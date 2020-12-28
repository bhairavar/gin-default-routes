package gin_default_routes

import "github.com/gin-gonic/gin"

func Static(route *gin.Engine)  {
	route.Static("/public", "./public")
}