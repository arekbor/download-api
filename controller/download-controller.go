package controller

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func DownloadController(ctx *gin.Context) {
	paramFilename := ctx.Param("filename")
	dirPath := os.Getenv("dirpath")
	_, err := ioutil.ReadDir(dirPath)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	targetPath := filepath.Join(dirPath, paramFilename)
	if _, err := os.Stat(targetPath); errors.Is(err, os.ErrNotExist) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+paramFilename)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(targetPath)
}
