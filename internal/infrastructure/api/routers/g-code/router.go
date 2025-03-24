package gcode

import (
	"github.com/gin-gonic/gin"
	"github.com/santirufiner/slicerwise/internal/domain/actions/gcode_parser"
	v1 "github.com/santirufiner/slicerwise/internal/infrastructure/api/routers/g-code/v1"
	"github.com/santirufiner/slicerwise/pkg/logger/middleware/logger"
	"github.com/santirufiner/slicerwise/pkg/logger/middleware/logging"
	"github.com/sirupsen/logrus"
)

type Router struct {
	l                    *logrus.Logger
	parseMetadataHandler *gcode_parser.ParseMetadataHandler
}

func NewRouter(
	l *logrus.Logger,
	parseMetadataHandler *gcode_parser.ParseMetadataHandler,
) *Router {
	return &Router{
		l:                    l,
		parseMetadataHandler: parseMetadataHandler,
	}
}

func (router *Router) Register(engine *gin.Engine) {
	router.v1(engine)
}

func (router *Router) v1(engine *gin.Engine) {
	api := engine.Group("api/v1/slicer-wise/:user_id/g-code")
	api.Use(logger.NewLoggerMiddleware(router.l).Middleware(), logging.Logger(router.l))

	api.POST("/parse", v1.ParseMetadata(router.parseMetadataHandler, router.l))
}
