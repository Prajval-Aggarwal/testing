package handler

import (
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func BuyGarage(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	
}

func GetGarageList(ctx *gin.Context) {

}

func UpgradeGarage(ctx *gin.Context) {

}

func AddCarToGarage(ctx *gin.Context) {

}
