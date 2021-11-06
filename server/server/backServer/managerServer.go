package backServer

import (
	"orderserve/dal/backendDal"
	"orderserve/model/backendModel"
)

type IManager interface {
	// 注册页面添加管理员
	AddManager(userName string,passWord string) (err error)
	// 管理员添加管理员
	AddManagerByAdmin(userInfo backendModel.Manager) (err error)
	// 根据管理员名称 查询 是否已经存在
	GetManagerByName(userName string) (manager *backendModel.Manager, err error)
	// 根据 管理员账号 密码 查询管理员
	GetManagerByPass(userName string,passWord string) (manager *backendModel.Manager, err error)
	// 根据ID 查询 管理员
	GetManagerById(id int64) (manager *backendModel.Manager, err error)
	// 获取所有管理员列表
	GetManagerAll() (managersData []backendModel.Manager,err error)
	// 根据ID 更新管理员头像
	EditManagerAvatarById(id int64,fileName string) (err error)
	// 根据ID  更新管理员基本信息
	EditManagerInfoById(adminInfo backendModel.Manager)(err error)
	// 根据ID 删除管理员账号
	RemoveManagerById(id int64) (err error)
}

type ManagerServer struct {}

func NewManagerServer() IManager {
	return &ManagerServer{}
}

var md = backendDal.NewManagerDal()

// 添加用户
func (ms *ManagerServer) AddManager(userName string,passWord string) (err error) {
	return md.InsertManager(userName,passWord)
}
// 管理员添加用户
func (ms *ManagerServer) AddManagerByAdmin(userInfo backendModel.Manager) (err error) {
	return md.InsertManagerByAdmin(userInfo)
}
// 根据用户名 查询 用户 是否已经存在
func (ms *ManagerServer) GetManagerByName(userName string) (manager *backendModel.Manager, err error) {
	return md.QueryManagerByName(userName)
}
// 根据 用户名 密码 查询管理员
func (ms *ManagerServer) GetManagerByPass(userName string,passWord string) (manager *backendModel.Manager, err error) {
	return md.QueryManagerByPass(userName,passWord)
}
// 根据ID 查询 用户
func (ms *ManagerServer) GetManagerById(id int64) (manager *backendModel.Manager, err error) {
	return md.QueryManagerById(id)
}
// 获取所有用户列表
func (ms *ManagerServer) GetManagerAll() (managersData []backendModel.Manager,err error) {
	return md.QueryManagerAll()
}
// 根据ID 更新用户头像
func (ms *ManagerServer) EditManagerAvatarById(id int64,fileName string) (err error) {
	return md.UpdateManagerAvatarById(id,fileName)
}
// 根据ID  更新用户基本信息
func (ms *ManagerServer) EditManagerInfoById(adminInfo backendModel.Manager)(err error) {
	return md.UpdataManagerInfoById(adminInfo)
}
// 根据ID 删除管理员账号
func (ms *ManagerServer) RemoveManagerById(id int64) (err error) {
	return md.DeleteManagerById(id)
}