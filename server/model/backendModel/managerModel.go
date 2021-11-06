package backendModel

import "orderserve/model"

//用户
type Manager struct {
	Id         int64          `xorm:"pk autoincr" json:"id"`
	UserName   string         `xorm:"varchar(30)" json:"username"`
	UserNick   string         `xorm:"varchar(30)" json:"usernick"`
	UserIcon   string         `xorm:"varchar(255)" json:"user_icon"`
	Gender     string         `xorm:"varchar(20)" json:"user_gender"`
	Phone      string         `xorm:"varchar(11)" json:"phone"`
	PassWord   string         `xorm:"varchar(255)" json:"password"`
	IsDel      int            `xorm:"int" json:"is_del"`        // 0 删除 1 激活
	Power      string         `xorm:"varchar(12)" json:"power"` // 7 普通  77 管理 777 超级
	CreateTime model.JsonTime `xorm:"created" json:"create_time"`
}
