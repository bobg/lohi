package places

import (
	"context"

	"golang.org/x/time/rate"
)

type RateLimitedService struct {
	limiter *rate.Limiter
	next    Service
}

func NewRateLimitedService(limit rate.Limit, next Service) *RateLimitedService {
	return &RateLimitedService{
		limiter: rate.NewLimiter(limit, 1),
		next:    next,
	}
}

func (s *RateLimitedService) GetPlace(ctx context.Context, id string) (*Place, error) {
	if err := s.limiter.Wait(ctx); err != nil {
		return nil, err
	}
	return s.next.GetPlace(ctx, id)
}
