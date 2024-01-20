package repositories

import (
	"context"
)

type StoryRepository interface {
	GetHashtag(ctx context.Context, contentID []int) ([]string, error)
}
