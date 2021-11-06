package model

type Goods struct {
	Id            int64   `xorm:"pk autoincr" json:"id"`
	Name          string  `xorm:"varchar(30)" json:"name"`        // 商品名称
	Description   string  `xorm:"varchar(60)" json:"description"` // 商品描述
	Content       string  `xorm:"text" json:"content"`            // 商品详情
	Icon          string  `xorm:"varchar(255)" json:"icon"`       // 商品图标
	Price         float32 `xorm:"float" json:"price"`             // 现售单价
	OldPrice      float32 `xorm:"float" json:"old_price"`         // 原价
	GoodsUnit     string  `xorm:"varchar(30)" json:"goods_unit"`  // 商品单位
	TotalNum      int64   `xorm:"int" json:"total_num"`           // 商品总量
	SellCount     int64   `xorm:"int" json:"sell_count"`          // 已售数量
	MinNum        int64   `xorm:"int" json:"min_num"`             // 起送数量
	DeliveryPrice float32 `xorm:"float" json:"delivery_price"`    // 单位运费
	OverPrice     float32 `xorm:"float" json:"over_price"`        // 超出运费
	ShopId        int64   `xorm:"int" json:"shop_id"`             // 商家ID
	CateId        int64   `xorm:"int" json:"cate_id"`             // 类别ID 商家专属分类 ShopCates
	GiveId        string  `xorm:"varchar(255)" json:"give_id"`    // 服务ID
	SaleId        string  `xorm:"varchar(255)" json:"sale_id"`    // 活动ID
	IsHot         int     `xorm:"int" json:"is_hot"`              // 热卖 true 1 false 0
	IsTop         int     `xorm:"int" json:"is_top"`              // 置顶
	State         int     `xorm:"int" json:"state"`               // 是否销售
	IsMain        int     `xorm:"int" json:"is_main"`             // 是否首页
	Score         int     `xorm:"int" json:"score"`               // 评分
	IsFree        int     `xorm:"int" json:"is_free"`             //是否包邮
}
