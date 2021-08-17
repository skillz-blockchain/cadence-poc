package server

import (
	"cadence-poc/cadence/workers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	V1             *gin.RouterGroup
	WorkflowClient workers.CadenceAdapter
}

func NewServer(wfc workers.CadenceAdapter) *Server {
	s := &Server{Engine: gin.Default()}

	// Init Cadence
	s.WorkflowClient = wfc

	// Init API's routing
	s.V1 = s.Group("/v1")
	s.V1.Use(func(c *gin.Context) {
		c.Set("wfc", s.WorkflowClient)
	})
	InitRouter(s.V1)
	return s
}
