package model

type Evaluation struct {
	Id            int64   `xorm:"pk autoincr" json:"id"`
	// 父级
	ParentId 	int64 `xorm:"int" json:"parent_id"`
	// 用户
	UserId 	int64 `xorm:"int" json:"user_id"`
	// 用户头像
	UserIcon string `xorm:"varchar(255)" json:"user_icon"`
	// 用户名字
	UserName string `xorm:"varchar(60)" json:"user_name"`
	// 商家
	ShopId int64 `xorm:"int" json:"shop_id"`
	// 商品
	GoodId int64	`xorm:"int" json:"good_id"`
	// 商品名称
	GoodName string `xorm:"varchar(255)" json:"good_name"`
	// 时间
	CreateTime JsonTime  `xorm:"created" json:"create_time"`
	UpdateTime JsonTime  `xorm:"updated" json:"update_time"`
	// 图片
	Icon string `xorm:"varchar(255)" json:"icon"`
	// 文字
	ContentText string `xorm:"text" json:"content_text"`
	// 回复
	Reply string `xorm:"text" json:"reply"`
	// 好评
	IsFine	int `xorm:"int" json:"is_fine"`
	// 中评
	IsGeneral int `xorm:"int" json:"is_general"`
	// 差评
	IsBad int `xorm:"int" json:"is_bad"`
	// 评分
	Score float64 `xorm:"float" json:"score"`
	// 包装分数
	PackScore float64 `xorm:"float" json:"pack_score"`
	// 服务分数
	ServeScore float64 `xorm:"float" json:"serve_score"`
	// 配送分数
	DeliveryScore float64 `xorm:"float" json:"delivery_score"`
	// 状态
	State int `xorm:"int" json:"state"`
}
