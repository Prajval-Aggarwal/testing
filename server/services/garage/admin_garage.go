package garage

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddGarageService(ctx *gin.Context, addGarageReq request.AddGarageRequest) {

	var exists bool

	query := "SELECT EXISTS (SELECT * FROM garage_types WHERE type_id=?)"
	err := db.QueryExecutor(query, &exists, addGarageReq.GarageType)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check that no two same garages are on same locations
	query = "SELECT EXISTS (SELECT * FROM garages WHERE latitude=? AND longitude=?)"
	err = db.QueryExecutor(query, &exists, addGarageReq.Latitude, addGarageReq.Longitude)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if exists {
		response.ShowResponse(utils.GARAGE_ALREADY_PRESENT, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//	var newGarage model.Garage
	newGarage := model.Garage{
		GarageName:    addGarageReq.GarageName,
		Latitude:      addGarageReq.Latitude,
		Longitude:     addGarageReq.Longitude,
		Level:         addGarageReq.Level,
		CoinsRequired: addGarageReq.CoinsRequired,
		GarageType:    addGarageReq.GarageType,
		Locked:        true,
	}

	err = db.CreateRecord(&newGarage)
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
	if updateReq.Longitude != 0 {
		garageDetails.Longitude = updateReq.Longitude
	}
	if updateReq.Level != 0 {
		garageDetails.Level = updateReq.Level
	}
	if updateReq.CoinsRequired != 0 {
		garageDetails.CoinsRequired = updateReq.CoinsRequired
	}

	if updateReq.GarageType != 0 {
		garageDetails.GarageType = updateReq.GarageType
	}

	err = db.UpdateRecord(&garageDetails, updateReq.GarageId, "garage_id").Error
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.GARAGE_UPDATE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, garageDetails, ctx)

}
func GetAllGarageListService(ctx *gin.Context) {
	var garageList = []model.Garage{}
	var dataresp response.DataResponse
	// Get the query parameters for skip and limit from the request
	skipParam := ctx.DefaultQuery("skip", "0")
	limitParam := ctx.DefaultQuery("limit", "10")

	// Convert skip and limit to integers
	skip, err := strconv.Atoi(skipParam)
	if err != nil {
		response.ShowResponse("Invalid skip value", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		response.ShowResponse("Invalid limit value", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Build the SQL query with skip and limit
	query := fmt.Sprintf("SELECT * FROM garages ORDER BY created_at DESC LIMIT %d OFFSET %d", limit, skip)

	err = db.QueryExecutor(query, &garageList)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var totalCount int
	countQuery := "SELECT COUNT(*) FROM garages"
	err = db.QueryExecutor(countQuery, &totalCount)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	dataresp.TotalCount = totalCount
	dataresp.Data = garageList

	response.ShowResponse(utils.GARAGE_LIST_FETCHED, utils.HTTP_OK, utils.SUCCESS, dataresp, ctx)
}

func AddGargeTypes(ctx *gin.Context) {
	slice := []string{
		"The Great Spot",
		"Princes Palace",
		"The Bear's Hideaway",
		"Red's Hotspot ",
		"The Mu",
	}

	for i, val := range slice {
		newType := model.GarageTypes{
			TypeId:   i,
			TypeName: val,
		}
		err := db.CreateRecord(&newType)
		if err != nil {
			break
		}
	}

}

func GetGarageTypes(ctx *gin.Context) {
	var garageTypeList = []model.GarageTypes{}
	var dataresp response.DataResponse
	// Get the query parameters for skip and limit from the request

	// Build the SQL query with skip and limit
	query := "SELECT * FROM garage_types ORDER BY type_id"

	err := db.QueryExecutor(query, &garageTypeList)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var totalCount int
	countQuery := "SELECT COUNT(*) FROM garage_types"
	err = db.QueryExecutor(countQuery, &totalCount)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	dataresp.TotalCount = totalCount
	dataresp.Data = garageTypeList

	response.ShowResponse(utils.GARAGE_LIST_FETCHED, utils.HTTP_OK, utils.SUCCESS, dataresp, ctx)
}
