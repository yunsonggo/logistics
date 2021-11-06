package backendModel

type BackendMenu struct {
	Id           int64  `xorm:"pk autoincr" json:"id"`
	ParentId     int64  `xorm:"int" json:"parent_id"`             // 0 一级菜单
	MenuTitle    string `xorm:"varchar(30)" json:"menu_title"`    // 菜单名称
	IsShow       int    `xorm:"int" json:"is_show"`               // 1 不显示  2 显示
	ManagerPower string `xorm:"varchar(11)" json:"manager_power"` // 7 77 777
	MenuPath     string `xorm:"varchar(255)" json:"menu_path"`
	Icon         string `xorm:"varchar(50)" json:"icon"`
}
