package controller

import (
	"downloader-api/security"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginController(ctx *gin.Context) {
	var user user
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if !verifyUser(user.Username, user.Password) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "wrong pwd or login")
		return
	}
	token, err := getJWTToken(user.Username, time.Hour)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, token)
}

func verifyUser(username string, password string) bool {
	return username == os.Getenv("adminlogin") && password == os.Getenv("adminpwd")
}

func getJWTToken(username string, duration time.Duration) (string, error) {
	maker, _ := security.NewJwtMaker(os.Getenv("jwtSecret"))
	token, err := maker.CreateToken(username, duration)
	if err != nil {
		return "", err
	}
	return token, nil
}
