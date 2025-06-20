package api

import (
	"fmt"
	"github.com/santirufiner/slicerwise/internal/infrastructure/api/routers"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run(port int, l *logrus.Logger, routers ...routers.Router) error {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	for _, r := range routers {
		r.Register(router)
	}

	l.Infof("starting server on port %d", port)
	return router.Run(fmt.Sprintf(":%d", port))
}
