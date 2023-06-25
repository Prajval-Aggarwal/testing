package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/car"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// EquipCarService equips a car for a player.
//
// @Summary Equip Car
// @Description Equip a car for a player
// @Tags Car
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param equipRequest body  request.CarRequest true "Equip Car Request"
// @Success 200 {object} response.Success "Car equipped successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Router /equip-car [put]
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

// @Summary Buy Car
// @Description Buy a car for a player
// @Tags Car
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param carRequest body  request.CarRequest true "Buy Car Request"
// @Security ApiKeyAuth
// @Success 200 {object} response.Success "Car added to player successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 404 {object} response.Success "Car not found"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /buy-car [post]
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

// SellCarService sells a car for a player.
//
// @Summary Sell Car
// @Description Sell a car for a player
// @Tags Car
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param sellCarRequest body request.CarRequest true "Sell Car Request"
// @Security ApiKeyAuth
// @Success 200 {object} response.Success "Car sold successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 404 {object} response.Success "Car not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /sell-car [post]
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

func RepairCarHandler(ctx *gin.Context) {
	var sellCarRequest request.CarRequest
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player is from token is:", playerId)
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

	//car.RepairCarService(ctx, sellCarRequest, playerId.(string))
}
