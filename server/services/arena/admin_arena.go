package arena

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

func AddArenaService(ctx *gin.Context, addArenaReq request.AddArenaRequest) {
	//	var newArena model.Arena

	var exists bool

	query := "SELECT EXISTS (SELECT * FROM arena_levels WHERE type_id=?)"
	err := db.QueryExecutor(query, &exists, addArenaReq.ArenaLevel)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check that no two same Arenas are on same locations
	query = "SELECT EXISTS (SELECT * FROM arenas WHERE latitude=? AND longitude=?)"
	err = db.QueryExecutor(query, &exists, addArenaReq.Latitude, addArenaReq.Longitude)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if exists {
		response.ShowResponse(utils.ARENA_ALREADY_PRESENT, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	newArena := model.Arena{
		ArenaName:  addArenaReq.ArenaName,
		Latitude:   addArenaReq.Latitude,
		Longitude:  addArenaReq.Longitude,
		ArenaLevel: addArenaReq.ArenaLevel,
	}

	err = db.CreateRecord(&newArena)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.ARENA_ADD_SUCCESS, utils.HTTP_OK, utils.SUCCESS, newArena, ctx)
}

func DeleteArenaService(ctx *gin.Context, deleteReq request.DeletArenaReq) {
	//validate the Arena id
	if !db.RecordExist("Arenas", deleteReq.ArenaId, "Arena_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	query := "DELETE FROM arenas WHERE Arena_id =?"
	err := db.RawExecutor(query, deleteReq.ArenaId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	response.ShowResponse(utils.ARENA_DELETE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func UpdateArenaService(ctx *gin.Context, updateReq request.UpdateArenaReq) {
	var ArenaDetails model.Arena

	//check if the Arena exists or not
	if !db.RecordExist("arenas", updateReq.ArenaId, "arena_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err := db.FindById(&ArenaDetails, updateReq.ArenaId, "arena_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//Null check on the inputs
	if updateReq.ArenaName != "" {
		ArenaDetails.ArenaName = updateReq.ArenaName
	}

	if updateReq.Latitude != 0 {
		ArenaDetails.Latitude = updateReq.Latitude
	}
	if updateReq.Longitude != 0 {
		ArenaDetails.Longitude = updateReq.Longitude
	}
	if updateReq.ArenaLevel != 0 {
		ArenaDetails.ArenaLevel = updateReq.ArenaLevel
	}

	err = db.UpdateRecord(&ArenaDetails, updateReq.ArenaId, "arena_id").Error
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.ARENA_UPDATE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, ArenaDetails, ctx)

}
func GetAllArenaService(ctx *gin.Context) {
	var ArenaList = []model.Arena{}
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
	query := fmt.Sprintf("SELECT * FROM arenas ORDER BY created_at DESC LIMIT %d OFFSET %d", limit, skip)

	err = db.QueryExecutor(query, &ArenaList)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var totalCount int
	countQuery := "SELECT COUNT(*) FROM arenas"
	err = db.QueryExecutor(countQuery, &totalCount)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	dataresp.TotalCount = totalCount
	dataresp.Data = ArenaList

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, dataresp, ctx)
}

func GetArenaTypes(ctx *gin.Context) {
	var arenaTypeList = []model.ArenaLevels{}
	var dataresp response.DataResponse
	// Get the query parameters for skip and limit from the request

	// Build the SQL query with skip and limit
	query := "SELECT * FROM arena_levels ORDER BY type_id "

	err := db.QueryExecutor(query, &arenaTypeList)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var totalCount int
	countQuery := "SELECT COUNT(*) FROM arena_levels"
	err = db.QueryExecutor(countQuery, &totalCount)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	dataresp.TotalCount = totalCount
	dataresp.Data = arenaTypeList

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, dataresp, ctx)
}

func AddGargeTypes(ctx *gin.Context) {
	slice := []string{
		"Easy",
		"Medium",
		"Hard",
	}

	for i, val := range slice {
		newType := model.ArenaLevels{
			TypeId:   i,
			TypeName: val,
		}
		err := db.CreateRecord(&newType)
		if err != nil {
			break
		}
	}

}
