package main

import (
	"downloader-api/controller"
	"downloader-api/enviroment"
	"downloader-api/middleware"
	"downloader-api/security"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	enviroment.SetEnviromentVariables()
	router := gin.Default()
	apiGroup := router.Group("/api", middleware.AuthenticationMiddleware)

	port := fmt.Sprintf(":%v", os.Getenv("port"))
	defer router.Run(port)

	maker, _ := security.NewJwtMaker(os.Getenv("jwtSecret"))
	token, err := maker.CreateToken("a_r_e_k97", time.Minute)
	if err != nil {
		panic(err)
	}
	fmt.Printf("token wygenerowany: %s\n", token)

	apiGroup.GET("/directory/", controller.DirectoryController)
	apiGroup.GET("/download/:filename", controller.DownloadController)
}
