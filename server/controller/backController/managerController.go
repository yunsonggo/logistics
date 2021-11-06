package backController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"mime/multipart"
	"net/http"
	"orderserve/common"
	"orderserve/model/backendModel"
	"orderserve/server/backServer"
	"orderserve/tools"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var manServer = backServer.NewManagerServer()
// 获取登录用户数据
func PostManagerInfo(ctx *gin.Context) {
	res,_ := ctx.Get("managerInfo")
	resManager := res.(*backendModel.Manager)
	if resManager.Id > 0 {
		common.Success(ctx,resManager)
		return
	}
	common.Failed(ctx,"获取用户数据失败")
	return
}
// 获取所有用户数据
func PostManagerList(ctx *gin.Context) {
	res,err := manServer.GetManagerAll()
	if err != nil {
		common.Failed(ctx,"获取工作人员列表数据失败")
		return
	}
	for key,_ := range res {
		res[key].PassWord = ""
	}
	common.Success(ctx,res)
	return
}
// 登录用户头像上传
func PostUploadAvatar(ctx *gin.Context) {
	//adminAvatarFile,_ := ctx.Get("adminAvatarFile")
	adminAvatarFileHeader,_ := ctx.Get("adminAvatarFileHeader")
	//file :=  adminAvatarFile.(multipart.File)
	fileHeader := adminAvatarFileHeader.(*multipart.FileHeader)
	//随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(fileHeader.Filename)
	fileName := "backendUpload/admin/" + "avatar" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	err := ctx.SaveUploadedFile(fileHeader,filePath)
	if err != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"path":"",
		})
		return
	}
	// 如果存在旧头像 则删除旧头像
	adminOldAvatarUrl,_ := ctx.Get("adminOldAvatarUrl")
	if adminOldAvatarUrl != nil {
		oldUrl := adminOldAvatarUrl.(string)
		oldFilePath := "public/" + oldUrl
		_ = os.Remove(oldFilePath)
	}
	// 提前更新数据库
	editAvatarAdminId,_ := ctx.Get("editAvatarAdminId")
	adminId := editAvatarAdminId.(string)
	// 根据ID 更新 人员 头像
	id,_ := strconv.ParseInt(adminId,10,64)
	//var manServer = backServer.NewManagerServer()
	editAvatarErr := manServer.EditManagerAvatarById(id,fileName)
	if editAvatarErr != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"path":"",
		})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{
		"path":fileName,
	})
	return
}
// 管理员添加工作人员
func PostAddManager(ctx *gin.Context) {
	managerInfo,has := ctx.Get("managerData")
	if !has {
		common.Failed(ctx,"获取用户数据失败")
		return
	}
	managerData := managerInfo.(backendModel.Manager)
	// 判断用户是否存在
	queryRes,_ := manServer.GetManagerByName(managerData.UserName)
	if queryRes.Id > 0 {
		// 删除上传的头像文件
		if managerData.UserIcon != "" {
			filePath := "public/" + managerData.UserIcon
			_ = os.Remove(filePath)
		}
		common.Failed(ctx,"用户已经存在")
		return
	}
	// 实现入库
	password := managerData.PassWord
	pass := tools.EncoderSha256(password)
	managerData.PassWord = pass
	resErr := manServer.AddManagerByAdmin(managerData)
	if resErr != nil {
		// 删除上传的头像文件
		if managerData.UserIcon != "" {
			filePath := "public/" + managerData.UserIcon
			_ = os.Remove(filePath)
		}
		common.Failed(ctx,"添加工作人员失败")
		return
	}
	managerData.PassWord = ""
	common.Success(ctx,managerData)
 	return
}
// 管理员编辑工作人员 基本信息
func PostEditManager(ctx *gin.Context) {
	managerInfo,has := ctx.Get("managerData")
	if !has {
		common.Failed(ctx,"获取用户数据失败")
		return
	}
	managerData := managerInfo.(backendModel.Manager)
	// 实现入库
	resErr := manServer.EditManagerInfoById(managerData)
	if resErr != nil {
		common.Failed(ctx,"编辑工作人员基本信息失败")
		return
	}
	managerData.PassWord = ""
	common.Success(ctx,managerData)
	return
}
// 管理员删除工作人员账号
func PostDelManager(ctx *gin.Context) {
	managerInfo,has := ctx.Get("managerData")
	if !has {
		common.Failed(ctx,"获取用户数据失败")
		return
	}
	managerData := managerInfo.(backendModel.Manager)
	// 实现删除
	if managerData.Power == "777" {
		common.Failed(ctx,"超级管理员账号不允许删除!")
		return
	}
	id := managerData.Id
	resErr := manServer.RemoveManagerById(id)
	if resErr != nil {
		common.Failed(ctx,"删除工作人员账号失败")
		return
	}
	// 删除管理员头像文件
	if managerData.UserIcon != "" {
		avatarPath := "public/" + managerData.UserIcon
		_ = os.Remove(avatarPath)
	}
	common.Success(ctx,managerData)
	return
}