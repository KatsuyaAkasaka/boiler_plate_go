package cloud

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"path/filepath"

	e "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/error"
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/globalsign/mgo/bson"
)

type Repo struct {
	session   *session.Session
	awsConfig *config.AWSInfo
}

func InitUploader() *Repo {
	s, err := session.NewSession(&config.GetConf().AWS.Config)
	if err != nil {
		log.Fatal(err)
	}
	return &Repo{
		session:   s,
		awsConfig: config.GetConf().AWS,
	}
}

// UploadFileToS3 saves a file to aws bucket and returns the url to // the file and an error if there's any
func (repo *Repo) UploadImageToS3(fileHeader *multipart.FileHeader) (string, e.Err) {
	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	f, err := fileHeader.Open()
	if err != nil {
		log.Error(err)
		return "", e.System.Unexported
	}

	buffer := make([]byte, size)
	f.Read(buffer)

	tempFileName := "images/" + bson.NewObjectId().Hex() + filepath.Ext(fileHeader.Filename)

	_, er := s3.New(repo.session).PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(repo.awsConfig.S3.PublicBucket),
		Key:           aws.String(tempFileName),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(int64(size)),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})
	if er != nil {
		log.Error(er)
		return "", e.System.S3Err
	}

	return tempFileName, nil
}
