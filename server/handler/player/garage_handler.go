package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/garage"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// BuyGarageService buys a garage for a player.
//
// @Summary Buy Garage
// @Description Buy a garage for a player
// @Tags Garage
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param buyRequest body request.GarageRequest true "Buy Garage Request"
// @Security ApiKeyAuth
// @Success 200 {object} response.Success "Garage bought successfully"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 404 {object} response.Success "Garage not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /garage/buy [post]
func BuyGarageHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get(utils.PLAYER_ID)
	fmt.Println("player id from token is:", playerId)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
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

// UpgradeGarageService upgrades a player's garage.
//
// @Summary Upgrade Garage
// @Description Upgrade a player's garage to the next level
// @Tags Garage
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param upgradeRequest body request.GarageRequest true "Upgrade Request"
// @Success 200 {object} response.Success "Garage upgraded successfully"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /garage/upgrade [put]
func UpgradeGarageHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get(utils.PLAYER_ID)
	fmt.Println("player id from token is:", playerId)

	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
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

// AddCarToGarageService adds a car to a player's garage.
//
// @Summary Add Car to Garage
// @Description Add a car to a player's garage
// @Tags Garage
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param addCarRequest body request.AddCarRequest true "Add Car Request"
// @Success 200 {object} response.Success "Car added to garage successfully"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /garage/add-car [post]
func AddCarToGarageHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get(utils.PLAYER_ID)
	fmt.Println("player id from token is:", playerId)

	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
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

// GetPlayerGarageListService retrieves the list of garages owned by a player.
//
// @Summary Get Player Garage List
// @Description Retrieve the list of garages owned by a player
// @Tags Garage
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Success 200 {object} response.Success "Data fetched successfully"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /garage/get [get]
func GetPlayerGarageListHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get(utils.PLAYER_ID)
	fmt.Println("player id from token is:", playerId)

	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	garage.GetPlayerGarageListService(ctx, playerId.(string))
}
