package subcription

import "go-twitter-test/subscribe_service/internal/repository"

type Service struct {
	subscriptionRepository repository.SubcribeRepository
}

func NewService(subscriptionRepository repository.Repository) *Service {
	return &Service{
		subscriptionRepository: subscriptionRepository,
	}
}
