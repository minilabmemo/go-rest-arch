package models

import (
	"context"
)

// Card ...
type Card struct {
	ID      int64  `json:"id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// ArticleUsecase represent the article's usecases
type CardUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Card, string, error)
	// GetByID(ctx context.Context, id int64) (Card, error)
	// Update(ctx context.Context, ar *Article) error
	// GetByTitle(ctx context.Context, title string) (Article, error)
	// Store(context.Context, *Article) error
	// Delete(ctx context.Context, id int64) error
}

// // ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Card, nextCursor string, err error)
	// GetByID(ctx context.Context, id int64) (Article, error)
	// GetByTitle(ctx context.Context, title string) (Article, error)
	// Update(ctx context.Context, ar *Article) error
	// Store(ctx context.Context, a *Article) error
	// Delete(ctx context.Context, id int64) error
}
