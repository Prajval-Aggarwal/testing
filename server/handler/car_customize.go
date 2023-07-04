package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/car"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// ColorCustomizationService handles color customization.
//
// @Summary Color Customization
// @Description Customizes the color of a player's car.
// @Accept json
// @Produce json
// @Tags Car-Customize
// @Param playerId header string true "The ID of the player"
// @Param colorReq body request.ColorCustomizationRequest true "Color customization request object"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Success
// @Failure 404 {object} response.Success
// @Failure 500 {object} response.Success
// @Router /car/customise/color [put]
func ColorCustomizeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId)

	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	var colorReq request.ColorCustomizationRequest
	err := utils.RequestDecoding(ctx, &colorReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, err.Error(), nil, ctx)
		return
	}
	err = colorReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.ColorCustomizationService(ctx, colorReq, playerId.(string))

}

// WheelCustomizeService handles wheel customization.
//
// @Summary Wheel Customization
// @Description Customizes the wheels of a player's car.
// @Accept json
// @Tags Car-Customize
// @Produce json
// @Param playerId header string true "The ID of the player"
// @Param wheelReq body request.WheelCustomizeRequest true "Wheel customization request object"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Success
// @Failure 404 {object} response.Success
// @Failure 500 {object} response.Success
// @Router /car/customise/wheels [put]
func WheelsCustomizeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	var wheelReq request.WheelCustomizeRequest
	err := utils.RequestDecoding(ctx, &wheelReq)
	if err != nil {

		fmt.Println("error in decoding")
		response.ShowResponse("Bad Request", utils.HTTP_BAD_REQUEST, err.Error(), nil, ctx)
		return
	}

	err = wheelReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//call the service method
	car.WheelCustomizeService(ctx, wheelReq, playerId.(string))

}

// InteriorCustomizeService handles interior customization.
//
// @Summary Interior Customization
// @Description Customizes the interior of a player's car.
// @Accept json
// @Produce json
// @Tags Car-Customize
// @Param playerId header string true "The ID of the player"
// @Param interiorReq body request.InteriorCustomizeRequest true "Interior customization request object"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Success
// @Failure 404 {object} response.Success
// @Failure 500 {object} response.Success
// @Router /car/customise/interior [put]
func InteriorCustomizeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	var interiorReq request.InteriorCustomizeRequest
	err := utils.RequestDecoding(ctx, &interiorReq)
	if err != nil {

		fmt.Println("error in decoding")
		response.ShowResponse("Bad Request", utils.HTTP_BAD_REQUEST, err.Error(), nil, ctx)
		return
	}

	err = interiorReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//call the service method
	car.InteriorCustomizeService(ctx, interiorReq, playerId.(string))

}

// LicenseCustomizeService handles license plate customization.
//
// @Summary License Plate Customization
// @Description Customizes the license plate of a player's car.
// @Accept json
// @Tags Car-Customize
// @Produce json
// @Param playerId header string true "The ID of the player"
// @Param licenseRequest body request.LicenseRequest true "License plate customization request object"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Success
// @Failure 500 {object} response.Success
// @Router /car/customise/license [put]
func LicenseCustomizeHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	var licenseRequest request.LicenseRequest
	err := utils.RequestDecoding(ctx, &licenseRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = licenseRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	car.LicenseCustomizeService(ctx, licenseRequest, playerId.(string))

}

// @Summary Get car customization parts
// @Description Retrieve a list of car customization parts
// @Accept  json
// @Tags Car-Customize
// @Produce  json
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Success
// @Failure 500 {object} response.Success
// @Router /get-customization [get]
func GetCarCustomiseHandler(ctx *gin.Context) {
	car.GetCarCustomizationParts(ctx)
}

// @Summary Get car color categories
// @Description Retrieve a list of car color categories for a specific part
// @Tags Car-Customize
// @Accept  json
// @Produce  json
// @Param part body request.GetCarColorCategoriesRequest true "Car customization part"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Success
// @Failure 500 {object} response.Success
// @Router /get-color-category [get]
func GetCarColorCategoriesHandler(ctx *gin.Context) {
	var categoryReq request.GetCarColorCategoriesRequest
	err := utils.RequestDecoding(ctx, &categoryReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = categoryReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.GetCarColorCategories(ctx, categoryReq.Part)
}

// @Summary Get car color types
// @Description Retrieve a list of car color types for a specific part and color category
// @Accept  json
// @Tags Car-Customize
// @Produce  json
// @Param part body request.GetCarColorTypesRequest true "Car color type like fluorescent,military"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Success
// @Failure 500 {object} response.Success
// @Router /get-color-type [get]
func GetCarColorTypesHandler(ctx *gin.Context) {
	var carColorReq request.GetCarColorTypesRequest
	err := utils.RequestDecoding(ctx, &carColorReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = carColorReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.GetCarColorTypes(ctx, carColorReq)
}

// @Summary Get car color types
// @Description Retrieve a list of car color types for a specific part and color category
// @Accept  json
// @Tags Car-Customize
// @Produce  json
// @Param part body request.GetCarColorTypesRequest true "Car color type like fluorescent,military"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Success
// @Failure 500 {object} response.Success
// @Router /get-colors [get]
func GetCarColorsHandler(ctx *gin.Context) {
	var carColorReq request.GetCarColorRequest
	err := utils.RequestDecoding(ctx, &carColorReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	err = carColorReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	car.GetCarColors(ctx, carColorReq)
}
