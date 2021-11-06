package dal

import (
	"errors"
	"orderserve/model"
	"orderserve/tools"
)

type UserDalInterface interface {
	InsertUserBySms(phone string) (userId int64, err error)
	InsertUserByPass(name string, password string) (userId int64, err error)
	DelUserById(id int64) (err error)
	UpdateUserById(userData *model.User) (err error)
	QueryUserByPhone(phone string) (userData *model.User, err error)
	QueryUserById(id interface{}) (userData model.User, err error)
	QueryUserByName(name string) (userData *model.User, err error)
}

type UserDal struct{}

func NewUserDal() UserDalInterface {
	return &UserDal{}
}

// 手机短信 添加用户
func (ud *UserDal) InsertUserBySms(phone string) (userId int64, err error) {
	var user model.User
	_, err = DB.Where("phone = ?", phone).Get(&user)
	if err != nil {
		err = errors.New("添加用户查询手机号码失败")
		return
	}
	if user.Id > 0 {
		if user.IsDel == 0 {
			err = errors.New("该手机号已注册并注销,请联系管理员激活")
		} else {
			err = errors.New("该手机号已注册")
		}
		return 0, err
	}
	user.Phone = phone
	user.IsDel = 1
	_, err = DB.InsertOne(&user)
	if err != nil {
		err = errors.New("添加用户失败")
		return
	}
	res, _ := ud.QueryUserByPhone(phone)
	if res.Id > 0 {
		userId = res.Id
		err = nil
		return
	}
	return 0, err
}

// 账号密码 添加用户
func (ud *UserDal) InsertUserByPass(name string, password string) (userId int64, err error) {
	var user model.User
	pass := tools.EncoderSha256(password)
	_, err = DB.Where("username = ?", name).Get(&user)
	if err != nil {
		err = errors.New("添加用户查询用户名失败")
		return
	}
	if user.Id > 0 {
		if user.IsDel == 0 {
			err = errors.New("该用户已注册并注销,请联系管理员激活")
		} else {
			err = errors.New("该用户名已注册")
		}
		return 0, err
	}
	user.UserName = name
	user.PassWord = pass
	user.IsDel = 1
	_, err = DB.InsertOne(&user)
	if err != nil {
		err = errors.New("添加用户失败")
		return
	}
	res, _ := ud.QueryUserByName(name)
	if res.Id > 0 {
		return res.Id, nil
	}
	return 0, err
}

// 删除用户
func (ud *UserDal) DelUserById(id int64) (err error) {
	var user model.User
	_, err = DB.Where("id = ?", id).Get(&user)
	if err != nil {
		err = errors.New("删除用户失败")
		return
	}
	if user.Id > 0 {
		user.IsDel = 0
		_, err = DB.Where("id = ?", id).Update(&user)
		if err != nil {
			err = errors.New("删除用户失败")
			return
		}
		return nil
	}
	err = errors.New("用户不存在")
	return
}

// 根据ID 更新用户
func (ud *UserDal) UpdateUserById(userData *model.User) (err error) {
	id := userData.Id
	var user model.User
	_, err = DB.Where("id = ?", id).Get(&user)
	if err != nil {
		err = errors.New("更新用户数据失败")
		return
	}
	_, err = DB.Where("id = ?", id).Update(userData)
	if err != nil {
		err = errors.New("更新用户数据失败")
		return
	}
	return nil
}

// 根据手机 查询用户
func (ud *UserDal) QueryUserByPhone(phone string) (userData *model.User, err error) {
	var user model.User
	_, err = DB.Where("phone = ?", phone).Get(&user)
	if err != nil {
		err = errors.New("查询手机号失败")
		return
	}
	return &user, nil
}

//根据ID 查询用户
func (ud *UserDal) QueryUserById(id interface{}) (userData model.User, err error) {
	_, err = DB.Where("id = ?", id).Get(&userData)
	return
}

//根据用户名 查询用户
func (ud *UserDal) QueryUserByName(name string) (userData *model.User, err error) {
	var user model.User
	_, err = DB.Where("username = ?", name).Get(&user)
	if err != nil {
		err = errors.New("查询用户失败")
		return
	}
	return &user, nil
}
