package entity

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/config"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type Model struct {
	CreatedAt time.Time      `sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time      `sql:"DEFAULT:current_timestamp on update current_timestamp"`
	DeletedAt gorm.DeletedAt `sql:"index"`
}

const (
	idSeparater            = "_"
	cloudFrontPrefix       = "https://assets.wantty.app/"
	publisedMovieURLPrefix = "movies/converts/"
	publisedImageURLPrefix = "images/"
)

func createShaID(elems []string) string {
	targetStr := ""
	for _, elem := range elems {
		targetStr += elem
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(targetStr)))
}

type Pagenation int

type PagenationInfo struct {
	Current int
	Offset  int
	Limit   int
}

func getPagenation(i int) (*Pagenation, e.Err) {
	p := Pagenation(i)
	if p <= 0 {
		return nil, e.System.BadRequest
	}
	return &p, nil
}

func (p Pagenation) getPagenationNum() int {
	return int(p)
}

func (p Pagenation) getPagenationInfo(size int) *PagenationInfo {
	pageNum := p.getPagenationNum()
	return &PagenationInfo{
		Current: pageNum,
		Offset:  (pageNum - 1) * size,
		Limit:   size,
	}
}

func IsExpired(expTime int64) bool {
	return time.Now().Unix() > expTime
}

type contentID string

func Decode(input interface{}, result interface{}) e.Err {
	if err := mapstructure.Decode(input, result); err != nil {
		return e.System.Unexported
	}
	return nil
}

type Card struct {
	ID       string `json:id`
	Brand    string `json:brand`
	Country  string `json:country`
	CVCCheck string `json:cvc_check`
	ExpMonth int    `json:exp_month`
	ExpYear  int    `json:exp_year`
	Funding  string `json:funding`
	Last4    string `json:last4`
	// Object   string `json:object`
}
type Cards []Card

func (cs Cards) ToSlice() []Card {
	var cards []Card
	for _, v := range cs {
		cards = append(cards, v)
	}
	return cards
}

type Subscription struct {
	ID               string    `json:id`
	CurrentPeriodEnd time.Time `json:current_period_end`
	UnitAmmount      int       `json:unit_ammount`
	Status           string    `json:status`
}
type Subscriptions []Subscription

func (ss Subscriptions) ToSlice() []Subscription {
	var subscriptions []Subscription
	for _, v := range ss {
		subscriptions = append(subscriptions, v)
	}
	return subscriptions
}

func BuildImagePathForCDN(path string) string {
	slice := strings.Split(path, "/")
	fileName := slice[len(slice)-1]
	return cloudFrontPrefix + publisedImageURLPrefix + fileName
}

func CheckPass(pass string) bool {
	return config.GetConf().Gateway.AdminPass == pass
}
