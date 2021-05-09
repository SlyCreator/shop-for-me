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
	IsPhoneInDB(email string) bool
	IsEmailInDB(email string) bool
	VerifyCredential(email string,password string) interface{}
	CreateResetCode(passwordReset dto.PasswordResetDTO) entity.PasswordReset
	VerifyResetToken(email string,token string) bool
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

func (a *authService) IsPhoneInDB(phone string) bool {
	res := a.authRepository.FindPhone(phone)
	return !(res.Error != nil)//returns true when res is found


}

func (a *authService) IsEmailInDB(email string) bool {
	res := a.authRepository.FindEmail(email)
	return !(res.Error != nil) //returns true when email is not found and negate it to false
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

func (a authService) CreateResetCode(passwordReset dto.PasswordResetDTO) entity.PasswordReset{
	tokenToCreate := entity.PasswordReset{}
	//read about smapping and how it work
	err := smapping.FillStruct(&tokenToCreate,smapping.MapFields(&passwordReset))
	if err != nil {
		log.Fatalf("Failed to map")
	}
	res := a.authRepository.CreateResetCode(tokenToCreate)
	return res
}

func (a *authService) VerifyResetToken(email string,token string) bool {
	res := a.authRepository.VerifyResetToken(email,token)
	return res.Error == nil

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