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
	//check if the garage exists of not
	if !db.RecordExist("garages", buyRequest.GarageId, "garage_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

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

	//add garage to players account
	ownedGarage := model.OwnedGarage{
		PlayerId:    playerId,
		GarageId:    buyRequest.GarageId,
		GarageLevel: 1,
	}

	err = db.CreateRecord(&ownedGarage)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.GARAGE_BOUGHT_SUCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func UpgradeGarageService(ctx *gin.Context, upgradeRequest request.GarageRequest, playerId string) {

	var exists bool
	query := "SELECT EXISTS(SELECT * FROM owned_garages WHERE player_id =? AND garage_id =?"
	err := db.QueryExecutor(query, &exists, playerId, upgradeRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse("Record not fond", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//get current level of that garage
	var ownedGarageStatus model.OwnedGarage
	query = "SELECT garage_level from owned_garages WHERE player_id =? AND garage_id =?"
	err = db.RawExecutor(query, &ownedGarageStatus, playerId, upgradeRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//get player detials
	var playerDetails model.Player
	err = db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//get the upgrades for that current garage
	var garageUpgrades model.GarageUpgrades
	query = "SELECT * FROM garage_upgrades WHERE garage_id=? AND upgrade_level=?"
	db.QueryExecutor(query, &garageUpgrades, upgradeRequest.GarageId, ownedGarageStatus.GarageLevel+1)

	//check if player is compatible to upgrade the garage
	if playerDetails.Coins < garageUpgrades.UpgradeAmount {
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//reduce the playercoins and upgarde the garage level and its car limit
	playerDetails.Coins -= garageUpgrades.UpgradeAmount

	ownedGarageStatus.GarageLevel += 1
	ownedGarageStatus.CarLimit = garageUpgrades.CarLimit

	//update it in database
	err = db.UpdateRecord(&playerDetails, playerId, "player_id").Error
	if err != nil {
		response.ShowResponse(utils.FAILED_TO_UPDATE, utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	query = "UPDATE owned_garages SET garage_level=? , car_limit=? WHERE player_id=? AND garage_id=?"
	err = db.RawExecutor(query, ownedGarageStatus.GarageLevel, garageUpgrades.CarLimit, playerId, upgradeRequest.GarageId)
	if err != nil {
		response.ShowResponse(utils.FAILED_TO_UPDATE, utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.GARAGE_UPGRADED, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
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

func GetPlayerGarageListService(ctx *gin.Context, playerId string) {
	var garageList struct {
		GarageId    string  `json:"garageId"`
		GarageName  string  `json:"garageName"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitute"`
		GarageLevel string  `json:"garageLevel"`
	}

	query := "SELECT pg.garage_id, g.garage_name, g.longitude, g.latitude, pg.garage_level FROM owned_garages pg JOIN garages g ON pg.garage_id = g.garage_id WHERE pg.player_id = ?"

	err := db.QueryExecutor(query, &garageList, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, garageList, ctx)

}

func AddCarToGarageService(ctx *gin.Context, addCarRequest request.AddCarRequest, playerId string) {
	//check if the car is bought
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM owned_garages WHERE player_id =? AND car_id =?"
	err := db.QueryExecutor(query, &exists, playerId, addCarRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse("Record not fond", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check if the garage is bought or not
	query = "SELECT EXISTS(SELECT * FROM owned_garages WHERE player_id =? AND garage_id =?"
	err = db.QueryExecutor(query, &exists, playerId, addCarRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check the car limit of that garage
	var ownedGarageDetails model.OwnedGarage
	query = "SELECT * FROM owned_garages WHERE player_id=? ADN garage_id=?"
	err = db.QueryExecutor(query, &ownedGarageDetails, playerId, addCarRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//get the count of cars in that garage and then compare
	var count int64
	query = "SELECT count(*) FROM  garage_car_lists WHERE player_id=? AND garage_id=?"
	err = db.QueryExecutor(query, &count, playerId, addCarRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if count == int64(ownedGarageDetails.CarLimit) {
		response.ShowResponse("Car Limit reached upgarde the garage to increse the limit", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//add car to the garage
	addCar := model.GarageCarList{
		GarageId: addCarRequest.GarageId,
		CarId:    addCarRequest.CarId,
		PlayerId: playerId,
	}

	err = db.CreateRecord(&addCar)
	if err != nil {
		response.ShowResponse(utils.ADD_CAR_TO_GARAGE_FAILED, utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse("Car added to garage sucessfully", utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}
