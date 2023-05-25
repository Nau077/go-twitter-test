package subs_service

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (s *Service) CancelSub(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		paramValue := c.Query("user")
		msg, err := s.subscriptionRepository.CancelSubcription(ctx, paramValue)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
		}

		c.JSON(200, gin.H{
			"result": msg,
		})
	}
}
