package param

type SmsParam struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code"`
	TplId string `json:"tpl_id" binding:"required"`
	Key   string `json:"key" binding:"required"`
}
