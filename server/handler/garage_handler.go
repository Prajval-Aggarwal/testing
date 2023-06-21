package handler

import (
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func BuyGarageHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

}

func GetAllGarageListHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
}

func UpgradeGarageHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
}

func AddCarToGarageHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
}

// gives the list of garages owned by users
func GetPlayerGarageListHandler(ctx *gin.Context) {

}
