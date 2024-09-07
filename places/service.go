package places

import "context"

// Service is a service that can supply Google Place API objects given a place ID.
type Service interface {
	GetPlace(context.Context, string) (*Place, error)
}
