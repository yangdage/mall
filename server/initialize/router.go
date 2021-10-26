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

		// 品牌管理API
		web.POST("/brand/create", api.WebCreateBrand)
		web.DELETE("/brand/delete", api.WebDeleteBrand)
		web.PUT("/brand/update", api.WebUpdateBrand)
		web.GET("/brand/list", api.WebGetBrandList)
		web.GET("/brand/option", api.WebGetBrandOption)

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
		web.GET("/product/info", api.WebGetProductInfo)
		web.GET("/product/list", api.WebGetProductList)

		// 订单管理API
		web.DELETE("/order/delete", api.WebDeleteOrder)
		web.PUT("/order/update", api.WebUpdateOrder)
		web.GET("/order/list", api.WebGetOrderList)
		web.GET("/order/detail", api.WebGetOrderDetail)
		web.POST("/order/set/save", api.WebSaveOrderSet)
		web.GET("/order/set/info", api.WebGetOrderSetInfo)

		// 用户管理API
		web.DELETE("/user/delete", api.WebDeleteUser)
		web.PUT("/user/update", api.WebUpdateUser)
		web.GET("/user/list", api.WebGetUserList)

		// 数据统计API
		web.GET("/statistics/info", api.WebGetStatisticsInfo)
		web.GET("/today/order/info", api.WebGetTodayOrderInfo)
		web.GET("/week/order/info", api.WebGetWeekInfo)
	}

	// 微信小程序用户接口
	app := engine.Group("/app")
	{
		app.POST("/login", api.AppUserLogin)

		app.GET("/product/list", api.AppGetProductList)
		app.GET("/product/detail", api.AppGetProductDetail)

		app.GET("/cart/add", api.AppAddCart)
		app.GET("/cart/delete", api.AppDeleteCart)
		app.GET("/cart/clear", api.AppClearCart)
		app.GET("/cart/info", api.AppGetCartInfo)

		app.POST("/order/create", api.AppCreateOrder)
		app.GET("/order/list", api.AppGetOrderList)

		app.POST("/address/add", api.AppAddAddress)
		app.GET("/address/delete", api.AppDeleteAddress)
		app.GET("/address/update", api.AppUpdateAddress)
		app.GET("/address/info", api.AppGetAddressUpdateInfo)
		app.GET("/address/list", api.AppGetAddressList)

		app.POST("/collection/add", api.AppAddCollection)
		app.GET("/collection/delete", api.AppDeleteCollection)
		app.GET("/collection/list", api.AppGetCollectionList)

		app.GET("/footmark/add", api.AppAddFootmark)
		app.GET("/footmark/delete", api.AppDeleteFootmark)
		app.GET("/footmark/list", api.AppGetFootmarknList)
	}

	// 启动、监听端口
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("server start error: %s", err)
	}
}