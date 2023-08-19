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
	// Get player details
	var playerDetails model.Player
	err := db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Get garage details
	var garageDetails model.Garage
	// Check if the garage exists or not
	if !db.RecordExist("garages", buyRequest.GarageId, "garage_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	err = db.FindById(&garageDetails, buyRequest.GarageId, "garage_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Check if the player is eligible to buy the garage; if yes, add it to the player's garage list, else return error
	if playerDetails.Coins < uint64(garageDetails.CoinsRequired) || playerDetails.Level <= garageDetails.Level {
		response.ShowResponse(utils.UPGRADE_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Begin a database transaction
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	defer tx.Commit() // Defer the transaction commit to the end of the function

	// Deduct coins from player's account
	playerDetails.Coins -= uint64(garageDetails.CoinsRequired)
	query := "UPDATE players SET coins=? WHERE player_id=?"
	tx.Exec(query, playerDetails.Coins, playerId)
	if err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Add the garage to the player's account
	ownedGarage := model.OwnedGarage{
		PlayerId:    playerId,
		GarageId:    buyRequest.GarageId,
		GarageLevel: 1,
	}

	err = db.CreateRecord(&ownedGarage)
	if err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.GARAGE_BOUGHT_SUCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func UpgradeGarageService(ctx *gin.Context, upgradeRequest request.GarageRequest, playerId string) {
	// Check if the garage exists in player's owned garages
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM owned_garages WHERE player_id=? AND garage_id=?)"
	err := db.QueryExecutor(query, &exists, playerId, upgradeRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Get current level of that garage
	var ownedGarageStatus model.OwnedGarage
	query = "SELECT garage_level FROM owned_garages WHERE player_id=? AND garage_id=?"
	err = db.RawExecutor(query, &ownedGarageStatus, playerId, upgradeRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Get player details
	var playerDetails model.Player
	err = db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Get the upgrades for the current garage level
	var garageUpgrades model.GarageUpgrades
	query = "SELECT * FROM garage_upgrades WHERE garage_id=? AND upgrade_level=?"
	err = db.QueryExecutor(query, &garageUpgrades, upgradeRequest.GarageId, ownedGarageStatus.GarageLevel+1)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Check if player has enough coins to upgrade the garage

	if playerDetails.Coins < garageUpgrades.UpgradeAmount {
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Reduce the player's coins and upgrade the garage level and car limit
	playerDetails.Coins -= garageUpgrades.UpgradeAmount
	ownedGarageStatus.GarageLevel += 1
	ownedGarageStatus.CarLimit = garageUpgrades.CarLimit

	// Update player and garage details in the database using a transaction
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	defer tx.Commit() // Defer the transaction commit to the end of the function

	// Update player's coins in the database
	query = "UPDATE players SET coins=? WHERE player_id=?"
	tx.Exec(query, playerDetails.Coins, playerId)
	if err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		response.ShowResponse(utils.FAILED_TO_UPDATE, utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Update garage level and car limit in the database
	query = "UPDATE owned_garages SET garage_level=?, car_limit=? WHERE player_id=? AND garage_id=?"
	err = tx.Exec(query, ownedGarageStatus.GarageLevel, ownedGarageStatus.CarLimit, playerId, upgradeRequest.GarageId).Error
	if err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		response.ShowResponse(utils.FAILED_TO_UPDATE, utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.GARAGE_UPGRADED, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
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
	// Check if the car is owned by the player
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM owned_garages WHERE player_id=? AND car_id=?)"
	err := db.QueryExecutor(query, &exists, playerId, addCarRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Check if the garage is owned by the player
	query = "SELECT EXISTS(SELECT * FROM owned_garages WHERE player_id=? AND garage_id=?)"
	err = db.QueryExecutor(query, &exists, playerId, addCarRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Get the car limit of that garage
	var ownedGarageDetails model.OwnedGarage
	query = "SELECT * FROM owned_garages WHERE player_id=? AND garage_id=?"
	err = db.QueryExecutor(query, &ownedGarageDetails, playerId, addCarRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Get the count of cars in that garage and then compare with the car limit
	var count int64
	query = "SELECT count(*) FROM garage_car_lists WHERE player_id=? AND garage_id=?"
	err = db.QueryExecutor(query, &count, playerId, addCarRequest.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if count == int64(ownedGarageDetails.CarLimit) {
		response.ShowResponse(utils.CAR_LIMIT_REACHED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Add the car to the garage
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

	response.ShowResponse(utils.CAR_ADDED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func GetGarageOwnersService(ctx *gin.Context, getReq request.GarageRequest) {
	var resp []struct {
		PlayerId   string  `json:"playerId"`
		PlayerName string  `json:"playerName"`
		OVR        float64 `json:"ovr"`
	}
	query := `SELECT
    p.player_id,
    p.player_name,
    SUM(cs.OVR) AS TotalCarOVR
	FROM
		players p
	JOIN
		owned_garages og ON p.player_id = og.player_id
	JOIN
		owned_cars oc ON p.player_id = oc.player_id
	JOIN
		player_car_stats cs ON oc.car_id = cs.car_id
	WHERE
		og.garage_id = ?
	GROUP BY
		p.player_id, p.player_name;`

	err := db.QueryExecutor(query, &resp, getReq.GarageId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, resp, ctx)

}
