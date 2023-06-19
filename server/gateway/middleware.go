package gateway

import (
	// "fmt"
	// "gym/server/response"

	"fmt"
	"main/server/response"
	"main/server/services/token"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func AdminAuthorization(ctx *gin.Context) {

	fmt.Println("inside middleware")
	tokenString := ctx.Request.Header.Get("Authorization")

	claims, err := token.DecodeToken(tokenString)
	if err != nil {
		// response.ErrorResponse(ctx, 401, err.Error())
		response.ShowResponse("Unauthorized", int64(utils.HTTP_UNAUTHORIZED), err.Error(), "", ctx)
		ctx.Abort()
		return
	}

	if claims.Role == "admin" {
		ctx.Next()
	} else {
		// response.ErrorResponse(ctx, 403, "Access Denied")
		response.ShowResponse("Access Denied", int64(utils.HTTP_FORBIDDEN), "Forbidden", "", ctx)

		ctx.Abort()
		return

	}
	//set the token details into context for further processing in handler function
	ctx.Next()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
