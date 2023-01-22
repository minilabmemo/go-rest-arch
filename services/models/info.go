package models

// Service Info
type Info struct {
	Name       string `json:"name" validate:"required"`
	Version    string `json:"version" validate:"required"`
	StartupMsg string `json:"startupMsg" validate:"required"`
}

type InfoUpdate struct {
	Name       string `json:"name" validate:"required"`
	StartupMsg string `json:"startupMsg" validate:"required"`
}

// InfoUsecase 代表 the Info's usecases
type InfoUsecase interface {
	GetInfo() (Info, error)
	Update(*InfoUpdate) error
}
