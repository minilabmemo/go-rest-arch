package usecase

import (
	"context"

	"github.com/minilabmemo/go-rest-arch/services/config"
	"github.com/minilabmemo/go-rest-arch/services/models"
	"github.com/pkg/errors"
)

type cardUsecase struct {
	config        config.CofigDefinition
	cardRpository models.CardRpository
}

// NewcardUsecase 會產生一個新的 cardUsecase 物件代表 models.cardUsecase interface[介面]
func NewCardUsecase(config config.CofigDefinition, cr models.CardRpository) models.CardUsecase {
	return &cardUsecase{
		config:        config,
		cardRpository: cr,
	}
}

//implement CardUsecase usecase logic ,
func (cu *cardUsecase) Fetch(ctx context.Context) ([]models.Card, string, error) {

	return cu.cardRpository.FetchAll(ctx)

}

func (cu *cardUsecase) Update(ctx context.Context, id string, card *models.CardUpdate) (string, error) {

	return cu.cardRpository.Update(ctx, id, card)
}

func (cu *cardUsecase) Store(ctx context.Context, card *models.CardUpdate) (string, error) {

	if card.Title == "" {
		return "", errors.Errorf("title is empty.")
	}

	return cu.cardRpository.Store(ctx, card)
}

func (cu *cardUsecase) Delete(ctx context.Context, id string) (string, error) {

	return cu.cardRpository.Delete(ctx, id)
}
