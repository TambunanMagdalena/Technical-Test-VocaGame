package routes

import (
	controller "template-go/app/controllers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func ConfigureRouter(e *echo.Echo, controller *controller.Main) {

	v1 := e.Group("/v1")
	{
		v1.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	masterItem := v1.Group("/master-item")
	{
		masterItem.POST("", controller.Item.Create)
		masterItem.GET("", controller.Item.Get)
		masterItem.PATCH("", controller.Item.Update)
		masterItem.DELETE("", controller.Item.Delete)
		masterItem.GET("/export", controller.Item.Export)
	}
}