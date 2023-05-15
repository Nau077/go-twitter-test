package routes

import (
	"time"

	"github.com/Nau077/cassandra-golang-sv/internal/services/post"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func NewPostHTTPHandler(r *gin.Engine, s *post.Service) {
	// handler := &HTTPPost{}

	api := r.Group("/post")
	api.GET("/", s.GetRandom)
	api.POST("/add", timeout.New(
		timeout.WithTimeout(4000*time.Millisecond),
		// timeout.WithHandler(handler.InsertData),
		// timeout.WithResponse(testResponse),
	))
}
