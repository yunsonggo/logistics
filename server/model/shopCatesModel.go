package model

//商家 专属 商品小类
type ShopCates struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	ScName string `xorm:"varchar(30)" json:"sc_name"`
	ScIcon string `xorm:"varchar(255)" json:"sc_icon"`
	ShopId int64 `xorm:"int" json:"shop_id"`
	GoodIds string `xorm:"varchar(255)" json:"good_ids"`
	State int `xorm:"int" json:"state"`
	Position int `xorm:"int" json:"position"`
	Desc string `xorm:"varchar(255)" json:"desc"`
}
