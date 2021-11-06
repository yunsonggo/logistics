package model
// 已售商品
type SaleGoods struct {
	Id           int64    `xorm:"pk autoincr" json:"id"`            // ID
	SaleCode string `xorm:"varchar(255)" json:"sale_code"` 	// 已售商品识别码
	OrderId 	 int64 		`xorm:"int" json:"order_id"`			// 订单ID
	UserId       int64    `xorm:"int" json:"userid"`                // 用户ID
	SelectPay    int      `xorm:"int" json:"select_pay"`            // 0 微信 1 支付宝 ...
	ShopId       int64    `xorm:"int" json:"shop_id"`               // 商家ID
	ShopName     string   `xorm:"varchar(60)" json:"shop_name"`     // 商家名称
	Liaison      string   `xorm:"varchar(30)" json:"liaison"`       // 联系人
	LiaisonPhone string   `xorm:"varchar(11)" json:"liaison_phone"` // 联系人电话
	GoodId		 int64 `xorm:"int" json:"good_id"`	// 商品ID
	GoodName string `xorm:"varchar(30)" json:"good_name"` // 商品名称
	GoodPrice float64 `xorm:"float" json:"good_price"` // 当前商品 当前成交价格
	GoodCount int `xorm:"int" json:"good_count"` // 成交数量
	Exchange     string   `xorm:"varchar(30)" json:"exchange"`      // 调换货信息
	Invoice      string   `xorm:"varchar(30)" json:"invoice"`       // 发票信息
	CreateTime   JsonTime `xorm:"created" json:"create_time"`       // 时间
	UpdateTime   JsonTime `xorm:"updated" json:"update_time"`
	CateId int64 `xorm:"int" json:"cate_id"` // 分类ID
	Icon string `xorm:"varchar(255)" json:"icon"`
}
