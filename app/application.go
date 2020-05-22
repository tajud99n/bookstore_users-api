package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tajud99n/bookstore_utils-go/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapURLS()

	logger.Info("About to start the application")
	router.Run(":8080")
}
