package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/car"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

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
