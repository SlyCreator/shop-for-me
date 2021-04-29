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
	VerifyNumber(ctx *gin.Context)
	ForgetPassword(ctx *gin.Context)
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

	if !c.authService.IsDuplicatePhone(registerDTO.Phone) {
		response := helper.BuildErrorResponse("Failed to process request", "Phone number has been used by another user", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Email has been used by another user", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)

}

func (c authController) Login(ctx *gin.Context)  {
	var loginDTO  dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request",errDTO.Error(),helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	authResult := c.authService.VerifyCredential(loginDTO.Phone,loginDTO.Password)

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

func (c authController) ForgetPassword(ctx *gin.Context)  {
	return
}

func (c authController) VerifyNumber(ctx *gin.Context)  {
	var updatePasswordDTO dto.UpdatePasswordDTO
	errDTO := ctx.ShouldBind(&updatePasswordDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request",errDTO.Error(),helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}
}