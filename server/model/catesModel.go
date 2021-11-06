package model

// 商家分类
type Cates struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`
	Name        string `xorm:"varchar(30)" json:"name"`        //类别名称
	Description string `xorm:"varchar(60)" json:"description"` // 类别描述
	Icon        string `xorm:"varchar(255)" json:"icon"`       //类别图标
	IsHot       int    `xorm:"int" json:"is_hot"`              //常用 true 1 false 0
	IsTop       int    `xorm:"int" json:"is_top"`              //置顶
	State       int    `xorm:"int" json:"state"`               // 是否激活
	IsMain      int    `xorm:"int" json:"is_main"`             //是否主类
	ShopIds      string  `xorm:"varchar(255)" json:"shop_ids"` // 商家ID
	Position  int `xorm:"int" json:"position"` // 位置 排位
}
