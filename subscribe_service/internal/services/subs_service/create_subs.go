package subs_service

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateSubs(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := s.subscriptionRepository.CreateSubcription(ctx)
		if err != nil {
			return 0, err
		}

		c.JSON(200, gin.H{
			"id": id,
		})
	}
}
