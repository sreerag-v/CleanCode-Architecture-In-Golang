package http

import (
	"github.com/gin-gonic/gin"
	"my-clean.go/pkg/api/handler"
)

type ServerHTTP struct {
	engine *gin.Engine
}
func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()
	// Auth middleware
	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
