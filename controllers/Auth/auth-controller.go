package Auth

import (
	"github.com/gin-gonic/gin"
	"github.com/slycreator/shop-for-me/dto"
	"github.com/slycreator/shop-for-me/entity"
	"github.com/slycreator/shop-for-me/helper"
	"github.com/slycreator/shop-for-me/service"
	"net/http"
	"strconv"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	VerifyToken(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	PasswordReset(ctx *gin.Context)
}
type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController :
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	//fmt.Println(ctx)
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if c.authService.IsPhoneInDB(registerDTO.Phone) {
		response := helper.BuildErrorResponse("Failed to process request", "Phone number has been used by another user", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	if  c.authService.IsEmailInDB(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Email has been used by another user", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)

	}else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}


}

func (c authController) Login(ctx *gin.Context)  {
	var loginDTO  dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request",errDTO.Error(),helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}


	authResult := c.authService.VerifyCredential(loginDTO.Email,loginDTO.Password)

	if v,ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID,10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!",v)
		ctx.JSON(http.StatusOK,response)
		return
	}

	response := helper.BuildErrorResponse("Email or Password doesnt match","Invalid Credential",helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized,response)
}

func (c *authController) PasswordReset(ctx *gin.Context)  {
	var passwordResetDTO  dto.PasswordResetDTO
	errDTO := ctx.ShouldBind(&passwordResetDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsEmailInDB(passwordResetDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Email does not exist", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)

	}else {
		createdToken := c.authService.CreateResetCode(passwordResetDTO)
		token := createdToken.Token
		response := helper.BuildResponse(true, "OK!", token)
		ctx.JSON(http.StatusCreated, response)
	}

}

func (c authController) VerifyToken(ctx *gin.Context)  {
	var verifyResetTokenDTO dto.VerifyResetTokenDTO
	errDTO := ctx.ShouldBind(&verifyResetTokenDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request",errDTO.Error(),helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	if !c.authService.VerifyResetToken(verifyResetTokenDTO.Email,verifyResetTokenDTO.Token) {
		response := helper.BuildErrorResponse("Token is incorrect","Invalid token",helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,response)
	}
	response := helper.BuildResponse(true, "OK!", "Confirmed pls reset update your password")
	ctx.JSON(http.StatusCreated, response)
}

func (c authController) UpdatePassword(ctx *gin.Context)  {
	var verifyResetTokenDTO dto.VerifyResetTokenDTO
	errDTO := ctx.ShouldBind(&verifyResetTokenDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request",errDTO.Error(),helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
}

