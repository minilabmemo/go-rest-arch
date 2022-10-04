package service

import "github.com/minilabmemo/go-rest-arch/internal/models"

func Users() ([]models.User, error) {

	us := make([]models.User, 0)
	us = append(us, models.User{Name: "A"})

	return us, nil

}
