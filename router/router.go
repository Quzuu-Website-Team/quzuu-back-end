package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/controller"
)

func StartService() {
	router := gin.Default()
	router.GET("/", controller.HomeController)

	AuthRoute(router)
	UserRoute(router)
	EmailRoute(router)
	OptionsRoute(router)
	EventRoute(router)
	err := router.Run(config.TCP_ADDRESS)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
