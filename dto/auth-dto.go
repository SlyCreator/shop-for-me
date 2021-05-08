package dto

//RegisterDTO : is used when client post from /register url
type RegisterDTO struct {
	FirstName			string		`json:"first_name" form:"first_name" binding:"required"`
	LastName			string		`json:"last_name"form:"last_name" binding:"required"`
	Email				string		`json:"email" form:"email" binding:"required,email" `
	Password			string		`json:"password" form:"password" binding:"required"`
	DataOFBirth			string		`json:"data_of_birth" form:"data_of_birth" binding:"required"`
	Phone				string		`json:"phone" form:"password" binding:"required"`
	Citizenship			string		`json:"citizenship" form:"citizenship"`
}

//LoginDTO is use to to receive data on registration
type LoginDTO struct {
	Email				string		`json:"email" form:"email" binding:"required,email" `
	Password			string		`json:"password" form:"password" binding:"required"`
}

type PasswordResetDTO struct {
	Email				string		`json:"email" form:"email" binding:"required,email" `
}

//RegisterDTO is use to to receive data on registration
type UpdatePasswordDTO struct {
	OldPassword	string		`json:"password" form:password binding:required`
	NewPassword1	string		`json:"password" form:password binding:required`
	NewPassword2	string		`json:"password" form:password binding:required`
}