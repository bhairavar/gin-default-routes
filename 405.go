package gin_default_routes

import "github.com/gin-gonic/gin"

func NoMethods(route *gin.Engine){
	route.NoMethod(func(c *gin.Context) {
		res := "We are yet to build that space ship."
		c.AbortWithStatusJSON(405, res)
	})
}


