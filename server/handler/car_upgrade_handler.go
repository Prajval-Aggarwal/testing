package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/car"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// UpgradeEngineService upgrades the engine of a player's car.
//
// @Summary Upgrade Engine
// @Description Upgrade the engine of a player's car
// @Tags Car-Upgrades
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param upgradeRequest body request.CarUpgradesRequest true "Upgrade Request"
// @Success 200 {object} response.Success "Engine upgraded successfully"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /car/upgrade/engine [put]
func UpgradeEngineHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is ", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var upgradeEngineRequest request.CarUpgradesRequest
	err := utils.RequestDecoding(ctx, &upgradeEngineRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = upgradeEngineRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.UpgradeEngineService(ctx, upgradeEngineRequest, playerId.(string))

}

// UpgradeTurboService upgrades the turbo of a player's car.
//
// @Summary Upgrade Turbo
// @Description Upgrade the engine of a player's car
// @Tags Car-Upgrades
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param upgradeRequest body request.CarUpgradesRequest true "Upgrade Request"
// @Success 200 {object} response.Success "Turbo upgraded successfully"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /car/upgrade/turbo [put]
func UpgradeTurboHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is ", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var upgradeTurboRequest request.CarUpgradesRequest
	err := utils.RequestDecoding(ctx, &upgradeTurboRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = upgradeTurboRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.UpgradeTurboService(ctx, upgradeTurboRequest, playerId.(string))
}

// UpgradeIntakeService upgrades the intake of a player's car.
//
// @Summary Upgrade intake
// @Description Upgrade the engine of a player's car
// @Tags Car-Upgrades
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param upgradeRequest body request.CarUpgradesRequest true "Upgrade Request"
// @Success 200 {object} response.Success "intake upgraded successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /car/upgrade/intake [put]
func UpgradeIntakeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is ", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var upgradeIntakeRequest request.CarUpgradesRequest
	err := utils.RequestDecoding(ctx, &upgradeIntakeRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = upgradeIntakeRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.UpgradeIntakeService(ctx, upgradeIntakeRequest, playerId.(string))

}

// UpgradenitrousService upgrades the nitrous of a player's car.
//
// @Summary Upgrade nitrous
// @Description Upgrade the engine of a player's car
// @Tags Car-Upgrades
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param upgradeRequest body request.CarUpgradesRequest true "Upgrade Request"
// @Success 200 {object} response.Success "nitrous upgraded successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /car/upgrade/nitrous [put]
func UpgradeNitrousHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is ", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var upgradeNitrousRequest request.CarUpgradesRequest
	err := utils.RequestDecoding(ctx, &upgradeNitrousRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = upgradeNitrousRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.UpgradeNitrousService(ctx, upgradeNitrousRequest, playerId.(string))
}

// UpgradebodyService upgrades the body of a player's car.
//
// @Summary Upgrade body
// @Description Upgrade the engine of a player's car
// @Tags Car-Upgrades
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param upgradeRequest body request.CarUpgradesRequest true "Upgrade Request"
// @Success 200 {object} response.Success "body upgraded successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /car/upgrade/body [put]
func UpgradeBodyHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is ", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var upgradeBodyRequest request.CarUpgradesRequest
	err := utils.RequestDecoding(ctx, &upgradeBodyRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = upgradeBodyRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.UpgradeBodyService(ctx, upgradeBodyRequest, playerId.(string))

}

// UpgradetiresService upgrades the tires of a player's car.
//
// @Summary Upgrade tires
// @Description Upgrade the engine of a player's car
// @Tags Car-Upgrades
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param upgradeRequest body request.CarUpgradesRequest true "Upgrade Request"
// @Success 200 {object} response.Success "tires upgraded successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /car/upgrade/tires [put]
func UpgradeTiresHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is ", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var upgradeTiresRequest request.CarUpgradesRequest
	err := utils.RequestDecoding(ctx, &upgradeTiresRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = upgradeTiresRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.UpgradeTiresService(ctx, upgradeTiresRequest, playerId.(string))

}

// UpgradetransmissionService upgrades the transmission of a player's car.
//
// @Summary Upgrade transmission
// @Description Upgrade the engine of a player's car
// @Tags Car-Upgrades
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Param upgradeRequest body request.CarUpgradesRequest true "Upgrade Request"
// @Success 200 {object} response.Success "transmission upgraded successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /car/upgrade/transmission [put]
func UpgradeTransmissionHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id from token is ", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var upgradeTransmissionRequest request.CarUpgradesRequest
	err := utils.RequestDecoding(ctx, &upgradeTransmissionRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = upgradeTransmissionRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.UpgradeTransmissionService(ctx, upgradeTransmissionRequest, playerId.(string))

}
