package repository

import (
	"gorm.io/gorm"
	interfaces "my-clean.go/pkg/repository/interface"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}


