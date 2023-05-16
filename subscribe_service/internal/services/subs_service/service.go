package subs_service

import "go_subs_service/subscribe_service/internal/repository"

type Service struct {
	subscriptionRepository repository.SubcribeRepository
}

func NewService(subscriptionRepository repository.Repository) *Service {
	return &Service{
		subscriptionRepository: subscriptionRepository,
	}
}
