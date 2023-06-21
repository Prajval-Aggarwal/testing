package provider

import (
	// "fmt"
	"main/server/response"
	"main/server/services/token"
	"main/server/utils"

	"fmt"

	"github.com/gin-gonic/gin"
)

func UserAuthorization(ctx *gin.Context) {

	fmt.Println("inside middleware")
	tokenString := ctx.Request.Header.Get("Authorization")

	claims, err := token.DecodeToken(tokenString)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		ctx.Abort()
		return
	}
	err = claims.Valid()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		ctx.Abort()
		return
	}
	if claims.Role == "player" {
		ctx.Next()
	} else {
		response.ShowResponse(utils.ACCESS_DENIED, utils.HTTP_FORBIDDEN, utils.FAILURE, nil, ctx)
		ctx.Abort()
		return
	}

	ctx.Set("playerId", claims.Id)

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
