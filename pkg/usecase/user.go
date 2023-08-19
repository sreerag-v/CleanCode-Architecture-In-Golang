package usecase
import (

	interfaces "my-clean.go/pkg/repository/interface"
	services "my-clean.go/pkg/usecase/interface"
)


type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}
