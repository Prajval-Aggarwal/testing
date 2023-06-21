package handler

import (
	"main/server/request"
	"main/server/response"
	"main/server/services/car"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func EquipCar(ctx *gin.Context) {

}

func BuyCarHandler(ctx *gin.Context) {
	var carRequest request.BuyCarRequest
	playerId, _ := ctx.Get("playerId")

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
