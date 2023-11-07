package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/logger"
	"web_app/settings"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("version: %s", settings.Conf.Version))
	})
	return r
}
