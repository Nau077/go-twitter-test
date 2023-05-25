package subs_service

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (s *Service) GetSubsList(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		paramValue := c.Query("user")
		users, err := s.subscriptionRepository.GetSubcriptionList(ctx, paramValue)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
		}

		c.JSON(200, gin.H{
			"users": users,
		})
	}
}
