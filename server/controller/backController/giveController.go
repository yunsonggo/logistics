package backController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"mime/multipart"
	"net/http"
	"orderserve/common"
	"orderserve/model"
	"orderserve/server/backServer"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var gis = backServer.NewGiveServer()

//获取所有服务数据
func PostGiveList(ctx *gin.Context) {
	res,resErr := gis.GetGivesAll()
	if resErr != nil {
		common.Failed(ctx,"获取服务数据失败")
		return
	}
	common.Success(ctx,res)
	return
}
// 服务图标上传
func PostGiveIconUpload(ctx *gin.Context) {
	giveIconFileHeader,_ := ctx.Get("giveIconFileHeader")
	fileHeader := giveIconFileHeader.(*multipart.FileHeader)
	//随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(fileHeader.Filename)
	fileName := "upload/gives/" + "icon" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	saveFileErr := ctx.SaveUploadedFile(fileHeader,filePath)
	if saveFileErr != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"path":"",
		})
		return
	}
	// 如果存在旧icon  则删除旧icon
	oldIconUrl,_ := ctx.Get("oldGiveIconUrl")
	if oldIconUrl != nil {
		oldUrl := oldIconUrl.(string)
		oldFilePath := "public/" + oldUrl
		_ = os.Remove(oldFilePath)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"path":fileName,
	})
	return
}
// 更新服务数据
func PostGiveEdit(ctx *gin.Context) {
	giveInfo,hasData := ctx.Get("giveData")
	if !hasData {
		common.Failed(ctx,"获取服务数据失败")
		return
	}
	giveData := giveInfo.(model.Give)
	id := giveData.Id
	err := gis.EditGiveOne(id,giveData)
	if err != nil {
		common.Failed(ctx,"更新服务数据失败")
		return
	}
	common.Success(ctx,giveData)
	return
}
// 添加服务数据
func PostGiveAdd(ctx *gin.Context) {
	giveInfo,hasData := ctx.Get("giveData")
	if !hasData {
		common.Failed(ctx,"获取服务数据失败")
		return
	}
	giveData := giveInfo.(model.Give)
	err := gis.AddGiveOne(giveData)
	if err != nil {
		common.Failed(ctx,"添加服务数据失败")
		return
	}
	common.Success(ctx,giveData)
	return
}