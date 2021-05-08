package repository

import (
	"github.com/slycreator/shop-for-me/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type AuthRepository interface {
	CreateUser(user entity.User) entity.User
	IsDuplicatePhone(phone string) (tx *gorm.DB)
	IsDuplicateEmail(phone string) (tx *gorm.DB)
	VerifyCredential(phone string) interface{}
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
func (db *authConnection) IsDuplicatePhone(phone string) (tx *gorm.DB){
	var user entity.User
	return db.connection.Where("phone = ?",phone).Take(&user)
}


func (db *authConnection) IsDuplicateEmail(email string) (tx *gorm.DB){
	var user entity.User
	return db.connection.Where("email = ?",email).Take(&user)
}

func (db authConnection) VerifyCredential(email string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?",email).Take(&user)
	if res.Error != nil {
		return nil
	}
	return user
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