package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	router := gin.Default()
	router.GET("/healthz", func(ctx *gin.Context) {
		_, _ = bcrypt.GenerateFromPassword([]byte("demo"), bcrypt.MinCost)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	_ = router.Run(":8080")
}
