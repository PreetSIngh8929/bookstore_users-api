package app

import (
	// "../bookstore-oauth-go/oauth"
	"github.com/PreetSIngh8929/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("about to start the application")
	router.Run("localhost:8081")

}
