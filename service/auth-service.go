package service

import (
	"github.com/mashingan/smapping"
	"github.com/slycreator/shop-for-me/dto"
	"github.com/slycreator/shop-for-me/entity"
	"github.com/slycreator/shop-for-me/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	CreateUser(user dto.RegisterDTO) entity.User
	IsDuplicatePhone(phone string) bool
	IsDuplicateEmail(email string) bool
	VerifyCredential(email string,password string) interface{}
}

type authService struct {
	authRepository repository.AuthRepository
}
//NewAuthService creates a new instance of AuthService
func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return &authService{
		authRepository: authRepo,
	}
}

func (a *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := a.authRepository.CreateUser(userToCreate)
	return res
}

func (a *authService) IsDuplicatePhone(phone string) bool {
	res := a.authRepository.IsDuplicatePhone(phone)
	return !(res.Error == nil)
}

func (a *authService) IsDuplicateEmail(email string) bool {
	res := a.authRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (a *authService) VerifyCredential(email string, password string) interface{} {
	res := a.authRepository.VerifyCredential(email)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (a authService) UpdatePassword()  {
	
}
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}