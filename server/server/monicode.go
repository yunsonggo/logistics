package server

//模拟发送短信验证码
type ResponseSms struct {
	Message   string
	RequestId string
	BizId     string
	Code      string
}
