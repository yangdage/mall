package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall.com/api"
	"mall.com/global"
	"mall.com/middleware"
)

func Router() {

	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	// 静态资源请求映射
	engine.Static("/image", global.Config.Upload.SavePath)

	// 后台管理员前端接口
	web := engine.Group("/web")

	{
		// 用户登录API
		web.GET("/captcha", api.WebGetCaptcha)
		web.POST("/login", api.WebUserLogin)

		// 开启JWT认证,以下接口需要认证成功才能访问
		web.Use(middleware.JwtAuth())

		// 文件上传API
		web.POST("/upload", api.WebFileUpload)

		// 类目管理API
		web.POST("/category/create", api.WebCreateCategory)
		web.DELETE("/category/delete", api.WebDeleteCategory)
		web.PUT("/category/update", api.WebUpdateCategory)
		web.GET("/category/list", api.WebGetCategoryList)
		web.GET("/category/option", api.WebGetCategoryOption)

		// 商品管理API
		web.POST("/product/create", api.WebCreateProduct)
		web.DELETE("/product/delete", api.WebDeleteProduct)
		web.PUT("/product/update", api.WebUpdateProduct)
		web.PUT("/product/status/update", api.WebUpdateProductStatus)
		web.GET("/product/info", api.WebGetProductInfo)
		web.GET("/product/list", api.WebGetProductList)

		// 订单管理API
		web.DELETE("/order/delete", api.WebDeleteOrder)
		web.PUT("/order/update", api.WebUpdateOrder)
		web.GET("/order/list", api.WebGetOrderList)
		web.GET("/order/detail", api.WebGetOrderDetail)

		// 数据统计API
		web.GET("/data/overview/info", api.WebGetDataOverviewInfo)
		web.GET("/today/order/data/info", api.WebGetTodayOrderDataInfo)
		web.GET("/week/data/info", api.WebGetWeekDataInfo)
	}

	// 微信小程序用户接口
	app := engine.Group("/app")

	{
		app.POST("/login", api.AppUserLogin)

		// 商品API
		app.GET("/product/list", api.AppGetProductList)
		app.GET("/product/search", api.AppGetProductSearchList)
		app.GET("/product/detail", api.AppGetProductDetail)

		// 分类API
		app.GET("/category/option", api.AppGetCategoryOption)

		// 购物车API
		app.POST("/cart/add", api.AppAddCart)
		app.DELETE("/cart/delete", api.AppDeleteCart)
		app.DELETE("/cart/clear", api.AppClearCart)
		app.GET("/cart/info", api.AppGetCartInfo)

		// 商品订单API
		app.POST("/order/create", api.AppCreateOrder)
		app.GET("/order/list", api.AppGetOrderList)

		// 收货地址API
		app.POST("/address/add", api.AppAddAddress)
		app.DELETE("/address/delete", api.AppDeleteAddress)
		app.PUT("/address/update", api.AppUpdateAddress)
		app.GET("/address/info", api.AppGetAddressUpdateInfo)
		app.GET("/address/list", api.AppGetAddressList)

		// 商品收藏API
		app.POST("/collection/add", api.AppAddCollection)
		app.DELETE("/collection/delete", api.AppDeleteCollection)
		app.GET("/collection/list", api.AppGetCollectionList)
		//
		// 商品浏览记录API
		app.POST("/browse/save", api.AppSaveBrowseRecord)
		app.DELETE("/browse/delete", api.AppDeleteBrowseRecord)
		app.GET("/browse/list", api.AppGetBrowseRecordList)
	}

	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}