package service

import (
	"log"

	"github.com/AbdulAffif/hallobumil_dev/api/dto"
	"github.com/AbdulAffif/hallobumil_dev/api/entity"
	"github.com/AbdulAffif/hallobumil_dev/api/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.Result
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.Result {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	UpdatedUser := service.userRepository.UpdateUser(userToUpdate)
	return UpdatedUser
}

