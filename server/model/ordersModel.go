package model

type Orders struct {
	Id            int64    `xorm:"pk autoincr" json:"id"`            // 订单ID
	OrderCode 	 string `xorm:"varchar(255)" json:"order_code"`		// 订单识别码
	UserId        int64    `xorm:"int" json:"userid"`                // 用户ID
	UserName      string   `xorm:"varchar(60)" json:"user_name"`     // 用户姓名
	UserPhone     string   `xorm:"varchar(11)" json:"user_phone"`    // 用户电话
	TotalPrice    float64  `xorm:"float" json:"total_price"`         // 总价
	SelectPay     string   `xorm:"string" json:"select_pay"`         // 0 微信 1 支付宝 ...
	ShopId        int64    `xorm:"int" json:"shop_id"`               // 商家ID
	ShopName      string   `xorm:"varchar(60)" json:"shop_name"`     // 商家名称
	ShopPhone     string   `xorm:"varchar(11)" json:"shop_phone"`    // 商家电话
	GoodsInfo     string   `xorm:"varchar(255)" json:"goods_info"`   // 商品信息 [商品 1]:_id:1_名称:aa_单价:11_数量11|[商品 2]:_id:2_名称:bb_单价:22_数量22,
	AddressId     int64    `xorm:"int" json:"address_id"`            // 地址ID
	Liaison       string   `xorm:"varchar(30)" json:"liaison"`       // 联系人
	Gender        string   `xorm:"varchar(12)" json:"gender"`        // 联系人性别
	LiaisonPhone  string   `xorm:"varchar(11)" json:"liaison_phone"` // 联系人电话
	Address       string   `xorm:"varchar(255)" json:"address"`      // 具体地址
	Lng           float64  `xorm:"Double" json:"lng"`                // 经度
	Lat           float64  `xorm:"Double" json:"lat"`                // 纬度
	Exchange      string   `xorm:"varchar(30)" json:"exchange"`      // 调换货信息
	Remark        string   `xorm:"varchar(255)" json:"remark"`       // 备注信息
	Invoice       string   `xorm:"varchar(30)" json:"invoice"`       // 发票信息
	CreateTime    JsonTime `xorm:"created" json:"create_time"`       // 时间
	UpdateTime    JsonTime `xorm:"updated" json:"update_time"`
	State         int      `xorm:"int" json:"state"`                  // 订单状态 0: 未付款 1：已付款未送达 2: 已完成
	DeliveryPrice float64  `xorm:"float" json:"delivery_price"`       // 运费
	GoodsPrice    float64  `xorm:"float" json:"goods_price"`          // 商品总价格
	Metre         float64  `xorm:"float" json:"metre"`                // 运输路径距离km
	IsUnion       int      `xorm:"int" json:"is_union"`               // 是否签约商家
	SaleGoodsId   string   `xorm:"varchar(255)" json:"sale_goods_id"` // 已销售商品ID 数组
	DeliveryDateTime string `xorm:"varchar(255)" json:"delivery_datetime"` // 订单预计送达时间
	SaleGoodsList []SaleGoods `json:"sale_goods_list"`
	ShopInfo Supplier `json:"shop_info"`
}
