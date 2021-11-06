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

var suppServer = backServer.NewSuppServer()

// 获取商家列表数据
func PostSuppList(ctx *gin.Context) {
	res,resErr := suppServer.GetSuppDataAll()
	if resErr != nil {
		common.Failed(ctx,"获取商家数据失败")
		return
	}
	common.Success(ctx,res)
	return
}
// 商家banner上传
func PostUploadBanner(ctx *gin.Context) {
	bannerFileHeader,_ := ctx.Get("bannerFileHeader")
	fileHeader := bannerFileHeader.(*multipart.FileHeader)
	//随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(fileHeader.Filename)
	fileName := "upload/supp/banner/" + "banner" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	saveFileErr := ctx.SaveUploadedFile(fileHeader,filePath)
	if saveFileErr != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"path":"",
		})
		return
	}
	// 如果存在旧banner 则删除旧banner
	oldBannerUrl,_ := ctx.Get("oldBannerUrl")
	if oldBannerUrl != nil {
		oldUrl := oldBannerUrl.(string)
		oldFilePath := "public/" + oldUrl
		_ = os.Remove(oldFilePath)
	}
	// 提前更新数据库
	editBannerSuppId,_ := ctx.Get("editBannerSuppId")
	suppId := editBannerSuppId.(string)
	id,_ := strconv.ParseInt(suppId,10,64)
	// 根据ID 更新商家banner
	editBannerErr := suppServer.EditSuppBanner(id,fileName)
	if editBannerErr != nil {
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
// 添加商家banner上传
func PostAddBanner(ctx *gin.Context) {
	bannerFileHeader,_ := ctx.Get("bannerFileHeader")
	fileHeader := bannerFileHeader.(*multipart.FileHeader)
	//随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(fileHeader.Filename)
	fileName := "upload/supp/banner/" + "banner" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	saveFileErr := ctx.SaveUploadedFile(fileHeader,filePath)
	if saveFileErr != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"path":"",
		})
		return
	}
	// 如果存在旧banner 则删除旧banner
	oldBannerUrl,_ := ctx.Get("oldBannerUrl")
	if oldBannerUrl != nil {
		oldUrl := oldBannerUrl.(string)
		oldFilePath := "public/" + oldUrl
		_ = os.Remove(oldFilePath)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"path":fileName,
	})
	return
}
// 商家Logo 上传
func PostUploadLogo(ctx *gin.Context) {
	logoFileHeader,_ := ctx.Get("logoFileHeader")
	fileHeader := logoFileHeader.(*multipart.FileHeader)
	//随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(fileHeader.Filename)
	fileName := "upload/supp/logo/" + "logo" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	saveFileErr := ctx.SaveUploadedFile(fileHeader,filePath)
	if saveFileErr != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"path":"",
		})
		return
	}
	// 如果存在旧logo  则删除旧logo
	oldLogoUrl,_ := ctx.Get("oldLogoUrl")
	if oldLogoUrl != nil {
		oldUrl := oldLogoUrl.(string)
		oldFilePath := "public/" + oldUrl
		_ = os.Remove(oldFilePath)
	}
	// 提前更新数据库
	editLogoSuppId,_ := ctx.Get("editLogoSuppId")
	suppId := editLogoSuppId.(string)
	id,_ := strconv.ParseInt(suppId,10,64)
	// 根据ID 更新商家logo
	editLogoErr := suppServer.EditSuppLogo(id,fileName)
	if editLogoErr != nil {
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
// 添加商家logo上传
func PostAddLogo(ctx *gin.Context) {
	logoFileHeader,_ := ctx.Get("logoFileHeader")
	fileHeader := logoFileHeader.(*multipart.FileHeader)
	//随机数
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	extString := filepath.Ext(fileHeader.Filename)
	fileName := "upload/supp/logo/" + "logo" + strconv.FormatInt(time.Now().Unix(),10) + code + extString
	filePath := "public/" + fileName
	saveFileErr := ctx.SaveUploadedFile(fileHeader,filePath)
	if saveFileErr != nil {
		ctx.JSON(http.StatusOK,gin.H{
			"path":"",
		})
		return
	}
	// 如果存在旧logo  则删除旧logo
	oldLogoUrl,_ := ctx.Get("oldLogoUrl")
	if oldLogoUrl != nil {
		oldUrl := oldLogoUrl.(string)
		oldFilePath := "public/" + oldUrl
		_ = os.Remove(oldFilePath)
	}
	ctx.JSON(http.StatusOK,gin.H{
		"path":fileName,
	})
	return
}
// 商家数据更新
func PostSuppEdit(ctx *gin.Context) {
	supplierData,hasData := ctx.Get("supplierData")
	if !hasData {
		common.Failed(ctx,"获取商家数据失败")
		return
	}
	supplierInfo := supplierData.(model.Supplier)
	id := supplierInfo.Id
	// 调用更新数据
	err := suppServer.EditSuppById(id,supplierInfo)
	if err != nil {
		common.Failed(ctx,"更新商家数据失败")
		return
	}
	common.Success(ctx,supplierInfo)
	return
}
// 添加商家
func PostSuppAdd(ctx *gin.Context) {
	supplierData,hasData := ctx.Get("supplierData")
	if !hasData {
		common.Failed(ctx,"获取商家数据失败")
		return
	}
	supplierInfo := supplierData.(model.Supplier)

	// 调用添加数据
	err := suppServer.AddSuppOne(supplierInfo)
	if err != nil {
		common.Failed(ctx,"更新商家数据失败")
		return
	}
	common.Success(ctx,supplierInfo)
	return
}