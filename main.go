package main

import (
	"downloader-api/controller"
	"downloader-api/enviroment"
	"downloader-api/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	enviroment.SetEnviromentVariables()
	router := gin.Default()
	apiGroup := router.Group("/api", middleware.AuthenticationMiddleware)

	port := fmt.Sprintf(":%v", os.Getenv("port"))
	defer router.Run(port)

	apiGroup.GET("/directory/", controller.DirectoryController)
	apiGroup.GET("/download/:filename", controller.DownloadController)

	router.POST("/login", controller.LoginController)
}
