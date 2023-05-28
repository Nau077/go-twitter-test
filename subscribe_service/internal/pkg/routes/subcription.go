package routes

import (
	"context"
	"go_subs_service/internal/services/subs_service"

	"github.com/gin-gonic/gin"
)

func NewSubsHTTPHandler(ctx context.Context, r *gin.Engine, s *subs_service.Service) *gin.Engine {
	api := r.Group("/")
	api.POST("/subs", s.CreateSubs(ctx))
	api.GET("/subsList", s.GetSubsList(ctx))
	api.DELETE("/subs", s.CancelSub(ctx))

	return r
}
