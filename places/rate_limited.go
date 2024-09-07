package places

import (
	"context"

	"golang.org/x/time/rate"
)

// RateLimitedService is a Service that limits the rate at which it calls another Service.
type RateLimitedService struct {
	limiter *rate.Limiter
	next    Service
}

// NewRateLimitedService creates a new RateLimitedService that "wraps" the given Service and limits the rate at which it calls it.
func NewRateLimitedService(limit rate.Limit, next Service) *RateLimitedService {
	return &RateLimitedService{
		limiter: rate.NewLimiter(limit, 1),
		next:    next,
	}
}

// GetPlace calls the wrapped Service's GetPlace method, subject to rate limiting.
func (s *RateLimitedService) GetPlace(ctx context.Context, id string) (*Place, error) {
	if err := s.limiter.Wait(ctx); err != nil {
		return nil, err
	}
	return s.next.GetPlace(ctx, id)
}
