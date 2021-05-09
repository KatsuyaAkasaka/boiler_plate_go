package charge

import (
	"strconv"
	"sync"
	"time"

	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/config"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
	"github.com/stripe/stripe-go/v72/webhook"
)

type Repo struct {
	api    *client.API
	config *config.StripeInfo
}

var duration = time.Millisecond * 500
var callCount = 20

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func InitStripe() *Repo {
	// sc := &client.API{}
	stripeConf := config.GetConf().Stripe
	stripe.DefaultLeveledLogger = &stripe.LeveledLogger{
		Level: stripe.LevelInfo,
	}

	config := &stripe.BackendConfig{}
	// mock利用の場合はlocalhostのstripe-mockサーバに接続
	if stripeConf.UseMock {
		config.URL = stripe.String("http://localhost:12111")
	}

	backends := &stripe.Backends{
		API:     stripe.GetBackendWithConfig(stripe.APIBackend, config),
		Uploads: stripe.GetBackendWithConfig(stripe.UploadsBackend, config),
	}
	sc := client.New(stripeConf.SecretKey, backends)

	return &Repo{
		api:    sc,
		config: stripeConf,
	}
}

func (repo Repo) CreatePrice(price int, uuid *entity.UUID) (string, e.Err) {
	productParam := &stripe.ProductParams{
		Name: stripe.String(uuid.ToStr() + ":" + strconv.Itoa(price)),
	}
	// productを作成して、そのproductにplanを紐付ける
	product, err := repo.api.Products.New(productParam)
	if err != nil {
		log.Error(err)
		log.Errorf("failed to creatte price_id. uuid: %s", uuid.ToStr())
		return "", e.System.StripeErr
	}
	nowUnixStr := strconv.FormatInt(time.Now().Unix(), 10)
	params := &stripe.PriceParams{
		UnitAmount: stripe.Int64(int64(price)),
		Currency:   stripe.String(string(stripe.CurrencyJPY)),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String("month"),
		},
		Product:  stripe.String(product.ID),
		Nickname: stripe.String(uuid.ToStr() + ":" + nowUnixStr),
	}
	priceData, err := repo.api.Prices.New(params)
	if err != nil {
		log.Error(err)
		return "", e.System.StripeErr
	}
	if priceData.UnitAmount != int64(price) {
		log.Error("expected amount and actualy created amount is different.")
		return "", e.System.StripeErr
	}
	log.Infof("created price successfuly: %+v", priceData)
	return priceData.ID, nil
}

func (repo Repo) CreateCustomer(uuid *entity.UUID, email *entity.Email) (string, e.Err) {
	nowStr := time.Now().Format(time.RFC3339)
	params := &stripe.CustomerParams{
		Name:        stripe.String(uuid.ToStr()),
		Email:       stripe.String(email.ToStr()),
		Description: stripe.String("created by " + uuid.ToStr() + " at " + nowStr),
	}
	c, err := repo.api.Customers.New(params)
	if err != nil {
		log.Errorf("[stripe] failed to creaete customer. uuid: %s, err: %+v", uuid.ToStr(), err)
		return "", e.System.StripeErr
	}
	log.Infof("[stripe] created customer successfuly: %+v", c)
	return c.ID, nil
}

func (repo Repo) Subscribe(customerID string, priceID string, cardToken string) (*entity.Subscription, e.Err) {
	cardParam := &stripe.CardParams{
		Customer: stripe.String(customerID),
		Token:    stripe.String(cardToken),
	}
	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price:    stripe.String(priceID),
				Quantity: stripe.Int64(1),
			},
		},
		Card: cardParam,
	}
	s, err := repo.api.Subscriptions.New(params)
	if err != nil {
		log.Error(err)
		return nil, e.System.StripeErr
	}
	log.Infof("[stripe] subscription created successfuly: %+v", s)
	return &entity.Subscription{
		ID:               s.ID,
		CurrentPeriodEnd: time.Unix(s.CurrentPeriodEnd, 0),
		Status:           string(s.Status),
	}, nil
}

func (repo Repo) UnSubscribe(subscriptionID string) e.Err {
	s, err := repo.api.Subscriptions.Update(subscriptionID, &stripe.SubscriptionParams{
		// サブスク有効日まで削除しないでいてくれる
		CancelAtPeriodEnd: stripe.Bool(true),
	})
	if err != nil {
		log.Errorf("[stripe] failed to unsubscribe from stripe subscriptionID: %s, err %+v", subscriptionID, err)
		return e.System.StripeErr
	}
	log.Infof("[stripe] subscription canceled successfuly: %+v", s)
	return nil
}

func (repo Repo) UnSubscribeImmidiately(subscriptionID string) e.Err {
	s, err := repo.api.Subscriptions.Update(subscriptionID, &stripe.SubscriptionParams{
		// サブスク即削除
		CancelAtPeriodEnd: stripe.Bool(false),
	})
	if err != nil {
		log.Errorf("[stripe] failed to unsubscribe from stripe subscriptionID: %s, err %+v", subscriptionID, err)
		return e.System.StripeErr
	}
	log.Infof("[stripe] subscription canceled successfuly: %+v", s)
	return nil
}

func (repo Repo) UpdateCard(customerID string, subscriptionID string, card *entity.Card) (*entity.Subscription, e.Err) {
	params := &stripe.SubscriptionParams{
		DefaultSource: &card.ID,
	}
	s, er := repo.api.Subscriptions.Update(subscriptionID, params)
	if er != nil {
		log.Error(er)
		return nil, e.System.StripeErr
	}
	log.Infof("[stripe] subscription updated successfuly: %+v", s)
	return &entity.Subscription{
		ID:               s.ID,
		CurrentPeriodEnd: time.Unix(s.CurrentPeriodEnd, 0),
		Status:           string(s.Status),
	}, nil
}

func (repo Repo) GetCardByToken(customerID string, token string) (*entity.Card, e.Err) {
	params := &stripe.CardParams{
		Customer: stripe.String(customerID),
		Token:    stripe.String(token),
	}
	card, err := repo.api.Cards.New(params)
	if err != nil {
		log.Error(err)
		return nil, e.System.StripeErr
	}
	return &entity.Card{
		ID:       card.ID,
		Brand:    string(card.Brand),
		Country:  card.Country,
		CVCCheck: string(card.CVCCheck),
		ExpMonth: int(card.ExpMonth),
		ExpYear:  int(card.ExpYear),
		Funding:  string(card.Funding),
		Last4:    card.Last4,
	}, nil
}

func (repo Repo) GetPriceCost(id string) (int, e.Err) {
	price, err := repo.api.Prices.Get(id, nil)
	if err != nil {
		log.Errorf("[stripe] failed to find card from stripe cardID: %s, err %+v", id, err)
		return 0, e.System.StripeErr
	}
	return int(price.UnitAmount), nil
}

func (repo Repo) GetCard(id string, customerID string) (*entity.Card, e.Err) {
	s, err := repo.api.Cards.Get(id, &stripe.CardParams{
		Customer: stripe.String(customerID),
	})
	if err != nil {
		log.Errorf("[stripe] failed to find card from stripe cardID: %s, err %+v", id, err)
		return nil, e.System.StripeErr
	}
	return &entity.Card{
		ID:       s.ID,
		Brand:    string(s.Brand),
		Country:  s.Country,
		CVCCheck: string(s.CVCCheck),
		ExpMonth: int(s.ExpMonth),
		ExpYear:  int(s.ExpYear),
		Funding:  string(s.Funding),
		Last4:    s.Last4,
	}, nil
}

func (repo Repo) GetCards(ids []string, customerID string) (*entity.Cards, e.Err) {
	pos := 0
	cards := []entity.Card{}
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	for {
		endPos := min(len(ids), pos+callCount)
		targetCardIDs := ids[pos:endPos]
		for _, cardID := range targetCardIDs {
			wg.Add(1)
			go func(cardID string) {
				defer wg.Done()
				card, err := repo.GetCard(cardID, customerID)
				if err != nil {
					log.Errorf("[stripe] failed to find card from stripe cardID: %s, err: %+v", cardID, err)
					return
				}
				mutex.Lock()
				cards = append(cards, *card)
				mutex.Unlock()
			}(cardID)
		}
		wg.Wait()
		if endPos >= len(ids) {
			break
		}
		pos = pos + callCount
		time.Sleep(duration)
	}
	var res entity.Cards = cards
	return &res, nil
}

func (repo Repo) GetSubscription(id string) (*entity.Subscription, e.Err) {
	s, err := repo.api.Subscriptions.Get(id, nil)
	if err != nil {
		log.Error(err)
		return nil, e.System.StripeErr
	}
	status := string(s.Status)
	// キャンセル予定の場合はステータスを上書きする
	if s.CancelAt >= time.Now().Unix() {
		status = "canceling"
	}
	return &entity.Subscription{
		ID:               s.ID,
		CurrentPeriodEnd: time.Unix(s.CurrentPeriodEnd, 0),
		UnitAmmount:      int(s.Items.Data[0].Price.UnitAmount),
		Status:           status,
	}, nil
}

func (repo Repo) GetSubscriptions(ids []string) (*entity.Subscriptions, e.Err) {
	pos := 0
	subscriptions := []entity.Subscription{}
	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	for {
		endPos := min(len(ids), pos+callCount)
		targetSubIDs := ids[pos:endPos]
		for _, subID := range targetSubIDs {
			wg.Add(1)
			go func(subID string) {
				defer wg.Done()
				subscription, err := repo.GetSubscription(subID)
				if err != nil {
					log.Errorf("[stripe] failed to find subscription from stripe. subscriptionID: %s, err: %+v", subID, err)
					return
				}
				mutex.Lock()
				subscriptions = append(subscriptions, *subscription)
				mutex.Unlock()
			}(subID)
		}
		wg.Wait()
		if endPos >= len(ids) {
			break
		}
		pos = pos + callCount
		time.Sleep(duration)
	}
	var res entity.Subscriptions = subscriptions
	return &res, nil
}

func (repo Repo) signatureCheck(body []byte, header string, sign string) (*stripe.Event, e.Err) {
	// 署名チェック
	event, err := webhook.ConstructEvent(body, header, sign)
	if err != nil {
		log.Error(err)
		return nil, e.System.StripeErr
	}
	return &event, nil
}

func (repo Repo) GetUnsubscribeIDByWebhook(body []byte, header string) (string, e.Err) {
	sign := repo.config.WebhookSecretForUnsubscribe
	event, err := repo.signatureCheck(body, header, sign)
	if err != nil {
		return "", err
	}
	var subID string
	// サブスク解約通知以外は受け取らない
	if event.Type == "customer.subscription.deleted" {
		subID = event.Data.Object["id"].(string)

	} else if !(event.Type == "invoice.payment_failed" && event.Data.Object["next_payment_attempt"] == nil) {
		subID = event.Data.Object["subscription"].(string)
		// 支払い失敗の通知以外は受け取らない
	} else {
		return "", e.System.BadRequest
	}
	return subID, nil
}

func (repo Repo) GetSuccessSubscribeIDByWebhook(body []byte, header string) (string, string, e.Err) {
	sign := repo.config.WebhookSecretForSubscribe
	event, err := repo.signatureCheck(body, header, sign)
	if err != nil {
		return "", "", err
	}
	// サブスク成功通知以外の通知は受け取らない
	if event.Type != "invoice.payment_succeeded" {
		return "", "", e.System.BadRequest
	}
	subID := event.Data.Object["subscription"].(string)
	cusID := event.Data.Object["customer"].(string)
	return subID, cusID, nil
}
