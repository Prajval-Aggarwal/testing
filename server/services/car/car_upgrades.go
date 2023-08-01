package car

import (
	"fmt"
	"main/server/db"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func UpgradeEngineService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	var upgradeCost uint64
	isUpgradable := true

	playerDetails, playerCarStats, playerCarUpgrades, carClassDetails, maxUpgradeLevel, classRating, err := utils.UpgradeData(playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	if playerCarUpgrades.Engine+1 == maxUpgradeLevel {
		isUpgradable = false
	}
	fmt.Println("car class details is", carClassDetails)
	//check for max upgrade
	if playerCarUpgrades.Engine == (maxUpgradeLevel) {
		//part cannnot be upgraded further
		response.ShowResponse(utils.UPGRADE_REACHED_MAX_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	query := "SELECT cost FROM upgrades WHERE upgrade_level=? AND class=?"
	err = db.QueryExecutor(query, &upgradeCost, playerCarUpgrades.Engine+1, carClassDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check for players money
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if playerDetails.Coins < upgradeCost {
		//player donot have enough coin to buy the upgrade
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("upgrade cost is", upgradeCost)
	playerDetails.Coins -= upgradeCost

	query = "UPDATE players SET coins=? WHERE player_id=?"
	tx.Exec(query, playerDetails.Coins, playerId)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//upgrade engine,power and overall rating of the car in player car stats and upgrade

	query = "UPDATE player_car_upgrades SET engine=? WHERE player_id=? AND car_id=? "
	err = tx.Exec(query, playerCarUpgrades.Engine+1, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var ovr = utils.RoundFloat(utils.CalculateOVR(classRating.ORMultiplier, float64(playerCarStats.Power+utils.UPGRADE_POWER), float64(playerCarStats.Grip), float64(playerCarStats.Weight)), 2)

	query = "UPDATE player_cars_stats set power=? , ovr=? WHERE player_id=? AND car_id=?"
	err = tx.Exec(query, (playerCarStats.Power + utils.UPGRADE_POWER), ovr, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if err := tx.Commit().Error; err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var nextCost uint64
	query = "SELECT cost FROM upgrades WHERE upgrade_level=? AND class=?"
	err = db.QueryExecutor(query, &nextCost, playerCarUpgrades.Engine+2, carClassDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	upgradeResponse := response.CarUpgradeResponse{
		UpgradedPart: "engine",
		Power:        (playerCarStats.Power + utils.UPGRADE_POWER),
		OVR:          ovr,
		Coins:        playerDetails.Coins,
	}
	if isUpgradable {
		upgradeResponse.Upgradable = true
		upgradeResponse.NextLevel = playerCarUpgrades.Engine + 2
		upgradeResponse.NextCost = nextCost
	} else {
		upgradeResponse.Upgradable = false
	}

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, upgradeResponse, ctx)

}

func UpgradeTurboService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	var upgradeCost uint64
	isUpgradable := true

	playerDetails, playerCarStats, playerCarUpgrades, carClassDetails, maxUpgradeLevel, classRating, err := utils.UpgradeData(playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("car class details is", carClassDetails)

	if playerCarUpgrades.Turbo+1 == maxUpgradeLevel {
		isUpgradable = false
	}
	//check for max upgrade
	if playerCarUpgrades.Turbo == (maxUpgradeLevel) {

		//part cannnot be upgraded further
		response.ShowResponse(utils.UPGRADE_REACHED_MAX_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check for players money
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if playerDetails.Coins < upgradeCost {
		//player donot have enough coin to buy the upgrade
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	playerDetails.Coins -= upgradeCost

	query := "UPDATE players SET coins=? WHERE player_id=?"
	tx.Exec(query, playerDetails.Coins, playerId)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//upgrade engine,power and overall rating of the car in player car stats and upgrade

	query = "UPDATE player_car_upgrades SET turbo=? WHERE player_id=? AND car_id=? "
	err = tx.Exec(query, playerCarUpgrades.Turbo+1, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var ovr = utils.RoundFloat(utils.CalculateOVR(classRating.ORMultiplier, float64(playerCarStats.Power+utils.UPGRADE_POWER), float64(playerCarStats.Grip), float64(playerCarStats.Weight)), 2)

	query = "UPDATE player_cars_stats set power=? , ovr=? WHERE player_id=? AND car_id=?"
	err = tx.Exec(query, (playerCarStats.Power + utils.UPGRADE_POWER), ovr, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if err := tx.Commit().Error; err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var nextCost uint64
	query = "SELECT cost FROM upgrades WHERE upgrade_level=? AND class=?"
	err = db.QueryExecutor(query, &nextCost, playerCarUpgrades.Turbo+2, carClassDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	upgradeResponse := response.CarUpgradeResponse{
		UpgradedPart: "turbo",
		Power:        (playerCarStats.Power + utils.UPGRADE_POWER),
		OVR:          ovr,
		Coins:        playerDetails.Coins,
	}
	if isUpgradable {
		upgradeResponse.Upgradable = true
		upgradeResponse.NextLevel = playerCarUpgrades.Turbo + 2
		upgradeResponse.NextCost = nextCost
	} else {
		upgradeResponse.Upgradable = false
	}

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, upgradeResponse, ctx)

}

func UpgradeIntakeService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	var upgradeCost uint64
	isUpgradable := true

	playerDetails, playerCarStats, playerCarUpgrades, carClassDetails, maxUpgradeLevel, classRating, err := utils.UpgradeData(playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("car class details is", carClassDetails)
	if playerCarUpgrades.Intake+1 == maxUpgradeLevel {
		isUpgradable = false
	}
	//check for max upgrade
	if playerCarUpgrades.Intake == (maxUpgradeLevel) {

		//part cannnot be upgraded further
		response.ShowResponse(utils.UPGRADE_REACHED_MAX_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check for players money
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if playerDetails.Coins < upgradeCost {
		//player donot have enough coin to buy the upgrade
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	playerDetails.Coins -= upgradeCost

	query := "UPDATE players SET coins=? WHERE player_id=?"
	tx.Exec(query, playerDetails.Coins, playerId)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//upgrade engine,power and overall rating of the car in player car stats and upgrade

	query = "UPDATE player_car_upgrades SET intake=? WHERE player_id=? AND car_id=? "
	err = tx.Exec(query, playerCarUpgrades.Intake+1, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var ovr = utils.RoundFloat(utils.CalculateOVR(classRating.ORMultiplier, float64(playerCarStats.Power+utils.UPGRADE_POWER), float64(playerCarStats.Grip), float64(playerCarStats.Weight)), 2)

	query = "UPDATE player_cars_stats set power=? , ovr=? WHERE player_id=? AND car_id=?"
	err = tx.Exec(query, (playerCarStats.Power + utils.UPGRADE_POWER), ovr, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if err := tx.Commit().Error; err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var nextCost uint64
	query = "SELECT cost FROM upgrades WHERE upgrade_level=? AND class=?"
	err = db.QueryExecutor(query, &nextCost, playerCarUpgrades.Intake+2, carClassDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	upgradeResponse := response.CarUpgradeResponse{
		UpgradedPart: "intake",
		Power:        (playerCarStats.Power + utils.UPGRADE_POWER),
		OVR:          ovr,
		Coins:        playerDetails.Coins,
	}
	if isUpgradable {
		upgradeResponse.Upgradable = true
		upgradeResponse.NextLevel = playerCarUpgrades.Intake + 2
		upgradeResponse.NextCost = nextCost
	} else {
		upgradeResponse.Upgradable = false
	}

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, upgradeResponse, ctx)

}

func UpgradeBodyService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	var upgradeCost uint64
	isUpgradable := true

	playerDetails, playerCarStats, playerCarUpgrades, carClassDetails, maxUpgradeLevel, classRating, err := utils.UpgradeData(playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("car class details is", carClassDetails)
	if playerCarUpgrades.Body+1 == maxUpgradeLevel {
		isUpgradable = false
	}
	//check for max upgrade
	if playerCarUpgrades.Body == (maxUpgradeLevel) {

		//part cannnot be upgraded further
		response.ShowResponse(utils.UPGRADE_REACHED_MAX_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check for players money
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if playerDetails.Coins < upgradeCost {
		//player donot have enough coin to buy the upgrade
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	playerDetails.Coins -= upgradeCost

	query := "UPDATE players SET coins=? WHERE player_id=?"
	tx.Exec(query, playerDetails.Coins, playerId)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//upgrade engine,power and overall rating of the car in player car stats and upgrade

	query = "UPDATE player_car_upgrades SET body=? WHERE player_id=? AND car_id=? "
	err = tx.Exec(query, playerCarUpgrades.Body+1, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var ovr = utils.RoundFloat(utils.CalculateOVR(classRating.ORMultiplier, float64(playerCarStats.Power), float64(playerCarStats.Grip+uint64(utils.UPGRADE_GRIP)), float64(playerCarStats.Weight)), 2)

	query = "UPDATE player_cars_stats set grip=? , ovr=? WHERE player_id=? AND car_id=?"
	err = tx.Exec(query, (playerCarStats.Grip + uint64(utils.UPGRADE_GRIP)), ovr, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if err := tx.Commit().Error; err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var nextCost uint64
	query = "SELECT cost FROM upgrades WHERE upgrade_level=? AND class=?"
	err = db.QueryExecutor(query, &nextCost, playerCarUpgrades.Body+2, carClassDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	upgradeResponse := response.CarUpgradeResponse{
		UpgradedPart: "body",
		Grip:         playerCarStats.Grip +uint64(utils.UPGRADE_GRIP),
		OVR:          ovr,
		Coins:        playerDetails.Coins,
	}
	if isUpgradable {
		upgradeResponse.Upgradable = true
		upgradeResponse.NextLevel = playerCarUpgrades.Body + 2
		upgradeResponse.NextCost = nextCost
	} else {
		upgradeResponse.Upgradable = false
	}

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, upgradeResponse, ctx)

}

func UpgradeTiresService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	var upgradeCost uint64
	isUpgradable := true

	playerDetails, playerCarStats, playerCarUpgrades, carClassDetails, maxUpgradeLevel, classRating, err := utils.UpgradeData(playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("car class details is", carClassDetails)
	if playerCarUpgrades.Tires+1 == maxUpgradeLevel {
		isUpgradable = false
	}
	//check for max upgrade
	if playerCarUpgrades.Tires == (maxUpgradeLevel) {

		//part cannnot be upgraded further
		response.ShowResponse(utils.UPGRADE_REACHED_MAX_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check for players money
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if playerDetails.Coins < upgradeCost {
		//player donot have enough coin to buy the upgrade
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	playerDetails.Coins -= upgradeCost

	query := "UPDATE players SET coins=? WHERE player_id=?"
	tx.Exec(query, playerDetails.Coins, playerId)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//upgrade engine,power and overall rating of the car in player car stats and upgrade

	query = "UPDATE player_car_upgrades SET tries=? WHERE player_id=? AND car_id=? "
	err = tx.Exec(query, playerCarUpgrades.Tires+1, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var ovr = utils.RoundFloat(utils.CalculateOVR(classRating.ORMultiplier, float64(playerCarStats.Power), float64(playerCarStats.Grip+uint64(utils.UPGRADE_GRIP)), float64(playerCarStats.Weight)), 2)

	query = "UPDATE player_cars_stats set grip=? , ovr=? WHERE player_id=? AND car_id=?"
	err = tx.Exec(query, (playerCarStats.Grip + uint64(utils.UPGRADE_GRIP)), ovr, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if err := tx.Commit().Error; err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var nextCost uint64
	query = "SELECT cost FROM upgrades WHERE upgrade_level=? AND class=?"
	err = db.QueryExecutor(query, &nextCost, playerCarUpgrades.Tires+2, carClassDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	upgradeResponse := response.CarUpgradeResponse{
		UpgradedPart: "tires",
		Grip:         (playerCarStats.Grip + uint64(utils.UPGRADE_GRIP)),
		OVR:          ovr,
		Coins:        playerDetails.Coins,
	}
	if isUpgradable {
		upgradeResponse.Upgradable = true
		upgradeResponse.NextLevel = playerCarUpgrades.Tires + 2
		upgradeResponse.NextCost = nextCost
	} else {
		upgradeResponse.Upgradable = false
	}

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, upgradeResponse, ctx)

}

func UpgradeTransmissionService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	var upgradeCost uint64
	isUpgradable := true

	playerDetails, playerCarStats, playerCarUpgrades, carClassDetails, maxUpgradeLevel, classRating, err := utils.UpgradeData(playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("car class details is", carClassDetails)
	if playerCarUpgrades.Transmission+1 == maxUpgradeLevel {
		isUpgradable = false
	}
	//check for max upgrade
	if playerCarUpgrades.Transmission == (maxUpgradeLevel) {

		//part cannnot be upgraded further
		response.ShowResponse(utils.UPGRADE_REACHED_MAX_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check for players money
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if playerDetails.Coins < upgradeCost {
		//player donot have enough coin to buy the upgrade
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	playerDetails.Coins -= upgradeCost

	query := "UPDATE players SET coins=? WHERE player_id=?"
	tx.Exec(query, playerDetails.Coins, playerId)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//upgrade engine,power and overall rating of the car in player car stats and upgrade

	query = "UPDATE player_car_upgrades SET transmission=? WHERE player_id=? AND car_id=? "
	err = tx.Exec(query, playerCarUpgrades.Transmission+1, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var ovr = utils.RoundFloat(utils.CalculateOVR(classRating.ORMultiplier, float64(playerCarStats.Power), float64(playerCarStats.Grip), float64(playerCarStats.Weight)), 2)

	query = "UPDATE player_cars_stats set shift_time=? , ovr=? WHERE player_id=? AND car_id=?"
	err = tx.Exec(query, (playerCarStats.ShiftTime + utils.UPGRADE_SHIFT_TIME), ovr, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if err := tx.Commit().Error; err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	var nextCost uint64
	query = "SELECT cost FROM upgrades WHERE upgrade_level=? AND class=?"
	err = db.QueryExecutor(query, &nextCost, playerCarUpgrades.Transmission+2, carClassDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	upgradeResponse := response.CarUpgradeResponse{
		UpgradedPart: "transmission",
		Shift_Time:   (playerCarStats.ShiftTime + utils.UPGRADE_SHIFT_TIME),
		OVR:          ovr,
		Coins:        playerDetails.Coins,
	}
	if isUpgradable {
		upgradeResponse.Upgradable = true
		upgradeResponse.NextLevel = playerCarUpgrades.Transmission + 2
		upgradeResponse.NextCost = nextCost
	} else {
		upgradeResponse.Upgradable = false
	}

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, upgradeResponse, ctx)

}

func UpgradeNitrousService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {
	var upgradeCost uint64
	//isUpgradable := true

	playerDetails, _, playerCarUpgrades, carClassDetails, maxUpgradeLevel, _, err := utils.UpgradeData(playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	fmt.Println("car class details is", carClassDetails)
	//check for max upgrade
	if playerCarUpgrades.Transmission == (maxUpgradeLevel) {

		//part cannnot be upgraded further
		response.ShowResponse(utils.UPGRADE_REACHED_MAX_LEVEL, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check for players money
	tx := db.BeginTransaction()
	if tx.Error != nil {
		response.ShowResponse(tx.Error.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if playerDetails.Coins < upgradeCost {
		//player donot have enough coin to buy the upgrade
		response.ShowResponse(utils.NOT_ENOUGH_COINS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	playerDetails.Coins -= upgradeCost

	query := "UPDATE players SET coins=? WHERE player_id=?"
	tx.Exec(query, playerDetails.Coins, playerId)
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//upgrade engine,power and overall rating of the car in player car stats and upgrade

	query = "UPDATE player_car_upgrades SET nitrous=? WHERE player_id=? AND car_id=? "
	err = tx.Exec(query, playerCarUpgrades.Nitrous+1, playerId, upgradeRequest.CarId).Error
	if err != nil {
		tx.Rollback()
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	if err := tx.Commit().Error; err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// var nextCost uint64
	// query = "SELECT cost FROM upgrades WHERE upgrade_level=? AND class=?"
	// err = db.QueryExecutor(query, &nextCost, playerCarUpgrades.Engine+2, carClassDetails)
	// if err != nil {
	// 	response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
	// 	return
	// }
	// upgradeResponse := response.CarUpgradeResponse{
	// 	UpgradedPart: "engine",
	// 	Power:        (playerCarStats.Power + utils.UPGRADE_POWER),
	// 	OVR:          ovr,
	// 	Coins:        playerDetails.Coins,
	// }
	// if isUpgradable {
	// 	upgradeResponse.Upgradable = true
	// 	upgradeResponse.NextLevel = playerCarUpgrades.Engine + 2
	// 	upgradeResponse.NextCost = nextCost
	// } else {
	// 	upgradeResponse.Upgradable = false
	// }

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func GetCarStatsService(ctx *gin.Context, playerId string, getReq request.CarRequest) {
	var statsResponse response.CarStatResponse
	query := "SELECT power,grip,shift_time,weight,ovr,durability,nitrous_time FROM player_cars_stats WHERE player_id=? AND car_id=?"
	err := db.QueryExecutor(query, &statsResponse, playerId, getReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var nitrousLevel uint64
	query = "SELECT nitrous FROM player_car_upgrades WHERE player_id=? AND car_id=?"
	err = db.QueryExecutor(query, &nitrousLevel, playerId, getReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//calculate temporary power of the car after for the nitrous
	tempPower := (statsResponse.Power + (10*((statsResponse.Power)+(nitrousLevel*10)))/100)
	statsResponse.TempPower = tempPower

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, statsResponse, ctx)

}
