package models

import (
	"context"
)

type Card struct {
	ID         string `bson:"_id" json:"id" validate:"required" example:"63c39d754be73bb880adb763" ` //mongo _id link to id
	CardUpdate `bson:",inline" json:",inline"`
}

type CardUpdate struct {
	Title   string `json:"title" validate:"required" example:"learn golang" `
	Content string `json:"content"  example:"finished CRUD APIs"`
}

// CardUsecase represent the card's usecases
type CardUsecase interface {
	Fetch(ctx context.Context) ([]Card, string, error)
	// GetByID(ctx context.Context, id int64) (Card, error) //TODO
	Update(ctx context.Context, id string, card *CardUpdate) (string, error)
	Store(ctx context.Context, card *CardUpdate) (string, error)
	Delete(ctx context.Context, id string) (string, error)
}

// // CardeRpository represent the card's repository operations
type CardRpository interface {
	//FetchByCursor(ctx context.Context, cursor string, num int64) (res []Card, nextCursor string, err error) //TODO
	//FetchByPage(ctx context.Context, cursor string, num int64) (res []Card, nextCursor string, err error) //TODO
	FetchAll(ctx context.Context) (res []Card, total string, err error)
	Update(ctx context.Context, id string, card *CardUpdate) (string, error)
	Store(ctx context.Context, card *CardUpdate) (string, error)
	Delete(ctx context.Context, id string) (string, error)
}
