package controller

import (
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type directory struct {
	Filename string    `json:"filename"`
	Size     int64     `json:"size"`
	ModTime  time.Time `json:"modtime"`
}

func DirectoryController(ctx *gin.Context) {
	dirPath := os.Getenv("dirpath")
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	mappedDirectory := mapDirectoryToFiles(files)
	ctx.IndentedJSON(http.StatusOK, mappedDirectory)
}

func mapDirectoryToFiles(files []fs.FileInfo) []directory {
	var directoryList []directory
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		folder := &directory{
			Filename: file.Name(),
			Size:     file.Size(),
			ModTime:  file.ModTime(),
		}
		directoryList = append(directoryList, *folder)
	}
	return directoryList
}
