package middleware

import (
	"downloader-api/security"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(ctx *gin.Context) {
	var maker security.Maker
	token := ctx.GetHeader("Authorization")
	maker, _ = security.NewJwtMaker(os.Getenv("jwtSecret"))
	_, err := maker.VerifyToken(token)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	ctx.Next()
}
