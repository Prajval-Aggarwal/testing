package gateway

import (
	// "fmt"
	// "gym/server/response"

	"fmt"
	"main/server/response"
	"main/server/services/token"

	"github.com/gin-gonic/gin"
)

func UserAuthorization(c *gin.Context) {

	// fmt.Println("inside middleware")
	// token, err := c.Request.Cookie("cookie")
	// if err != nil {
	// 	response.ErrorResponse(c, 400, err.Error())
	// 	c.Abort()
	// 	return
	// }

	// claims, err := DecodeToken(token.Value)
	// if err != nil {
	// 	response.ErrorResponse(c, 401, err.Error())
	// 	c.Abort()
	// 	return
	// }
	// err = claims.Valid()
	// if err != nil {
	// 	response.ErrorResponse(c, 401, err.Error())
	// 	c.Abort()
	// 	return
	// }
	// if claims.Role == "user" {
	// 	c.Next()
	// } else {
	// 	response.ErrorResponse(c, 403, "Access Denied")
	// 	c.Abort()
	// 	return
	// }

}

func AdminAuthorization(ctx *gin.Context) {

	fmt.Println("inside middleware")
	tokenString := ctx.Request.Header.Get("Authorization")

	claims, err := token.DecodeToken(tokenString)
	if err != nil {
		response.ShowResponse(err.Error(), 400, "failure", nil, ctx)
		ctx.Abort()
		return
	}
	err = claims.Valid()
	if err != nil {
		response.ShowResponse(err.Error(), 400, "failure", nil, ctx)
		ctx.Abort()
		return
	}
	if claims.Role == "admin" {
		ctx.Next()
	} else {
		response.ShowResponse("Access denied", 403, "failure", nil, ctx)
		ctx.Abort()
		return

	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
