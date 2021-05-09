package config

import "os"

type StripeInfo struct {
	ProductionMode              bool
	PublishableKey              string
	SecretKey                   string
	WebhookSecretForUnsubscribe string
	WebhookSecretForSubscribe   string
	UseMock                     bool
	ProductID                   string
	LogLevel                    int
}

func parseStripeConf(stripeConf map[string]interface{}) *StripeInfo {
	publishable_key := os.Getenv("STRIPE_PUBLISHABLE_KEY")
	if publishable_key == "" {
		publishable_key = stripeConf["publishable_key"].(string)
	}
	secret_key := os.Getenv("STRIPE_SECRET_KEY")
	if secret_key == "" {
		secret_key = stripeConf["secret_key"].(string)
	}
	webhook_secret_for_unsubscribe := os.Getenv("STRIPE_WEBHOOK_SECRET_FOR_UNSUBSCRIBE")
	if webhook_secret_for_unsubscribe == "" {
		webhook_secret_for_unsubscribe = stripeConf["webhook_secret_for_unsubscribe"].(string)
	}
	webhook_secret_for_subscribe := os.Getenv("STRIPE_WEBHOOK_SECRET_FOR_SUBSCRIBE")
	if webhook_secret_for_subscribe == "" {
		webhook_secret_for_subscribe = stripeConf["webhook_secret_for_subscribe"].(string)
	}
	return &StripeInfo{
		ProductionMode:              stripeConf["production_mode"].(bool),
		PublishableKey:              publishable_key,
		SecretKey:                   secret_key,
		WebhookSecretForUnsubscribe: webhook_secret_for_unsubscribe,
		WebhookSecretForSubscribe:   webhook_secret_for_subscribe,
		UseMock:                     stripeConf["use_mock"].(bool),
		LogLevel:                    stripeConf["log_level"].(int),
	}
}
