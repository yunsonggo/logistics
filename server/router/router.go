package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orderserve/common"
	"orderserve/controller/backController"
	"orderserve/controller/webController"
	"orderserve/middleware"
)

func NewRouter() *gin.Engine {
	// 初始化引擎
	app := gin.Default()
	// 设置静态文件
	app.StaticFS("/api/static",http.Dir("./public"))
	// 启用sessions
	common.InitSession(app)
	// 启用cors
	app.Use(middleware.Cors())
	// web前台接口
	//自由访问
	webGroup := app.Group("/api")
	{
		// 发送验证码
		webGroup.POST("/sms/send", webController.SendCode)
		// 手机短信登陆
		webGroup.POST("/sms/login", webController.UserLogin)
		// 获取 城市信息
		webGroup.GET("/cities",webController.GetCityInfo)
		// 获取 首页轮播商家
		webGroup.GET("/profile/shopping",webController.GetProShopping)
		// 获取首页 商家分类
		webGroup.GET("/profile/category",webController.GetProCate)
		// 获取公共 底部导航
		webGroup.GET("/tab/menu",webController.GetTab)
		// 获取首页 展示商家
		webGroup.POST("/profile/supplier",webController.PostProSupplier)
		// 获取 搜索数据
		webGroup.POST("/profile/search",webController.PostSearch)
		// 获取 商家信息
		webGroup.GET("/shop/info",webController.GetSupplierInfo)
		// 获取商家信息 根据分类ID
		webGroup.GET("/profile/shop/cate",webController.GetSupplierByCate)
		// 获取 商家参与服务
		webGroup.POST("/shop/give",webController.PostSupplierGive)
		// 获取 商家参与活动
		webGroup.POST("/shop/sale",webController.PostSupplierSale)
		// 获取 商品信息
		webGroup.POST("/shop/goods",webController.PostStateGoodsAll)
		// 获取 商家内部 商品分类
		webGroup.POST("/shop/shopcate",webController.PostSupplierCatesAll)
		// 获取 商家评论 包括商品
		webGroup.POST("/good/evaluation",webController.PostEvaByGoodId)

		// 需要登陆权限:
		webAuthGroup := webGroup.Group("/auth")
		webAuthGroup.Use(middleware.WebAuth())
		{
			// 获取 用户信息
			webAuthGroup.POST("/user/info",webController.PostUserById)
			// 获取 用户地址
			webAuthGroup.POST("/user/address/list",webController.PostUserAddr)
			// 新增 用户收货地址
			webAuthGroup.POST("/user/address/insertone",webController.PostAddUserAddr)
			// 编辑 用户收货地址
			webAuthGroup.POST("/user/address/editone",webController.PostEditUserAddr)
			// 删除 用户收货地址
			webAuthGroup.POST("/user/address/removeone",webController.PostDelUserAddr)
			// 提交订单
			webAuthGroup.POST("/order/send",webController.PostOrder)
			// 获取用户所有订单
			webAuthGroup.POST("/order/all",webController.PostOrderAll)
		}
	}
	// 后端控制台
	backendGroup := app.Group("/manager")
	{
		// 验证 验证码
		backendGroup.POST("/captcha/verify",backController.VerifyCaptcha)
		// 登录
		backendGroup.POST("/login",backController.PostLogin)
		// 注册用户
		backendGroup.POST("/sign/up",backController.PostSignUp)
		// 后台需要登录权限
		backAuthGroup := backendGroup.Group("/auth")
		backAuthGroup.Use(middleware.BackendAuth())
		{
			// 获取登录用户数据
			backAuthGroup.POST("/manager/info",backController.PostManagerInfo)
			// 工作人员列表
			backAuthGroup.POST("/manager/admin/list",backController.PostManagerList)
			// 工作人员 头像上传
			backAuthGroup.POST("/manager/admin/avatar/upload",backController.PostUploadAvatar)
			// 管理员 编辑 工作人员基本信息
			backAuthGroup.POST("/manager/admin/edit",backController.PostEditManager)
			// 管理员添加工作人员
			backAuthGroup.POST("/manager/admin/add",backController.PostAddManager)
			// 管理员 删除 工作人员账号
			backAuthGroup.POST("/manager/admin/del",backController.PostDelManager)
			// 获取菜单数据
			backAuthGroup.POST("/manager/menu/list",backController.PostMenuAll)
			// 添加菜单
			backAuthGroup.POST("/manager/menu/add",backController.PostMenuAdd)
			// 更新菜单
			backAuthGroup.POST("/manager/menu/edit",backController.PostMenuEdit)
			// 获取商家列表数据
			backAuthGroup.POST("/manager/supp/list",backController.PostSuppList)
			// 更新商家banner上传
			backAuthGroup.POST("/manager/supplier/banner/upload",backController.PostUploadBanner)
			// 更新商家Logo 上传
			backAuthGroup.POST("/manager/supplier/logo/upload",backController.PostUploadLogo)
			// 商家数据更新
			backAuthGroup.POST("/manager/supp/edit",backController.PostSuppEdit)
			// 添加商家banner上传
			backAuthGroup.POST("/manager/addsupplier/banner/upload",backController.PostAddBanner)
			// 添加商家logo上传
			backAuthGroup.POST("/manager/addsupplier/logo/upload",backController.PostAddLogo)
			// 添加商家
			backAuthGroup.POST("/manager/supp/add",backController.PostSuppAdd)
			// 修改商家类别图标
			backAuthGroup.POST("/manager/category/icon/upload",backController.PostEditCateIcon)
			// 更新商家类别数据
			backAuthGroup.POST("/manager/category/edit",backController.PostEditCategory)
			// 添加商家类别数据
			backAuthGroup.POST("/manager/category/add",backController.PostAddCategory)
			// 获取所有服务数据
			backAuthGroup.POST("/manager/give/list",backController.PostGiveList)
			// 修改服务ICON
			backAuthGroup.POST("/manager/give/icon/upload",backController.PostGiveIconUpload)
			// 更新服务数据
			backAuthGroup.POST("/manager/give/edit",backController.PostGiveEdit)
			// 添加服务数据
			backAuthGroup.POST("/manager/give/add",backController.PostGiveAdd)
			// 获取所有活动数据
			backAuthGroup.POST("/manager/salepromote/list",backController.PostSalePromoteList)
			// 修改活动数据
			backAuthGroup.POST("/manager/salepromote/edit",backController.PostSalePromoteEdit)
			// 新增活动数据
			backAuthGroup.POST("/manager/salepromote/add",backController.PostSalePromoteAdd)
			// 所有用户数据
			backAuthGroup.POST("/manager/user/list",backController.PostUserAll)
			// 获取所有订单数据
			backAuthGroup.POST("/manager/order/list",backController.PostOrderAll)
			// 根据订单ID 获取已售商品
			backAuthGroup.POST("/manager/order/salegood",backController.PostGoodsByOrderId)

		}
	}
	return app
}
