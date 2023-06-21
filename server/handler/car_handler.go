package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/car"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func EquipCarHandler(ctx *gin.Context) {
	var equipRequest request.CarRequest
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	err := utils.RequestDecoding(ctx, &equipRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = equipRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	car.EquipCarService(ctx, equipRequest, playerId.(string))

}

func BuyCarHandler(ctx *gin.Context) {
	var carRequest request.CarRequest
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	err := utils.RequestDecoding(ctx, &carRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = carRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.BuyCarService(ctx, carRequest, playerId.(string))
}

func SellCarHandler(ctx *gin.Context) {
	var sellCarRequest request.CarRequest
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	err := utils.RequestDecoding(ctx, &sellCarRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = sellCarRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	car.SellCarService(ctx, sellCarRequest, playerId.(string))

}
