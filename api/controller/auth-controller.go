package controller

import (
	"github.com/AbdulAffif/hallobumil_dev/api/dto"
	"github.com/AbdulAffif/hallobumil_dev/api/entity"
	"github.com/AbdulAffif/hallobumil_dev/api/helper"
	"github.com/AbdulAffif/hallobumil_dev/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthController interface {
	LoginEmail(ctx *gin.Context)
	LoginPhone(ctx *gin.Context)
	Register(ctx *gin.Context)
	Ping(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}
func (c *authController) LoginEmail(ctx *gin.Context) {
	var LoginDTO dto.LoginDTOEmail
	errDTO := ctx.ShouldBind(&LoginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	errValid := LoginDTO.Validate()
	if errValid != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest,errValid.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyEmail(LoginDTO.Email, LoginDTO.Password,LoginDTO.Identifier)
	if v, ok := authResult.(entity.JsonRegister); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.User.ID, 10))
		v.JwtToken = generatedToken
		response := helper.BuildResponse(http.StatusOK, "Login Success", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse(http.StatusUnauthorized, "invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)

}

func (c *authController) LoginPhone(ctx *gin.Context) {
	var LoginDTO dto.LoginDTOPhone
	errDTO := ctx.ShouldBind(&LoginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest, errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	errValid := LoginDTO.Validate()
	if errValid != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest,errValid.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyPhone(LoginDTO.Phone, LoginDTO.Password, LoginDTO.Identifier)
	if v, ok := authResult.(entity.JsonRegister); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.User.ID, 10))
		v.JwtToken = generatedToken
		response := helper.BuildResponse(http.StatusOK, "Login Success", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse(http.StatusUnauthorized, "invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)

}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO

	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest,errDTO.Error() , helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	errValid := registerDTO.Validate()
	if errValid != nil {
		response := helper.BuildErrorResponse(http.StatusBadRequest,errValid.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse(http.StatusConflict, "Duplicate Email", helper.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.User.ID, 10))
		createdUser.JwtToken = token
		reponse := helper.BuildResponse(http.StatusCreated, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, reponse)
	}
}

func (c *authController) Ping(ctx *gin.Context) {
	if !c.authService.Ping(){
		response := helper.BuildErrorResponse(http.StatusBadRequest, "Bad Request", helper.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	}else {
		response := helper.BuildResponse(http.StatusOK, "OK!", "")
		ctx.JSON(http.StatusOK, response)
	}
}
