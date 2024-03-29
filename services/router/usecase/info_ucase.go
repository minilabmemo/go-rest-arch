package usecase

import (
	"github.com/pkg/errors"

	internal "github.com/minilabmemo/go-rest-arch/services"
	"github.com/minilabmemo/go-rest-arch/services/config"
	"github.com/minilabmemo/go-rest-arch/services/models"
)

type infoUsecase struct {
	config config.CofigDefinition
}

// NewInfoUsecase 會產生一個新的 infoUsecase 物件代表 models.InfoUsecase interface[介面]
func NewInfoUsecase(config config.CofigDefinition) models.InfoUsecase {
	return &infoUsecase{
		config: config,
	}
}

//implement info usecase logic , In e.g.: ginrouter wrapper 'Usecase' struct for handle
//實作InfoUsecase[介面]裡面的方法， 範例中delevery 的ginrouter.NewInfoHandler 的參數會需要它
func (*infoUsecase) GetInfo() (models.Info, error) {
	return models.Info{Name: config.ConfigData.Service.Name, Version: internal.Version, StartupMsg: config.ConfigData.Service.StartupMsg}, nil

}

func (*infoUsecase) Update(body *models.InfoUpdate) error {
	if body.Name == "" {
		return errors.Errorf("no Name")
	}
	config.ConfigData.Service.Name = body.Name
	config.ConfigData.Service.StartupMsg = body.StartupMsg

	return nil
}
