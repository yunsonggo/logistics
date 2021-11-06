package middleware

import (
	"github.com/gin-gonic/gin"
	"orderserve/common"
	"orderserve/param/backParam"
	"orderserve/server/backServer"
	"strconv"
)

func BackendAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var ms = backServer.NewManagerServer()
		var jwtParam  backParam.JwtBackendParam
		err := ctx.ShouldBind(&jwtParam)
		if err != nil {
			common.Failed(ctx,"获取验证参数失败")
			ctx.Abort()
			return
		}
		tokenString := jwtParam.ManagerToken
		// admin头像上传参数 // supplier banner logo上传参数 // 服务图标
		if ctx.Request.PostFormValue("manager_token") != "" {
			adminManagerToken := ctx.Request.PostFormValue("manager_token")
			// 管理员头像参数
			_,adminAvatarFileHeader,aErr := ctx.Request.FormFile("admin_avatar_file")
			if adminAvatarFileHeader != nil {
				if aErr != nil {
					common.Failed(ctx,"获取头像参数失败")
					ctx.Abort()
					return
				}
				//设置 更新 admin 头像 旧数据 用于删除 旧图片
				oldAvatarUrl := ctx.Request.PostFormValue("old_avatar_url")
				// 设置 更新admin 头像 管理员ID
				editAvatarAdminId := ctx.Request.PostFormValue("admin_id")
				if oldAvatarUrl != "" {
					ctx.Set("adminOldAvatarUrl",oldAvatarUrl)		// 更新头像的 旧地址
				}
				ctx.Set("editAvatarAdminId",editAvatarAdminId) // 更新头像的管理员ID
				ctx.Set("adminAvatarFileHeader",adminAvatarFileHeader)  // 新的头像文件
			}
			// supp banner参数
			_,supplierBannerFileHeader,suppBannerErr := ctx.Request.FormFile("supp_banner_file")
			if supplierBannerFileHeader != nil {
				if suppBannerErr != nil {
					common.Failed(ctx,"获取banner参数失败")
					ctx.Abort()
					return
				}
				// 设置 更新 banner 旧数据
				oldBannerUrl := ctx.Request.PostFormValue("old_banner_url")
				// 设置 更新banner 商家ID
				editBannerSuppId := ctx.Request.PostFormValue("supp_id")
				if oldBannerUrl != "" {
					ctx.Set("oldBannerUrl",oldBannerUrl)			// 更新banner 的旧地址
				}
				ctx.Set("editBannerSuppId",editBannerSuppId)   // 更新banner 的商家ID
				ctx.Set("bannerFileHeader",supplierBannerFileHeader)    // 新的banner文件
			}
			// logo 参数
			_,supplierLogoFileHeader,suppLogoErr := ctx.Request.FormFile("supp_logo_file")
			if supplierLogoFileHeader != nil {
				if suppLogoErr != nil {
					common.Failed(ctx,"获取logo参数失败")
					ctx.Abort()
					return
				}
				// 设置 更新logo 旧数据
				oldLogoUrl := ctx.Request.PostFormValue("old_logo_url")
				// 设置 更新logo 商家ID
				editLogoSuppId := ctx.Request.PostFormValue("supp_id")
				if oldLogoUrl != "" {
					ctx.Set("oldLogoUrl",oldLogoUrl)
				}
				ctx.Set("editLogoSuppId",editLogoSuppId)
				ctx.Set("logoFileHeader",supplierLogoFileHeader)
			}
			// cateGoryIcon 参数
			_,cateGoryIconFileHeader,cateGoryIconErr := ctx.Request.FormFile("cate_icon_file")
			if cateGoryIconFileHeader != nil {
				if cateGoryIconErr != nil {
					common.Failed(ctx,"获取类别图标参数失败")
					ctx.Abort()
					return
				}
				// 设置 更新category icon 旧数据
				oldCateIconUrl := ctx.Request.PostFormValue("old_cateicon_url")
				// 设置 更新cateId
				editCateId := ctx.Request.PostFormValue("cate_id")
				if oldCateIconUrl != "" {
					ctx.Set("oldCateIconUrl",oldCateIconUrl)
				}
				ctx.Set("editCateId",editCateId)
				ctx.Set("cateGoryIconFileHeader",cateGoryIconFileHeader)
			}
			// give icon 参数
			_,giveIconFileHeader,giveIconErr := ctx.Request.FormFile("give_icon_file")
			if giveIconFileHeader != nil {
				if giveIconErr != nil {
					common.Failed(ctx,"获取服务图标参数失败")
					ctx.Abort()
					return
				}
				// 设置 更新give icon 旧数据
				oldGiveIconUrl := ctx.Request.PostFormValue("old_giveicon_url")
				// 设置 更新cateId
				editGiveId := ctx.Request.PostFormValue("give_id")
				if oldGiveIconUrl != "" {
					ctx.Set("oldGiveIconUrl",oldGiveIconUrl)
				}
				ctx.Set("editGiveId",editGiveId)
				ctx.Set("giveIconFileHeader",giveIconFileHeader)
			}

			// 设置TOKEN
			if tokenString == "" {
				tokenString = adminManagerToken
			}
		}
		// 解析TOKEN
		token,claims,parseErr := common.ParseToken(tokenString)
		if parseErr != nil || !token.Valid {
			common.Failed(ctx,"需要登陆")
			ctx.Abort()
			return
		}
		managerId := claims.UserId
		sessId := strconv.FormatInt(managerId,10)
		id := common.GetSess(ctx,"manager_" + sessId)
		res,resErr := ms.GetManagerById(id.(int64))
		if resErr != nil || res.Id == 0 {
			common.Failed(ctx,"需要注册")
			ctx.Abort()
			return
		}
		res.PassWord = ""
		//接收 添加菜单数据
		addMenuData := jwtParam.AddMenuData
		//接收 后台添加工作人员数据
		managerData := jwtParam.ManagerData
		if managerData.UserName != "" {
			ctx.Set("managerData",managerData)
		}
		//接收 编辑商家数据
		supplierData := jwtParam.SupplierData
		ctx.Set("supplierData",supplierData)
		// 接收 商家分类数据
		cateGoryData := jwtParam.CategoryData
		ctx.Set("categoryData",cateGoryData)
		// 接收 服务数据
		giveData := jwtParam.GiveData
		ctx.Set("giveData",giveData)
		// 接收 活动数据
		promoteData := jwtParam.SalePromoteData
		ctx.Set("promoteData",promoteData)
		// 接收 订单数据
		orderInfo := jwtParam.OrderInfo
		ctx.Set("orderInfo",orderInfo)
		//设置 登录用户数据
		ctx.Set("managerInfo",res)
		//设置 添加菜单数据
		ctx.Set("addMenuData",addMenuData)


		ctx.Next()
	}
}