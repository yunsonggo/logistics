package model

type Supplier struct {
	Id            int64   `xorm:"pk autoincr" json:"id"`
	Name          string  `xorm:"varchar(30)" json:"name"`        	// 商家名称
	Description   string  `xorm:"varchar(60)" json:"description"` 	// 商家公告
	Content       string  `xorm:"text" json:"content"`            	// 商家详情
	Phone         string  `xorm:"varchar(20)" json:"phone"`			// 固定电话
	Mobile        string  `xorm:"varchar(11)" json:"mobile"`		// 移动电话
	Address       string  `xorm:"varchar(255)" json:"address"`		// 地址
	Lng           float64 `xorm:"Double" json:"lng"`				// 经度
	Lat           float64 `xorm:"Double" json:"lat"`				// 维度
	Icon          string  `xorm:"varchar(255)" json:"icon"`         // 商家图标
	Banner		  string  `xorm:"varchar(255)" json:"banner"`		// banner
	TotalCount    int64   `xorm:"int" json:"total_count"`           // 订单总量
	MinNum        int64   `xorm:"int" json:"min_num"`               // 起送数量
	DeliveryPrice float32 `xorm:"float" json:"delivery_price"`      // 单位运费
	OverPrice     float32 `xorm:"float" json:"over_price"`          // 超出运费
	Rating        float32 `xorm:"float" json:"rating"`              // 商铺评分
	PackScore float64 `xorm:"float" json:"pack_score"`				// 包装分数
	ServeScore float64 `xorm:"float" json:"serve_score"`    		// 服务分数
	DeliveryScore float64 `xorm:"float" json:"delivery_score"`  	// 配送分数
	RatingCount   int64   `xorm:"int" json:"rating_count"`          // 评分总数
	OpeningHours  string  `xorm:"varchar(20)" json:"opening_hours"` // 营业时间
	GiveId        string   `xorm:"varchar(255)" json:"give_id"`     // 服务ID
	SaleId        string   `xorm:"varchar(255)" json:"sale_id"`		// 活动ID
	IsHot         int     `xorm:"int" json:"is_hot"` 				// 热卖 true 1 false 2
	IsTop         int     `xorm:"int" json:"is_top"` 				// 置顶
	IsNew         int     `xorm:"int" json:"is_new"`				// 新店
	State         int     `xorm:"int" json:"state"`   				// 状态 1 正常  2 关闭
	IsUnion		  int 	  `xorm:"int" json:"is_union" `				//是否联盟  1 签约 2 非签约
	IsMain        int     `xorm:"int" json:"is_main"` 				// 是否首页 1 是 2 非
	CateId 		  int64 `xorm:"int" json:"cate_id"` 				// 商家所属分类 对应 Cates
	CreateTime JsonTime  `xorm:"created" json:"create_time"`
}
