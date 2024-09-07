package places

import "context"

type Service interface {
	GetPlace(context.Context, string) (*Place, error)
}
