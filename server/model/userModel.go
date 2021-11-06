package model

//用户
type User struct {
	Id         int64  `xorm:"pk autoincr" json:"id"`
	UserName   string `xorm:"varchar(30)" json:"username"`
	UserIcon string `xorm:"varchar(255)" json:"user_icon"`
	Gender string `xorm:"varchar(20)" json:"user_gender"`
	Phone      string `xorm:"varchar(11)" json:"phone"`
	PassWord   string `xorm:"varchar(60)" json:"password"`
	IsDel      int    `xorm:"int" json:"is_del"` // 0 删除 1 激活
	ShopId string `xorm:"varchar(255)" json:"shop_id"` // 收藏店铺
	CreateTime JsonTime  `xorm:"created" json:"create_time"`
}
