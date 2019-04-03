package main

import (
	"github.com/gin-gonic/gin"
)

//fp_user
//passW0rd!

func main() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/test", testTheAPI)
		api.GET("/top", topTenNames)
	}
	router.Run()
}
