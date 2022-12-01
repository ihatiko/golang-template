package files

import "context"

type Service interface {
	SaveImage(ctx context.Context) (string, error)
}
