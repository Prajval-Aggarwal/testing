package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/car"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func ColorCustomizeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId)

	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

}

func WheelsCustomizeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var wheelReq request.WheelCustomizeRequest
	err := utils.RequestDecoding(ctx, &wheelReq)
	if err != nil {

		fmt.Println("error in decoding")
		return
	}

	//call the service method
	car.WheelCustomizeService(ctx, wheelReq, playerId.(string))

}

func InteriorCustomizeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	var interiorReq request.InteriorCustomizeRequest
	err := utils.RequestDecoding(ctx, &interiorReq)
	if err != nil {

		fmt.Println("error in decoding")
		return
	}

	//call the service method
	car.InteriorCustomizeService(ctx, interiorReq, playerId.(string))

}

func LicenseCustomizeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId)
	if !exists {
		response.ShowResponse("Unauthorised", utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	var licenseRequest request.LicenseRequest
	err := utils.RequestDecoding(ctx, &licenseRequest)
	if err != nil {
		// response.ShowResponse()
		return
	}

	car.LicenseCustomizeService(ctx, licenseRequest, playerId.(string))

}
