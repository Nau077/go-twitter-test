package subs_service

import (
	"go_subs_service/internal/repository"
)

type Service struct {
	subscriptionRepository repository.SubcribeRepository
}

func NewService(subscriptionRepository repository.SubcribeRepository) *Service {
	return &Service{
		subscriptionRepository: subscriptionRepository,
	}
}
