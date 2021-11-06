package param

type LoginParam struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"verifyCode" binding:"required"`
}
