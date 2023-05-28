package subs_service

import (
	"go_subs_service/internal/model"
	"net/http"

	"context"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=service.go -destination=mocks/mock_SubcribeRepository.go

type SubcribeRepository interface {
	CreateSubscription(ctx context.Context, user1 string, user2 string) (string, error)
	CancelSubcription(ctx context.Context, user string) (string, error)
	GetSubcriptionList(ctx context.Context, user string) ([]interface{}, error)
}

type Service struct {
	subscriptionRepository SubcribeRepository
}

func NewService(subscriptionRepository SubcribeRepository) *Service {
	return &Service{
		subscriptionRepository: subscriptionRepository,
	}
}

func (s *Service) CancelSub(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		paramValue := c.Query("user")
		msg, err := s.subscriptionRepository.CancelSubcription(ctx, paramValue)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"result": msg,
		})
	}
}

func (s *Service) CreateSubs(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var subsModel model.SubsModel
		c.BindJSON(&subsModel)
		id, err := s.subscriptionRepository.CreateSubscription(ctx, subsModel.User, subsModel.SubsUser)

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
