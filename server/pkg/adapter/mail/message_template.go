package mail

var (
	hostURL              = "https://wantty.app"
	signInEmailVerifyURL = hostURL + "/sign_in/email/verify?token="
	signUpEmailVerifyURL = hostURL + "/sign_up/email/verify?token="
	sendMailAddress      = "noreply@wantty.app"
	// sendMailAddress = "akasakatora1208@gmail.com"
	authMessage = &template{
		Title: "本人確認のお知らせ",
		Body: `サブスクリプション型コミュニティサービス「Wantty」をご利用いただきありがとうございます。
下記のリンクをクリックして、アカウント登録を完了させてください。
`,
	}
)

func getEmailTemplate(emailType string, option string) template {
	switch emailType {
	case emailTypes.signIn:
		return template{
			Title: "ログインの確認",
			Body: `「Wantty」にログインする場合は10分以内に下記のリンクをクリックしてください。
`,
		}
	case emailTypes.signUp:
		return template{
			Title: "本人確認のお知らせ",
			Body: `サブスクリプション型コミュニティサービス「Wantty」をご利用いただきありがとうございます。
10分以内に下記のリンクをクリックして、アカウント登録を完了させてください。
`,
		}
	default:
		return template{
			Title: "本人確認のお知らせ",
			Body: `サブスクリプション型コミュニティサービス「Wantty」をご利用いただきありがとうございます。
下記のリンクをクリックして、アカウント登録を完了させてください。
`,
		}
	}
}

func getSigninEmailVerifyURL(token string) string {
	return signInEmailVerifyURL + token
}

func getSignupEmailVerifyURL(token string) string {
	return signUpEmailVerifyURL + token
}
func getSuffixURL(emailType string, token string) string {
	url := ""
	switch emailType {
	case emailTypes.signIn:
		url = getSigninEmailVerifyURL(token)
	case emailTypes.signUp:
		url = getSignupEmailVerifyURL(token)
	default:
		url = hostURL
	}
	return url
}

func buildAuthMessage(emailType string, token string, option string) *template {
	url := getSuffixURL(emailType, token)
	emailTemplate := getEmailTemplate(emailType, option)
	emailTemplate.Body = emailTemplate.Body + url
	return &emailTemplate
}
