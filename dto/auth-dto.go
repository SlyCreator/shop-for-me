package dto

//RegisterDTO : is used when client post from /register url
type RegisterDTO struct {
	Phone				string		`json:"phone" form:"password" binding:"required"`
	Password			string		`json:"password" form:"password" binding:"required"`
	Email				string		`json:"email" form:"email" binding:"required,email" `
	FirstName			string		`json:"first_name" form:"first_name" binding:"required"`
	LastName			string		`json:"last_name"form:"last_name" binding:"required"`
	DataOFBirth			string		`json:"data_of_birth" form:"data_of_birth" binding:"required"`
	BVN					string		`json:"bvn" form:"bvn" binding:"required"`
	CountryOfResidence  string		`json:"country_of_residence" form:"country_of_residence" binding:"required"`
	Citizenship			string		`json:"citizenship" form:"citizenship" binding:"required"`
}

//RegisterDTO is use to to receive data on registration
type LoginDTO struct {
	Phone		string		`json:"phone" form:phone binding:required`
	Password	string		`json:"password" form:password binding:required`
}

//RegisterDTO is use to to receive data on registration
type UpdatePasswordDTO struct {
	OldPassword	string		`json:"password" form:password binding:required`
	NewPassword1	string		`json:"password" form:password binding:required`
	NewPassword2	string		`json:"password" form:password binding:required`
}