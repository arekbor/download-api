package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadController(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	err = ctx.SaveUploadedFile(file, os.Getenv("dirupload")+file.Filename)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(os.Getenv("dirupload") + file.Filename)
	ctx.String(http.StatusOK, fmt.Sprintf("%s successfully uploaded!", file.Filename))
}
