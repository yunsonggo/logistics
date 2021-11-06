package backServer

import (
	"orderserve/dal/backendDal"
	"orderserve/model/backendModel"
)

type IMenu interface {
	// 查询所有菜单
	GetMenuAll() (menusData []backendModel.BackendMenu,err error)
	// 添加菜单
	AddMenuOne(menuData *backendModel.BackendMenu) (err error)
	// 更新菜单
	EditMenuOne(menuData *backendModel.BackendMenu) (err error)
}

type MenuServer struct {}

func NewMenuServer() IMenu {
	return &MenuServer{}
}

var menuDal = backendDal.NewMenuDal()
// 查询所有菜单
func (mus *MenuServer) GetMenuAll() (menusData []backendModel.BackendMenu,err error) {
	return menuDal.QueryMenuAll()
}
// 添加菜单
func (mus *MenuServer) AddMenuOne(menuData *backendModel.BackendMenu) (err error) {
	return menuDal.InsertMenuOne(menuData)
}
// 更新菜单
func (mus *MenuServer) EditMenuOne(menuData *backendModel.BackendMenu) (err error) {
	return menuDal.UpdateMenuOne(menuData)
}