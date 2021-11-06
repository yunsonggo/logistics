package model

type UserAddress struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	UserId   int64   `xorm:"int" json:"user_id"`          // 用户
	UserName string `xorm:"varchar(20)" json:"user_name"` //联系人姓名
	UserGender string `xorm:"varchar(12)" json:"user_gender"` //联系人性别
	Name     string  `xorm:"varchar(30)" json:"name"`     // 地址名称
	Phone    string  `xorm:"varchar(20)" json:"phone"`    // 固定电话
	Mobile   string  `xorm:"varchar(11)" json:"mobile"`   // 移动电话
	Address  string  `xorm:"varchar(255)" json:"address"` // 地址街道
	Bottom string `xorm:"varchar(120)" json:"bottom"`     //门牌号
	Lng      float64 `xorm:"Double" json:"lng"`           // 维度
	Lat      float64 `xorm:"Double" json:"lat"`           // 经度
	Position int     `xorm:"int" json:"position"`         //排序位置  默认地址 为 0
	Tage     string  `xorm:"varchar(20)" json:"tage"`     //标签
}
