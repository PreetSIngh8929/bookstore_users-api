package app

import (
	// "../bookstore-oauth-go/oauth"

	"github.com/PreetSIngh8929/boookstore_utils-go/logger"
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
