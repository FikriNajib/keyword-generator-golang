package repositories

import (
	"context"
)

type VideosRepository interface {
	GetKeyword(ctx context.Context, contentID []int) ([]string, error)
}
