package routers

import "github.com/gin-gonic/gin"

type Router interface {
	Register(engine *gin.Engine)
}