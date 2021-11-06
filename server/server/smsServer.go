package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"math/rand"
	"orderserve/config"
	"orderserve/dal"
	"orderserve/model"
	"time"
)

type ISmsServer interface {
	SendSms(phone string, tplId string, key string) (codeStr string, err error)
	AliSendSms(phone string, tplId string, key string) (err error)
	CheckSms(phone string, code string) (resId int64, err error)
	RemoveSms(id int64) (err error)
}

type SmsServer struct{}

func NewSmsServer() ISmsServer {
	return &SmsServer{}
}

var smsDal = dal.NewSmsDal()

// 模拟发送验证码
func (ss *SmsServer) SendSms(phone string, tplId string, key string) (codeStr string, err error) {
	tpl := "ju-he-sms-" + tplId
	k := "ju-he-sms-" + key
	if tpl != config.Conf.AliSms.TplId || k != config.Conf.AliSms.Key {
		err = errors.New("发送验证码特征错误")
		return "", err
	}
	//产生验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	var response = ResponseSms{
		Message:   "OK",
		RequestId: code,
		BizId:     code,
		Code:      "OK",
	}
	//发送短信成功  调用 入库
	smsCode := model.SmsCode{
		Phone:      phone,
		BizId:      response.BizId,
		Code:       code,
		CreateTime: time.Now().Unix(),
	}
	err = smsDal.InsertSmsOne(smsCode)
	if err != nil {
		return "", err
	}
	codeStr = code
	return codeStr, nil
}

// 阿里云发送验证码
func (ss *SmsServer) AliSendSms(phone string, tplId string, key string) (err error) {
	tpl := "ju-he-sms-" + tplId
	k := "ju-he-sms-" + key
	if tpl != config.Conf.AliSms.TplId || k != config.Conf.AliSms.Key {
		err = errors.New("发送验证码特征错误")
		return err
	}
	//产生验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//获取阿里云配置
	smsConf := config.Conf.AliSms
	//创建连接
	client, err := dysmsapi.NewClientWithAccessKey(smsConf.RegionId, smsConf.AppKey, smsConf.AppSecret)
	if err != nil {
		return err
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = smsConf.SmsName
	request.TemplateCode = smsConf.TemplateCode
	request.PhoneNumbers = phone
	//固定格式
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)
	//获取返回数据
	response, err := client.SendSms(request)
	if err != nil {
		return err
	}
	//发送成功
	if response.Code == "OK" {
		smsCode := model.SmsCode{
			Phone:      phone,
			BizId:      response.BizId,
			Code:       code,
			CreateTime: time.Now().Unix(),
		}
		//入库
		err = smsDal.InsertSmsOne(smsCode)
		if err != nil {
			return err
		}
		return nil
	}
	//发送失败
	err = errors.New("发送短信失败")
	return err
}

// 验证短信验证码
func (ss *SmsServer) CheckSms(phone string, code string) (resId int64, err error) {
	resCode, resErr := smsDal.QuerySmsByPhone(phone, code)
	if resErr != nil || resCode.Id == 0 {
		return
	}
	return resCode.Id, nil
}

//删除 短信验证码
func (ss *SmsServer) RemoveSms(id int64) (err error) {
	return smsDal.DelSmsByPhone(id)
}
