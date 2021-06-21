package middleware

import (
	"log"
	"net/http"

	"github.com/AbdulAffif/hallobumil_dev/api/helper"
	"github.com/AbdulAffif/hallobumil_dev/api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		autHeader := c.GetHeader("Authorization")
		if autHeader == "" {
			response := helper.BuildErrorResponse(http.StatusBadRequest, "no token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(autHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[iss]: ", claims["iss"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse(http.StatusUnauthorized, err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
