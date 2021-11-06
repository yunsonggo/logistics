package dal

import (
	"orderserve/model"
)

// 短信数据操作接口
type SmsInterface interface {
	InsertSmsOne(data model.SmsCode) (err error)
	QuerySmsByPhone(phone string, code string) (resCode *model.SmsCode, err error)
	DelSmsByPhone(resId int64) (err error)
}

type SmsDal struct{}

func NewSmsDal() SmsInterface {
	return &SmsDal{}
}

// 插入 短信验证码
func (sd *SmsDal) InsertSmsOne(data model.SmsCode) (err error) {
	_, err = DB.InsertOne(&data)
	if err != nil {
		return err
	}
	return nil
}

//查询 短信验证码
func (sd *SmsDal) QuerySmsByPhone(phone string, code string) (resCode *model.SmsCode, err error) {
	var codeData model.SmsCode
	_, err = DB.Where("phone = ? and code = ?", phone, code).Get(&codeData)
	if err != nil {
		return nil, err
	}
	resCode = &codeData
	return
}

//删除 短信验证码
func (sd *SmsDal) DelSmsByPhone(resId int64) (err error) {
	codeData := model.SmsCode{
		Id: resId,
	}
	_, err = DB.Delete(&codeData)
	return err
}
