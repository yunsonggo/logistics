package backendDal

import (
	"orderserve/dal"
	"orderserve/model/backendModel"
)

type ManagerInterface interface {
	// 注册页面添加用户
	InsertManager(userName string,passWord string) (err error)
	// 管理员添加用户
	InsertManagerByAdmin(userInfo backendModel.Manager) (err error)
	// 根据管理员用户名 查询 是否已经存在
	QueryManagerByName(userName string) (manager *backendModel.Manager, err error)
	// 根据 管理员名称 密码 查询管理员
	QueryManagerByPass(userName string,passWord string) (manager *backendModel.Manager, err error)
	// 根据ID 查询 管理员
	QueryManagerById(id int64) (manager *backendModel.Manager, err error)
	// 获取所有管理员列表
	QueryManagerAll() (managersData []backendModel.Manager,err error)
	// 根据ID 更新管理员头像
	UpdateManagerAvatarById(id int64,fileName string) (err error)
	// 根据ID  更新管理员基本信息
	UpdataManagerInfoById(adminInfo backendModel.Manager)(err error)
	// 根据ID 删除管理员账号
	DeleteManagerById(id int64) (err error)
}

type ManagerDal struct {}

func NewManagerDal() ManagerInterface {
	return &ManagerDal{}
}
// 注册页面添加用户
func (md *ManagerDal) InsertManager(userName string,passWord string) (err error) {
	var manager backendModel.Manager
	manager.UserName = userName
	manager.PassWord = passWord
	manager.Power = "7"
	_,err = dal.DB.InsertOne(&manager)
	return
}
// 管理员添加用户
func (md *ManagerDal) InsertManagerByAdmin(userInfo backendModel.Manager) (err error) {
	_,err = dal.DB.InsertOne(&userInfo)
	return
}
// 根据用户名 查询 用户 是否已经存在
func (md *ManagerDal) QueryManagerByName(userName string) (manager *backendModel.Manager, err error) {
	var newManager backendModel.Manager
	_,err = dal.DB.Where("user_name = ?",userName).Get(&newManager)
	if err != nil {
		return nil,err
	}
	manager = &newManager
	return
}
// 根据 用户名 密码 查询管理员
func (md *ManagerDal) QueryManagerByPass(userName string,passWord string) (manager *backendModel.Manager, err error) {
	var newManager backendModel.Manager
	_,err = dal.DB.Where("user_name = ? and pass_word = ?",userName,passWord).Get(&newManager)
	if err != nil {
		return nil,err
	}
	manager = &newManager
	return
}
// 根据ID 查询 用户
func (md *ManagerDal) QueryManagerById(id int64) (manager *backendModel.Manager, err error) {
	var newManager backendModel.Manager
	_,err = dal.DB.Where("id = ?",id).Get(&newManager)
	if err != nil {
		return nil,err
	}
	manager = &newManager
	return
}
// 获取所有用户列表
func (md *ManagerDal) QueryManagerAll() (managersData []backendModel.Manager,err error) {
	err = dal.DB.Find(&managersData)
	return
}
// 根据ID 更新用户头像
func (md *ManagerDal) UpdateManagerAvatarById(id int64,fileName string) (err error) {
	newManager := new(backendModel.Manager)
	newManager.UserIcon = fileName
	_,err = dal.DB.Where("id = ?",id).Update(newManager)
	return
}
// 根据ID  更新用户基本信息
func (md *ManagerDal) UpdataManagerInfoById(adminInfo backendModel.Manager)(err error) {
	newManager := new(backendModel.Manager)
	newManager.Id = adminInfo.Id
	newManager.UserName = adminInfo.UserName
	newManager.UserIcon = adminInfo.UserIcon
	newManager.UserNick = adminInfo.UserNick
	newManager.Gender = adminInfo.Gender
	newManager.Phone = adminInfo.Phone
	newManager.Power = adminInfo.Power
	_,err = dal.DB.Where("id = ?",newManager.Id).Update(newManager)
	return
}
// 根据ID 删除管理员账号
func (md *ManagerDal) DeleteManagerById(id int64) (err error) {
	newManager := new(backendModel.Manager)
	newManager.Id = id
	_,err = dal.DB.Where("id = ?",id).Delete(newManager)
	return
}