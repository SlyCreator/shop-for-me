package service

import "github.com/slycreator/shop-for-me/controllers/repository"

type UserService interface {

}

type userService struct {
	userRepository repository.UserRepository
}