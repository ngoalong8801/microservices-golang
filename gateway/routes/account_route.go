package routes

import (
	"gateway-service/controller"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type RouteParams struct {
	fx.In
	Route             *gin.Engine
	AccountController *controller.AccountController
}

func InitAccountRoute(params RouteParams) {

	route := params.Route
	accountController := params.AccountController

	accountRoute := route.Group("/users")
	{
		accountRoute.GET("/:id", accountController.GetUser)
	}
}
