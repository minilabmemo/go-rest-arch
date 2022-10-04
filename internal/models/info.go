package models

import "github.com/minilabmemo/go-rest-arch/internal/config"

// Card ...
type Info struct {
	Name    string `json:"name" validate:"required"`
	Version string `json:"version" validate:"required"`
}

// ArticleUsecase represent the article's usecases
type InfoUsecase interface {
	Fetch(config.CofigDefinition) (Info, error)
	// GetByID(ctx context.Context, id int64) (Card, error)
	// Update(ctx context.Context, ar *Article) error
	// GetByTitle(ctx context.Context, title string) (Article, error)
	// Store(context.Context, *Article) error
	// Delete(ctx context.Context, id int64) error
}
