package model
// 活动
type Sale struct {
	Id int64 `xorm:"pk autoincr" json:"id"`
	Name string `xorm:"varchar(30)" json:"name"`
	Icon string `xorm:"varchar(255)" json:"icon"`
	State int `xorm:"int" json:"state"`
	Seat int `xorm:"int" json:"seat"`
	Select string `xorm:"varchar(10)" json:"select"`
	Desc string `xorm:"varchar(255)" json:"desc"`
}
