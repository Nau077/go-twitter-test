package subcription

import "context"

func (s *Service) CreateSubs(ctx context.Context) (int64, error) {
	id, err := s.subscriptionRepository.CreateSubcription(ctx)
	if err != nil {
		return 0, err
	}

	return id, nil
}
