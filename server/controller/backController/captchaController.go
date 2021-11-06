package backController

import (
	"github.com/gin-gonic/gin"
	captcha_client "orderserve/captcha-client"
	"orderserve/common"
	"orderserve/param/backParam"
)

func VerifyCaptcha(ctx *gin.Context) {
	var captchaParam backParam.CaptchaParam
	err := ctx.ShouldBind(&captchaParam)
	if err != nil {
		common.Failed(ctx,"获取验证码TOKEN失败")
		return
	}
	appId := "62602185bf8b6f56e5b1e70fe7da4e60"
	appSecret := "b93b80ea39156be4edcc8142b42f369e"
	captchaClient := captcha_client.NewCaptchaClient(appId,appSecret)
	captchaClient.SetTimeout(10000)
	//设置超时时间，单位毫秒，默认2秒
	//captchaClient.SetCaptchaUrl(captchaUrl)
	//私有化版本需要额外指定服务器，默认情况下不需要设置
	captchaResponse := captchaClient.VerifyToken(captchaParam.CaptchaToken)
	//captchaResponse := captchaClient.VerifyTokenAndIP(token, ip)
	//针对一些token冒用的情况，业务方可以采集客户端ip随token一起提交到验证码服务，验证码服务除了判断token的合法性还会校验提交业务参数的客户端ip和验证码颁发token的客户端ip是否一致
	//fmt.Println(captchaResponse.CaptchaStatus)
	//确保验证状态是SERVER_SUCCESS，SDK中有容错机制，在网络出现异常的情况会返回通过
	//fmt.Println(captchaResponse.Ip)
	//验证码服务采集到的客户端ip
	if captchaResponse.Result {
		/*token验证通过，继续其他流程*/
		common.Success(ctx,"ok")
		return
	} else {
		/*token验证失败，业务系统可以直接阻断该次请求或者继续弹验证码*/
		common.Failed(ctx,"验证失败")
	}
}