package routes

import (
	"test/handlers"
	"test/middlewares"

	"github.com/gin-gonic/gin"
)

type RouteType string

const (
	GET  RouteType = "GET"
	POST RouteType = "POST"
	PUT  RouteType = "PUT"
)

type Route struct {
	addr        string
	routeType   RouteType
	handler     func(c *gin.Context)
	middlewares []gin.HandlerFunc
}

func SetupRoutes(router *gin.Engine) {
	routes := []Route{
		{
			addr:        "/sendTransfers",
			routeType:   POST,
			handler:     handlers.SendHandlers,
			middlewares: []gin.HandlerFunc{middlewares.ParseData(), middlewares.CheckIfBalanceIsOk()},
		},
	}

	for _, route := range routes {
		var handlers []gin.HandlerFunc

		if route.middlewares != nil {
			handlers = append(handlers, route.middlewares...)
		}

		handlers = append(handlers, route.handler)

		switch route.routeType {
		case "GET":
			router.GET(route.addr, handlers...)
		case "POST":
			router.POST(route.addr, handlers...)
		case "PUT":
			router.PUT(route.addr, handlers...)
		}
	}

}
