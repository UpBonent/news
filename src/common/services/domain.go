package services

import "context"

type ArticleRepository interface {
	Insert(ctx context.Context) error
}
