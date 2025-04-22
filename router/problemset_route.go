package router

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller/problemset"
)

func ProblemSetRoute(router *gin.Engine) {
	routerGroup := router.Group("api/v1/problemsets")
	{
		routerGroup.GET("/lists", problemset.GetProblemsetsByEvent)
	}
}
