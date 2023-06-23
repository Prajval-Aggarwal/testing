package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/garage"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func BuyGarageHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is:", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var buyRequest request.GarageRequest
	err := utils.RequestDecoding(ctx, &buyRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = buyRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	garage.BuyGarageService(ctx, buyRequest, playerId.(string))

}

func GetAllGarageListHandler(ctx *gin.Context) {
	garage.GetAllGarageListService(ctx)
}

func UpgradeGarageHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is:", playerId)

	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var upgradeRequest request.GarageRequest
	err := utils.RequestDecoding(ctx, &upgradeRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = upgradeRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	garage.UpgradeGarageService(ctx, upgradeRequest, playerId.(string))
}

func AddCarToGarageHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is:", playerId)

	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var addCarRequest request.AddCarRequest
	err := utils.RequestDecoding(ctx, &addCarRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = addCarRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	garage.AddCarToGarageService(ctx, addCarRequest, playerId.(string))
}

// gives the list of garages owned by users
func GetPlayerGarageListHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is:", playerId)

	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	garage.GetPlayerGarageListService(ctx, playerId.(string))
}
