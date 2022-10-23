package controller

import (
	"errors"
	"fmt"
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
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	targetPath := filepath.Join(dirPath, paramFilename)
	if _, err := os.Stat(targetPath); errors.Is(err, os.ErrNotExist) {
		ctx.String(http.StatusBadRequest, fmt.Sprintln(err))
		return
	}
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+paramFilename)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(targetPath)
}
