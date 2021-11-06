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

var cgs = backServer.NewCategoryServer()

// 商家类别icon 上传
func PostEditCateIcon(ctx *gin.Context) {
	cateGoryIconFileHeader,_ := ctx.Get("cateGoryIconFileHeader")
	fileHeader := cateGoryIconFileHeader.(*multipart.FileHeader)
	//随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(fileHeader.Filename)
	fileName := "upload/cates/" + "icon" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	saveFileErr := ctx.SaveUploadedFile(fileHeader,filePath)
	if saveFileErr != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"path":"",
		})
		return
	}
	// 如果存在旧icon  则删除旧icon
	oldIconUrl,_ := ctx.Get("oldCateIconUrl")
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
// 商家类别更新
func PostEditCategory(ctx *gin.Context) {
	cateGoryData,hasData := ctx.Get("categoryData")
	if !hasData {
		common.Failed(ctx,"获取类别数据失败")
		return
	}
	categoryData := cateGoryData.(model.Cates)
	id := categoryData.Id
	err := cgs.EditCateGoryById(id,categoryData)
	if err != nil {
		common.Failed(ctx,"更新类别数据失败")
		return
	}
	common.Success(ctx,categoryData)
	return
}
// 添加商家类别
func PostAddCategory(ctx *gin.Context) {
	cateGoryData,hasData := ctx.Get("categoryData")
	if !hasData {
		common.Failed(ctx,"获取类别数据失败")
		return
	}
	categoryData := cateGoryData.(model.Cates)
	err := cgs.AddCategoryOne(categoryData)
	if err != nil {
		common.Failed(ctx,"添加类别数据失败")
		return
	}
	common.Success(ctx,categoryData)
	return
}