package garage

import (
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func AddGarageService(ctx *gin.Context, addGarageReq request.AddGarageRequest) {
	//	var newGarage model.Garage
	newGarage := model.Garage{
		GarageName:    addGarageReq.GarageName,
		Latitude:      addGarageReq.Latitude,
		Longituted:    addGarageReq.Longitute,
		Level:         addGarageReq.Level,
		CoinsRequired: addGarageReq.CoinsRequired,
		Locked:        true,
	}

	err := db.CreateRecord(&newGarage)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.GARAGE_ADD_SUCCESS, utils.HTTP_OK, utils.SUCCESS, newGarage, ctx)
}

func DeleteGarageService(ctx *gin.Context, deleteReq request.DeletGarageReq) {
	//validate the garage id
	if !db.RecordExist("garages", deleteReq.GarageId, "garage_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	query := "DELETE FROM garages WHERE garage_id =?"
	err := db.RawExecutor(query, deleteReq.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	response.ShowResponse(utils.GARAGE_DELETE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func UpdateGarageService(ctx *gin.Context, updateReq request.UpdateGarageReq) {
	var garageDetails model.Garage

	//check if the garage exists or not
	if !db.RecordExist("garages", updateReq.GarageId, "garage_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err := db.FindById(&garageDetails, updateReq.GarageId, "garage_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//Null check on the inputs
	if updateReq.GarageName != "" {
		garageDetails.GarageName = updateReq.GarageName
	}

	if updateReq.Latitude != 0 {
		garageDetails.Latitude = updateReq.Latitude
	}
	if updateReq.Longitute != 0 {
		garageDetails.Longituted = updateReq.Longitute
	}
	if updateReq.Level != 0 {
		garageDetails.Level = updateReq.Level
	}
	if updateReq.CoinsRequired != 0 {
		garageDetails.CoinsRequired = updateReq.CoinsRequired
	}

	err = db.UpdateRecord(&garageDetails, updateReq.GarageId, "garage_id").Error
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.GARAGE_UPDATE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, garageDetails, ctx)

}
func GetAllGarageListService(ctx *gin.Context) {
	var garageList []model.Garage

	query := "SELECT * FROM garages"
	err := db.QueryExecutor(query, &garageList)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.GARAGE_LIST_FETCHED, utils.HTTP_OK, utils.SUCCESS, garageList, ctx)

}

func GetGarageById(ctx *gin.Context) {

}
