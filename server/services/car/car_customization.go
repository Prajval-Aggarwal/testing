package car

import (
	"main/server/db"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func WheelCustomizeService(ctx *gin.Context, wheelReq request.WheelCustomizeRequest, playerId string) {
	if !utils.IsCarEquipped(playerId, wheelReq.CarId) {
		response.ShowResponse(utils.EQUIP_CORRECT_CAR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

}

func InteriorCustomizeService(ctx *gin.Context, interiorReq request.InteriorCustomizeRequest, playerId string) {

	//check the carId s of selected car only

	if !utils.IsCarEquipped(playerId, interiorReq.CarId) {
		response.ShowResponse(utils.EQUIP_CORRECT_CAR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)

		return
	}
	var exists bool
	//check that the color is present in tht database or not
	query := "SELECT EXISTS(SELECT * FROM owned_cars WHERE part =? AND color_name=?"
	err := db.QueryExecutor(query, &exists, "interior", interiorReq.ColorName)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//query to update interior color
	query = "UPDATE player_car_customizations SET color_code=? AND color_name=? WHERE player_id=? AND car_id=? AND part=?"

	err = db.QueryExecutor(query, interiorReq.ColorCode, interiorReq.ColorName, playerId, interiorReq.CarId, "interior")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.INTERIOR_CUSTOMIZED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func LicenseCustomizeService(ctx *gin.Context, licenseRequest request.LicenseRequest, playerId string) {
	//check if the car id is equiped or not
	if !utils.IsCarEquipped(playerId, licenseRequest.CarId) {
		response.ShowResponse(utils.EQUIP_CORRECT_CAR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	query := "UPDATE player_car_customizations SET value=? WHERE player_id=? AND car_id=? part=?"
	err := db.QueryExecutor(query, licenseRequest.Value, playerId, licenseRequest.CarId, "license")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.LICENSE_PLATE_CUSTOMIZED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}
