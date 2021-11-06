package webController

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq"
	"orderserve/common"
	"os"
)

func GetCityInfo(ctx *gin.Context) {
	filePath,_ := os.Getwd()
	fileName := "/model/citiesData.json"
	path := filePath + fileName
	jsonData := gojsonq.New().File(path).Get()
	common.Success(ctx,jsonData)
	return
}
