package car

import (
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func UpgradeEngineService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	var player model.Player
	var Car model.Car
	var carCurrentUpgrades model.PlayerCarUpgrades
	var engineDetails model.Engine

	query := "SELECT * FROM player_car_upgrades WHERE car_id = ? AND player_id = ? "
	err := db.QueryExecutor(query, &carCurrentUpgrades, upgradeRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if utils.AlreadyAtMax(carCurrentUpgrades.Engine) {
		//already at highest level
		response.ShowResponse(utils.PARTS_CANNOT_BE_UPGRADED, utils.HTTP_BAD_REQUEST, utils.SUCCESS, nil, ctx)
		return
	}

	//get the class corresponding to car_id

	err = db.FindById(&Car, upgradeRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//get the price of current engine upgrade

	query = "SELECT * FROM engines WHERE car_class=? AND level=?"
	err = db.QueryExecutor(query, &engineDetails, Car.Class, carCurrentUpgrades.Engine+1)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//logic for money comparision

	err = db.FindById(&player, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check the payment mode
	if upgradeRequest.PaymentMode == "Cash" {

		if player.Cash < int64(engineDetails.CashPrice) {

			response.ShowResponse(utils.CASH_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {

			//cut the cash
			//update the player.cash
			player.Cash = player.Cash - int64(engineDetails.CashPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
		}
	} else {

		if player.Coins < int64(engineDetails.CoinPrice) {
			response.ShowResponse(utils.COINS_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {
			//cut the coins
			player.Coins = player.Coins - int64(engineDetails.CoinPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}

		}
	}

	//var carCurrentUpgrades model.PlayerCarUpgrades

	//now do the upgradation of engine
	query = "UPDATE player_car_upgrades SET engine= ? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, carCurrentUpgrades.Engine+1, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the car stats for the player
	var playerCarStats model.PlayerCarsStats

	query = "SELECT * FROM player_cars_stats WHERE player_id =? and car_id =?"

	err = db.QueryExecutor(query, &playerCarStats, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(utils.STATS_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increasing power
	query = "UPDATE player_cars_stats SET power=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.Power + engineDetails.Power), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the overAll Rating of the car
	query = "UPDATE player_cars_stats SET ovr=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.OVR + engineDetails.OVR), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(&playerCarStats)

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func UpgradeTurboService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then intake level
	var player model.Player
	var Car model.Car
	var carCurrentUpgrades model.PlayerCarUpgrades
	var turboDetails model.Turbo

	query := "SELECT * FROM player_car_upgrades WHERE car_id = ? AND player_id = ? "
	err := db.QueryExecutor(query, &carCurrentUpgrades, upgradeRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if utils.AlreadyAtMax(carCurrentUpgrades.Engine) {
		//already at highest level
		response.ShowResponse(utils.PARTS_CANNOT_BE_UPGRADED, utils.HTTP_BAD_REQUEST, utils.SUCCESS, nil, ctx)
		return
	}

	//get the class corresponding to car_id

	err = db.FindById(&Car, upgradeRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//get the price of current turbo upgrade

	query = "SELECT * FROM turbo WHERE car_class=? AND level=?"
	err = db.QueryExecutor(query, &turboDetails, Car.Class, carCurrentUpgrades.Engine+1)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//logic for money comparision

	err = db.FindById(&player, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check the payment mode
	if upgradeRequest.PaymentMode == "Cash" {

		if player.Cash < int64(turboDetails.CashPrice) {

			response.ShowResponse(utils.CASH_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {

			//cut the cash
			//update the player.cash
			player.Cash = player.Cash - int64(turboDetails.CashPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
		}
	} else {

		if player.Coins < int64(turboDetails.CoinPrice) {
			response.ShowResponse(utils.COINS_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {
			//cut the coins
			player.Coins = player.Coins - int64(turboDetails.CoinPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}

		}
	}

	//var carCurrentUpgrades model.PlayerCarUpgrades

	//now do the upgradation of engine
	query = "UPDATE player_car_upgrades SET turbo= ? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, carCurrentUpgrades.Turbo+1, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the car stats for the player
	var playerCarStats model.PlayerCarsStats

	query = "SELECT * FROM player_cars_stats WHERE player_id =? and car_id =?"

	err = db.QueryExecutor(query, &playerCarStats, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(utils.STATS_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increasing power
	query = "UPDATE player_cars_stats SET power=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.Power + turboDetails.Power), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the overAll Rating of the car
	query = "UPDATE player_cars_stats SET ovr=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.OVR + turboDetails.OVR), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(&playerCarStats)

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func UpgradeIntakeService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then intake level
	var player model.Player
	var Car model.Car
	var carCurrentUpgrades model.PlayerCarUpgrades
	var intakeDetails model.Intake

	query := "SELECT * FROM player_car_upgrades WHERE car_id = ? AND player_id = ? "
	err := db.QueryExecutor(query, &carCurrentUpgrades, upgradeRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if utils.AlreadyAtMax(carCurrentUpgrades.Engine) {
		//already at highest level
		response.ShowResponse(utils.PARTS_CANNOT_BE_UPGRADED, utils.HTTP_BAD_REQUEST, utils.SUCCESS, nil, ctx)
		return
	}

	//get the class corresponding to car_id

	err = db.FindById(&Car, upgradeRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//get the price of current intake upgrade

	query = "SELECT * FROM intake WHERE car_class=? AND level=?"
	err = db.QueryExecutor(query, &intakeDetails, Car.Class, carCurrentUpgrades.Engine+1)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//logic for money comparision

	err = db.FindById(&player, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check the payment mode
	if upgradeRequest.PaymentMode == "Cash" {

		if player.Cash < int64(intakeDetails.CashPrice) {

			response.ShowResponse(utils.CASH_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {

			//cut the cash
			//update the player.cash
			player.Cash = player.Cash - int64(intakeDetails.CashPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
		}
	} else {

		if player.Coins < int64(intakeDetails.CoinPrice) {
			response.ShowResponse(utils.COINS_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {
			//cut the coins
			player.Coins = player.Coins - int64(intakeDetails.CoinPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}

		}
	}

	//var carCurrentUpgrades model.PlayerCarUpgrades

	//now do the upgradation of engine
	query = "UPDATE player_car_upgrades SET intake= ? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, carCurrentUpgrades.Intake+1, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the car stats for the player
	var playerCarStats model.PlayerCarsStats

	query = "SELECT * FROM player_cars_stats WHERE player_id =? and car_id =?"

	err = db.QueryExecutor(query, &playerCarStats, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(utils.STATS_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increasing power
	query = "UPDATE player_cars_stats SET power=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.Power + intakeDetails.Power), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the overAll Rating of the car
	query = "UPDATE player_cars_stats SET ovr=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.OVR + intakeDetails.OVR), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(&playerCarStats)

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func UpgradeNitrousService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then nitrous level
	var player model.Player
	var Car model.Car
	var carCurrentUpgrades model.PlayerCarUpgrades
	var nitrousDetails model.Nitrous

	query := "SELECT * FROM player_car_upgrades WHERE car_id = ? AND player_id = ? "
	err := db.QueryExecutor(query, &carCurrentUpgrades, upgradeRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if utils.AlreadyAtMax(carCurrentUpgrades.Engine) {
		//already at highest level
		response.ShowResponse(utils.PARTS_CANNOT_BE_UPGRADED, utils.HTTP_BAD_REQUEST, utils.SUCCESS, nil, ctx)
		return
	}

	//get the class corresponding to car_id

	err = db.FindById(&Car, upgradeRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//get the price of current nitrous upgrade

	query = "SELECT * FROM nitrous WHERE car_class=? AND level=?"
	err = db.QueryExecutor(query, &nitrousDetails, Car.Class, carCurrentUpgrades.Engine+1)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//logic for money comparision

	err = db.FindById(&player, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check the payment mode
	if upgradeRequest.PaymentMode == "Cash" {

		if player.Cash < int64(nitrousDetails.CashPrice) {

			response.ShowResponse(utils.CASH_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {

			//cut the cash
			//update the player.cash
			player.Cash = player.Cash - int64(nitrousDetails.CashPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
		}
	} else {

		if player.Coins < int64(nitrousDetails.CoinPrice) {
			response.ShowResponse(utils.COINS_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {
			//cut the coins
			player.Coins = player.Coins - int64(nitrousDetails.CoinPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}

		}
	}

	//var carCurrentUpgrades model.PlayerCarUpgrades

	//now do the upgradation of engine
	query = "UPDATE player_car_upgrades SET nitrous= ? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, carCurrentUpgrades.Nitrous+1, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the car stats for the player
	var playerCarStats model.PlayerCarsStats

	query = "SELECT * FROM player_cars_stats WHERE player_id =? and car_id =?"

	err = db.QueryExecutor(query, &playerCarStats, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(utils.STATS_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increasing power
	query = "UPDATE player_cars_stats SET nitrous_time=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (int64(playerCarStats.NitrousTime) + nitrousDetails.NitrousTime), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the overAll Rating of the car
	query = "UPDATE player_cars_stats SET ovr=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.OVR + nitrousDetails.OVR), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(&playerCarStats)

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func UpgradeBodyService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then nitrous level
	var player model.Player
	var Car model.Car
	var carCurrentUpgrades model.PlayerCarUpgrades
	var partDetail model.Body

	query := "SELECT * FROM player_car_upgrades WHERE car_id = ? AND player_id = ? "
	err := db.QueryExecutor(query, &carCurrentUpgrades, upgradeRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if utils.AlreadyAtMax(carCurrentUpgrades.Engine) {
		//already at highest level
		response.ShowResponse(utils.PARTS_CANNOT_BE_UPGRADED, utils.HTTP_BAD_REQUEST, utils.SUCCESS, nil, ctx)
		return
	}

	//get the class corresponding to car_id

	err = db.FindById(&Car, upgradeRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//get the price of current nitrous upgrade

	query = "SELECT * FROM body WHERE car_class=? AND level=?"
	err = db.QueryExecutor(query, &partDetail, Car.Class, carCurrentUpgrades.Engine+1)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//logic for money comparision

	err = db.FindById(&player, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check the payment mode
	if upgradeRequest.PaymentMode == "Cash" {

		if player.Cash < int64(partDetail.CashPrice) {

			response.ShowResponse(utils.CASH_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {

			//cut the cash
			//update the player.cash
			player.Cash = player.Cash - int64(partDetail.CashPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
		}
	} else {

		if player.Coins < int64(partDetail.CoinPrice) {
			response.ShowResponse(utils.COINS_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {
			//cut the coins
			player.Coins = player.Coins - int64(partDetail.CoinPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}

		}
	}

	//var carCurrentUpgrades model.PlayerCarUpgrades

	//now do the upgradation of engine
	query = "UPDATE player_car_upgrades SET body= ? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, carCurrentUpgrades.Body+1, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the car stats for the player
	var playerCarStats model.PlayerCarsStats

	query = "SELECT * FROM player_cars_stats WHERE player_id =? and car_id =?"

	err = db.QueryExecutor(query, &playerCarStats, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(utils.STATS_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increasing power
	query = "UPDATE player_cars_stats SET grip=? AND weight=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.Grip + partDetail.Grip), (playerCarStats.Weight + partDetail.Weight), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the overAll Rating of the car
	query = "UPDATE player_cars_stats SET ovr=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.OVR + partDetail.OVR), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(&playerCarStats)

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func UpgradeTiresService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then nitrous level
	var player model.Player
	var Car model.Car
	var carCurrentUpgrades model.PlayerCarUpgrades
	var partDetail model.Tires

	query := "SELECT * FROM player_car_upgrades WHERE car_id = ? AND player_id = ? "
	err := db.QueryExecutor(query, &carCurrentUpgrades, upgradeRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if utils.AlreadyAtMax(carCurrentUpgrades.Engine) {
		//already at highest level
		response.ShowResponse(utils.PARTS_CANNOT_BE_UPGRADED, utils.HTTP_BAD_REQUEST, utils.SUCCESS, nil, ctx)
		return
	}

	//get the class corresponding to car_id

	err = db.FindById(&Car, upgradeRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//get the price of current nitrous upgrade

	query = "SELECT * FROM body WHERE car_class=? AND level=?"
	err = db.QueryExecutor(query, &partDetail, Car.Class, carCurrentUpgrades.Engine+1)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//logic for money comparision

	err = db.FindById(&player, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check the payment mode
	if upgradeRequest.PaymentMode == "Cash" {

		if player.Cash < int64(partDetail.CashPrice) {

			response.ShowResponse(utils.CASH_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {

			//cut the cash
			//update the player.cash
			player.Cash = player.Cash - int64(partDetail.CashPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
		}
	} else {

		if player.Coins < int64(partDetail.CoinPrice) {
			response.ShowResponse(utils.COINS_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {
			//cut the coins
			player.Coins = player.Coins - int64(partDetail.CoinPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}

		}
	}

	//var carCurrentUpgrades model.PlayerCarUpgrades

	//now do the upgradation of engine
	query = "UPDATE player_car_upgrades SET tires= ? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, carCurrentUpgrades.Tires+1, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the car stats for the player
	var playerCarStats model.PlayerCarsStats

	query = "SELECT * FROM player_cars_stats WHERE player_id =? and car_id =?"

	err = db.QueryExecutor(query, &playerCarStats, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(utils.STATS_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increasing power
	query = "UPDATE player_cars_stats SET grip=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.Grip + partDetail.Grip), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the overAll Rating of the car
	query = "UPDATE player_cars_stats SET ovr=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.OVR + partDetail.OVR), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(&playerCarStats)

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func UpgradeTransmissionService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then nitrous level
	var player model.Player
	var Car model.Car
	var carCurrentUpgrades model.PlayerCarUpgrades
	var partDetail model.Transmission

	query := "SELECT * FROM player_car_upgrades WHERE car_id = ? AND player_id = ? "
	err := db.QueryExecutor(query, &carCurrentUpgrades, upgradeRequest.CarId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if utils.AlreadyAtMax(carCurrentUpgrades.Engine) {
		//already at highest level
		response.ShowResponse(utils.PARTS_CANNOT_BE_UPGRADED, utils.HTTP_BAD_REQUEST, utils.SUCCESS, nil, ctx)
		return
	}

	//get the class corresponding to car_id

	err = db.FindById(&Car, upgradeRequest.CarId, "car_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//get the price of current nitrous upgrade

	query = "SELECT * FROM body WHERE car_class=? AND level=?"
	err = db.QueryExecutor(query, &partDetail, Car.Class, carCurrentUpgrades.Engine+1)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//logic for money comparision

	err = db.FindById(&player, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//check the payment mode
	if upgradeRequest.PaymentMode == "Cash" {

		if player.Cash < int64(partDetail.CashPrice) {

			response.ShowResponse(utils.CASH_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {

			//cut the cash
			//update the player.cash
			player.Cash = player.Cash - int64(partDetail.CashPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
		}
	} else {

		if player.Coins < int64(partDetail.CoinPrice) {
			response.ShowResponse(utils.COINS_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {
			//cut the coins
			player.Coins = player.Coins - int64(partDetail.CoinPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}

		}
	}

	//var carCurrentUpgrades model.PlayerCarUpgrades

	//now do the upgradation of engine
	query = "UPDATE player_car_upgrades SET transmission= ? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, carCurrentUpgrades.Transmission+1, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the car stats for the player
	var playerCarStats model.PlayerCarsStats

	query = "SELECT * FROM player_cars_stats WHERE player_id =? and car_id =?"

	err = db.QueryExecutor(query, &playerCarStats, playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(utils.STATS_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increasing power
	query = "UPDATE player_cars_stats SET shift_time=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.ShiftTime + partDetail.ShiftTime), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//increase the overAll Rating of the car
	query = "UPDATE player_cars_stats SET ovr=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.OVR + partDetail.OVR), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(&playerCarStats)

	response.ShowResponse(utils.UPGRADE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}
