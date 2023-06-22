package garage

import (
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func BuyGarageService(ctx *gin.Context, buyRequest request.GarageRequest, playerId string) {

	//get player details
	var playerDetails model.Player
	err := db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//get garage details
	var garageDetails model.Garage
	err = db.FindById(&garageDetails, buyRequest.GarageId, "garage_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//check if the player is eligible to buy the garage if yes add it to player garage list else return error

	if playerDetails.Coins < int64(garageDetails.CoinsRequired) {
		if playerDetails.Level <= int(garageDetails.Level) {
			response.ShowResponse(utils.UPGRADE_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)

		return
	}

	response.ShowResponse(utils.GARAGE_BOUGHT_SUCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func GetPlayerGarageListService(ctx *gin.Context, playerId string) {
	var garageList struct {
		GarageId    string  `json:"garageId"`
		GarageName  string  `json:"garageName"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitute"`
		GarageLevel string  `json:"garageLevel"`
	}

	query := "SELECT pg.garage_id, g.garage_name, g.longitude, g.latitude, pg.garage_level FROM PlayerGarage pg JOIN Garage g ON pg.garage_id = g.garage_id WHERE pg.player_id = ?"

	err := db.QueryExecutor(query, &garageList, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, garageList, ctx)

}
