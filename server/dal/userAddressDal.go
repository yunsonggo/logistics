package dal

import "orderserve/model"

type UserAddressInterface interface {
	// 添加地址
	InsertUserAddrOne(userAdd *model.UserAddress) (err error)
	// 根据用户ID 查询所有地址
	QueryUserAddrById(userId int64) (userAdds []model.UserAddress,err error)
	// 根据用户ID  清除默认收货地址
	UpdateAddrPosition(userId int64) (err error)
	// 根据用户ID 地址ID  修改 一条
	UpdateOneAddrById(userId int64,userAddr *model.UserAddress) (err error)
	// 根据用户ID 地址ID  删除一条
	DeleteOneAddrById(userId int64,userAddr *model.UserAddress) (err error)
}

type UserAddrDal struct {}

func NewUserAddrDal() UserAddressInterface {
	return &UserAddrDal{}
}

// 添加地址
func (uad *UserAddrDal) InsertUserAddrOne(userAdd *model.UserAddress) (err error) {
	_,err =DB.InsertOne(userAdd)
	return
}

// 根据用户ID 查询所有地址
func (uad *UserAddrDal) QueryUserAddrById(userId int64) (userAdds []model.UserAddress,err error) {
	err = DB.Where("user_id = ?",userId).Asc("position").Find(&userAdds)
	return
}
// 根据用户ID  清除默认收货地址
func (uad *UserAddrDal) UpdateAddrPosition(userId int64) (err error) {
	userAdds,_ := uad.QueryUserAddrById(userId)
	for _,addr := range userAdds {
		if addr.Position == 0 {
			addr.Position = 1
			_,updateErr := DB.Cols("position").Update(&addr)
			if updateErr != nil {
				return updateErr
			}
		}
	}
	return nil
}
// 根据用户ID 地址ID  修改 一条
func (uad *UserAddrDal) UpdateOneAddrById(userId int64,userAddr *model.UserAddress) (err error) {
	_,err = DB.Where("id = ? and user_id = ?",userAddr.Id,userId).Cols("user_name","user_gender","name","phone","mobile","address","bottom","lng","lat","position","tage").Update(userAddr)
	return
}
// 根据用户ID 地址ID  删除一条
func (uad *UserAddrDal) DeleteOneAddrById(userId int64,userAddr *model.UserAddress) (err error) {
	_,err = DB.Where("id = ? and user_id = ?",userAddr.Id,userId).Delete(userAddr)
	return
}
