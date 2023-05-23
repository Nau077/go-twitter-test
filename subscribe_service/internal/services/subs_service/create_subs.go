package subs_service

import (
	"context"
	"go_subs_service/internal/model"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateSubs(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var subsModel model.SubsModel
		c.BindJSON(&subsModel)
		id, err := s.subscriptionRepository.CreateSubscription(ctx, subsModel.Id, subsModel.SubsId)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err,
			})
		}

		c.JSON(200, gin.H{
			"id": id,
		})
	}
}
