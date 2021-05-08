package repository

import (
	"math/rand"
	//"crypto/rand"
	"github.com/slycreator/shop-for-me/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"

)

type AuthRepository interface {
	CreateUser(user entity.User) entity.User
	FindPhone(phone string) (tx *gorm.DB)
	FindEmail(phone string) (tx *gorm.DB)
	VerifyCredential(phone string) interface{}
	CreateResetCode(passwordReset entity.PasswordReset) entity.PasswordReset
	UpdatePassword(user entity.User) entity.User
}

type authConnection struct {
	connection *gorm.DB
}

func NewAuthRepository(db *gorm.DB)  AuthRepository{
	return &authConnection{
		connection: db,
	}
}

func (db *authConnection) CreateUser(user entity.User)  entity.User{
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *authConnection) UpdatePassword(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

//IsDuplicatePhoneOrBVN : to verify if BVN Or Number has been used
func (db *authConnection) FindPhone(phone string) (tx *gorm.DB){
	var user entity.User
	return db.connection.Where("phone = ?",phone).Take(&user)
}


func (db *authConnection) FindEmail(email string) (tx *gorm.DB){
	var user entity.User
	return db.connection.Where(&user.Email,email).Take(&user)
}

func (db authConnection) VerifyCredential(email string) interface{} {
	var user entity.User
	//res := db.connection.Where("email = ?",email).Take(&user)
	res := db.connection.Where(&user.Email,email).Find(&user)
	if res.Error != nil {
		return nil
	}
	return user
}

func (db *authConnection) CreateResetCode(passwordReset entity.PasswordReset) entity.PasswordReset{
	min := 10
	max := 30
	passwordReset.Token = string(rand.Intn(max-min+1) + min)
	db.connection.Save(&passwordReset)
	return passwordReset
}

//This is where the password hashing happens
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}