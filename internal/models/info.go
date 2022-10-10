package models

// Service Info
type Info struct {
	Name    string `json:"name" validate:"required"`
	Version string `json:"version" validate:"required"`
}

// InfoUsecase 代表 the Info's usecases
type InfoUsecase interface {
	GetInfo() (Info, error)
	Update(info *Info) error
}
