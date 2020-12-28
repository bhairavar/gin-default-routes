package gin_default_routes

import "github.com/gin-gonic/gin"


func NotFound(route *gin.Engine) {
	route.NoRoute(func(c *gin.Context) {
		res := "Nothing Here!"
		c.AbortWithStatusJSON(404, res)
	})
}

