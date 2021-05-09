module github.com/KatsuyaAkasaka/boiler_plate_go

go 1.13

replace github.com/KatsuyaAkasaka/boiler_plate_go => ./

require (
	github.com/aws/aws-sdk-go v1.38.33
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-playground/assert/v2 v2.0.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/mock v1.4.4
	github.com/mitchellh/mapstructure v1.1.2
	github.com/onsi/ginkgo v1.16.2 // indirect
	github.com/onsi/gomega v1.11.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stripe/stripe-go/v72 v72.43.0
	go.uber.org/zap v1.10.0
	gorm.io/driver/mysql v1.0.6
	gorm.io/gorm v1.21.9
)
