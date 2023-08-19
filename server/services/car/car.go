package car

import (
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// EquipCarService handles the request to equip a car for the player.
func EquipCarService(ctx *gin.Context, equipRequest request.CarRequest, playerId string) {
	// Check if the car with the provided ID is owned by the player.
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM owned_cars WHERE player_id = ? AND car_id = ?)"
	err := db.QueryExecutor(query, &exists, playerId, equipRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Deselect all other cars for the player and select the specified car.
	query = "UPDATE owned_cars SET selected = (car_id = ?) WHERE player_id = ?"
	err = db.RawExecutor(query, equipRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.CAR_SELECETED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func BuyCarService(ctx *gin.Context, carRequest request.CarRequest, playerId string) {
	// Check if the player already owns the car.
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM owned_cars WHERE player_id = ? AND car_id = ?)"
	err := db.QueryExecutor(query, &exists, playerId, carRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if exists {
		response.ShowResponse(utils.CAR_ALREADY_BOUGHT, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Check if the car exists in the database.
	if !db.RecordExist("cars", carRequest.CarId, "car_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Fetch car details from the database.
	var carDetails model.Car
	err = db.FindById(&carDetails, carRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Fetch player details from the database.
	var playerDetails model.Player
	err = db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Check if the player has enough currency to buy the car.
	var amount uint64
	if carDetails.CurrType == "coins" {
		amount = playerDetails.Coins
	} else {
		amount = playerDetails.Cash
	}

	if amount < uint64(carDetails.CurrAmount) {
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Start a database transaction to handle the purchase.
	tx := db.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Deduct the currency from the player's account.
	if carDetails.CurrType == "coins" {
		playerDetails.Coins -= uint64(carDetails.CurrAmount)
	} else {
		playerDetails.Cash -= uint64(carDetails.CurrAmount)
	}

	err = tx.Where("player_id", playerId).Updates(&playerDetails).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Add the car to the player's account and make it the player's current selected car.
	playerCar := model.OwnedCars{
		PlayerId: playerId,
		CarId:    carRequest.CarId,
		Selected: true,
	}

	// Set bought car defaults.
	err = utils.SetPlayerCarDefaults(playerId, carRequest.CarId)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Create a record for the player's owned car.
	err = tx.Create(&playerCar).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Commit the transaction.
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.CAR_BOUGHT_SUCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func SellCarService(ctx *gin.Context, sellCarRequest request.CarRequest, playerId string) {
	// Check if the car to be sold is owned by the player.
	if !db.RecordExist("owned_cars", sellCarRequest.CarId, "car_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Fetch player and car details from the database.
	var playerDetails model.Player
	var carDetails model.Car

	err := db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = db.FindById(&carDetails, sellCarRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Calculate the reward for selling the car.
	var reward uint64
	if carDetails.CurrType == "coins" {
		reward = uint64(0.3 * float64(carDetails.CurrAmount))
	} else {
		reward = uint64(0.6 * float64(carDetails.CurrAmount))
	}

	// Start a database transaction to handle the car selling.
	tx := db.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Add the reward to the player's currency.
	if carDetails.CurrType == "coins" {
		playerDetails.Coins += reward
	} else {
		playerDetails.Cash += reward
	}

	err = tx.Where("player_id", playerId).Updates(&playerDetails).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Delete the car and its related records from the player's collection.
	err = utils.DeleteCarDetails("owned_cars", playerId, sellCarRequest.CarId, ctx)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = utils.DeleteCarDetails("player_cars_stats", playerId, sellCarRequest.CarId, ctx)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = utils.DeleteCarDetails("player_car_upgrades", playerId, sellCarRequest.CarId, ctx)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Commit the transaction.
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.CAR_SOLD_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func RepairCarService(ctx *gin.Context, repairCarRequest request.CarRequest, playerId string) {
	// Fetch player details from the database.
	var playerDetails model.Player
	err := db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Fetch car stats details from the database.
	var carDetails model.CarStats
	err = db.FindById(&carDetails, repairCarRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Fetch player car stats details from the database.
	var playerCarStats model.PlayerCarsStats
	query := "SELECT * FROM player_cars_stats WHERE car_id=? AND player_id=?"
	err = db.QueryExecutor(query, &playerCarStats, repairCarRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Calculate the difference in durability to repair.
	durabilityDiff := carDetails.Durability - playerCarStats.Durability

	// Check if the player has enough repair parts to perform the repair.
	if durabilityDiff*5 > playerDetails.RepairParts {
		response.ShowResponse(utils.NOT_ENOUGH_REPAIR_PARTS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Start a database transaction to handle the car repair.
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	// Update the player car's durability to the car's maximum durability.
	playerCarStats.Durability = carDetails.Durability
	query = "UPDATE player_cars_stats SET durability = ? WHERE player_id=? AND car_id=?"
	err = tx.Exec(query, carDetails.Durability, playerId, repairCarRequest.CarId).Error
	if err != nil {
		tx.Rollback() // Rollback the transaction if there's an error.
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if err := tx.Commit().Error; err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	response.ShowResponse(utils.CAR_REPAIR_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func GetAllCarsService(ctx *gin.Context) {

	var carDetails = []model.Car{}
	query := "SELECT * FROM cars"
	err := db.QueryExecutor(query, &carDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, carDetails, ctx)
}

func GetCarByIdService(ctx *gin.Context, getReq request.CarRequest) {
	var carResponse response.CarResponse
	var carDetails model.Car
	var carStats model.CarStats
	var customization []model.DefaultCustomization

	//check if the car is present with given car id or not

	if !db.RecordExist("cars", getReq.CarId, "car_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	query := "SELECT * FROM cars WHERE car_id=?"
	err := db.QueryExecutor(query, &carDetails, getReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	query = "SELECT * FROM default_customizations WHERE car_id=?"
	err = db.QueryExecutor(query, &customization, getReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	query = "SELECT * FROM car_stats WHERE car_id=?"
	err = db.QueryExecutor(query, &carStats, getReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	carResponse.CarName = carDetails.CarName
	carResponse.Car = carDetails
	carResponse.CarStats = carStats
	carResponse.Customization = customization

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, carResponse, ctx)
}
