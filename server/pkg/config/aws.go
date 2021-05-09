package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

type S3 struct {
	PublicBucket  string
	PrivateBucket string
}
type AWSInfo struct {
	AccessKey string
	SecretKey string
	Region    string
	S3        S3
	Config    aws.Config
	Arn       string
}

func parseAWSConf(awsConf map[string]interface{}) *AWSInfo {
	region := awsConf["region"].(string)
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	arn := os.Getenv("ARN")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	return &AWSInfo{
		AccessKey: accessKey,
		SecretKey: secretKey,
		S3: S3{
			PublicBucket:  awsConf["public_bucket"].(string),
			PrivateBucket: awsConf["private_bucket"].(string),
		},
		Config: aws.Config{
			Region: &region,
			Credentials: credentials.NewStaticCredentials(
				accessKey,
				secretKey,
				"",
			),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
		},
		Region: region,
		Arn:    arn,
	}
}
