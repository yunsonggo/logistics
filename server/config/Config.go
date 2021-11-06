package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"os"
)

type Config struct {
	App struct {
		AppName string `json:"app_name"`
		AppUrl  string `json:"app_url"`
		AppMode string `json:"app_mode"`
		AppWeb  string `json:"app_web"`
		AppJwt  string `json:"app_jwt"`
	}
	Session struct {
		SessionName   string `json:"session_name"`
		SessionSecret string `json:"session_secret"`
		SessionMode   string `json:"session_mode"`
	}
	Mysql struct {
		MysqlDriver string `json:"mysql_driver"`
		MysqlAddr   string `json:"mysql_addr"`
		MysqlShow   bool   `json:"mysql_show"`
	}
	Redis struct {
		RedisAddr string `json:"redis_addr"`
		RedisPwd  string `json:"redis_pwd"`
		RedisHold int    `json:"redis_hold"`
	}
	RabbitMq struct {
	}
	AliSms struct {
		SmsName      string `json:"sms_name"`
		TemplateCode string `json:"template_code"`
		AppKey       string `json:"app_key"`
		AppSecret    string `json:"app_secret"`
		RegionId     string `json:"region_id"`
		TplId        string `json:"tpl_id"`
		Key          string `json:"key"`
	}
	Jwt struct {
		JwtKey string `json:"jwt_key"`
		Issuer string `json:"issuer"`
	}
}

var Conf Config

func ConfInit() {
	var cfg = new(Config)
	file, err := os.Open("config/config.json")
	if err != nil {
		err := errors.New("读取配置文件失败")
		panic(err.Error())
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	err = decoder.Decode(&cfg)
	if err != nil {
		err := errors.New("解析配置文件失败")
		panic(err.Error())
	}
	Conf = *cfg
	if Conf.App.AppMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	return
}
