package captcha_client

type CaptchaResponse struct {
	Result        bool
	CaptchaStatus string
	Ip            string
	Tpc           string
	Uid           string
	Code          string
}
