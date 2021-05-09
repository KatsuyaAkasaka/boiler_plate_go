package mail

import (
	log "github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/adapter/logger"
	"github.com/KatsuyaAkasaka/boiler_plate_go/server/pkg/domain/entity"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type template struct {
	Title string
	Body  string
}

var emailTypes = struct {
	signUp string
	signIn string
}{
	"signUp",
	"signIn",
}

func SendSignInAuthEmail(email *entity.Email, token string) {
	sendEmail(email.ToStr(), buildAuthMessage(emailTypes.signIn, token, ""))
}

func SendSignUpAuthEmail(email *entity.Email, token string) {
	sendEmail(email.ToStr(), buildAuthMessage(emailTypes.signUp, token, ""))
}

// メール送信は完了するまで処理を止めず、失敗してもフロントにエラーを返さない
func sendEmail(to string, content *template) {
	go func(to string, content *template) {
		svc := ses.New(session.Must(session.NewSession()))
		input := &ses.SendEmailInput{
			Destination: &ses.Destination{
				ToAddresses: []*string{
					aws.String(to),
				},
			},
			Message: &ses.Message{
				Body: &ses.Body{
					Text: &ses.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(content.Body),
					},
				},
				Subject: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(content.Title),
				},
			},
			Source: aws.String(sendMailAddress),
		}
		_, err := svc.SendEmail(input)
		if err != nil {
			log.Error(err)
		}
	}(to, content)
	return
}
