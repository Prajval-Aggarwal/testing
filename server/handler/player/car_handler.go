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
// @Router /car/equip [put]
func EquipCarHandler(ctx *gin.Context) {
	var equipRequest request.CarRequest
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
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
// @Router /car/buy [post]
func BuyCarHandler(ctx *gin.Context) {
	var carRequest request.CarRequest
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
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
	fmt.Println("car request is", carRequest)
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
// @Router /car/sell [delete]
func SellCarHandler(ctx *gin.Context) {
	var sellCarRequest request.CarRequest
	playerId, exists := ctx.Get("playerId")
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
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

// @Summary Repair car
// @Description Repair a car by deducting the repair cost from the player's coins and updating the car's durability
// @Accept  json
// @Tags Car
// @Produce  json
// @Param repairCarRequest body request.CarRequest true "Repair car request"
// @Success 200 {object} response.Success "Car sold successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 404 {object} response.Success "Car not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /car/repair [post]
func RepairCarHandler(ctx *gin.Context) {
	var repairCarRequest request.CarRequest
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player is from token is:", playerId)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	err := utils.RequestDecoding(ctx, &repairCarRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = repairCarRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	car.RepairCarService(ctx, repairCarRequest, playerId.(string))
}

// @Summary Get car by ID
// @Description Retrieve car details, stats, and customizations by car ID
// @Accept  json
// @Tags Car
// @Produce  json
// @Param getReq body  request.CarRequest true "Get car by ID request"
// @Success 200 {object} response.Success "Data fetch success"
// @Failure 404 {object} response.Success "Not Found"
// @Failure 500 {object} response.Success "Internal Server Error"
// @Router /car/get-by-id [post]
func GetCarByIdHandler(ctx *gin.Context) {
	var getCarReq request.CarRequest
	err := utils.RequestDecoding(ctx, &getCarReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = getCarReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	car.GetCarByIdService(ctx, getCarReq)
}

// GetAllCarsServiceretrieves the list of all car.
//
// @Summary Get All Garage List
// @Description Retrieve the list of all car
// @Tags Car
// @Accept json
// @Produce json
// @Success 200 {object} response.Success "Garage list fetched successfully"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /car/get-all [get]
func GetAllCarsHandler(ctx *gin.Context) {
	car.GetAllCarsService(ctx)
}
