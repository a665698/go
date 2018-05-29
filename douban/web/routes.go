package web

import (
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	route := gin.Default()
	api := route.Group("/api")
	{
		api.GET("/list", GetList)
	}
	return route
}
