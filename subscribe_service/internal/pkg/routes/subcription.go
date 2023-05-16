package routes

import (
	"context"
	"go_subs_service/internal/services/subs_service"

	"github.com/gin-gonic/gin"
)

func NewSubsHTTPHandler(ctx context.Context, r *gin.Engine, s *subs_service.Service) *gin.Engine {
	xapi := r.Group("/")
	api.POST("/CreateSubs", s.CreateSubs(ctx))

	return r
}
