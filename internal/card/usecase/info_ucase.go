package usecase

import (
	"time"

	"github.com/minilabmemo/go-rest-arch/internal"
	"github.com/minilabmemo/go-rest-arch/internal/config"
	"github.com/minilabmemo/go-rest-arch/internal/models"
)

type infoUsecase struct {
	contextTimeout time.Duration
}

// NewInfoUsecase 會產生一個新的 infoUsecase 物件代表 models.InfoUsecase interface[介面]
func NewInfoUsecase(timeout time.Duration) models.InfoUsecase {
	return &infoUsecase{
		contextTimeout: timeout,
	}
}

//implement info usecase logic , ginrouter wrapper 'Usecase' for handle setting
//實作InfoUsecase[介面]裡面的方法， ginrouter.NewInfoHandler 的參數會需要它
func (*infoUsecase) Fetch(config.CofigDefinition) (models.Info, error) {
	return models.Info{Name: config.ConfigData.Service.Name, Version: internal.Version}, nil

}
