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

	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then engine level
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
	query = "UPDATE player_car_stats SET power=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.Power + engineDetails.Power), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(playerId, upgradeRequest.CarId)

}

func UpgradeTurboService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {
	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then engine level
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
	//get the price of current engine upgrade

	query = "SELECT * FROM engines WHERE car_class=? AND level=?"
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
	query = "UPDATE player_car_stats SET power=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.Power + turboDetails.Power), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(playerId, upgradeRequest.CarId)
}

func UpgradeIntakeService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {
	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then engine level
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
	//get the price of current engine upgrade

	query = "SELECT * FROM engines WHERE car_class=? AND level=?"
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
	query = "UPDATE player_car_stats SET power=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.Power + intakeDetails.Power), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(playerId, upgradeRequest.CarId)
}

func UpgradeNitrousService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

}

func UpgradeBodyService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {
	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then engine level
	var player model.Player
	var Car model.Car
	var carCurrentUpgrades model.PlayerCarUpgrades
	var bodyDetails model.Body

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
	err = db.QueryExecutor(query, &bodyDetails, Car.Class, carCurrentUpgrades.Engine+1)
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

		if player.Cash < int64(bodyDetails.CashPrice) {

			response.ShowResponse(utils.CASH_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {

			//cut the cash
			//update the player.cash
			player.Cash = player.Cash - int64(bodyDetails.CashPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
		}
	} else {

		if player.Coins < int64(bodyDetails.CoinPrice) {
			response.ShowResponse(utils.COINS_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {
			//cut the coins
			player.Coins = player.Coins - int64(bodyDetails.CoinPrice)
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
	query = "UPDATE player_car_stats SET grip=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (playerCarStats.Grip + bodyDetails.Grip), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(playerId, upgradeRequest.CarId)
}

func UpgradeTiresService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

}

func UpgradeTransmissionService(ctx *gin.Context, upgradeRequest request.CarUpgradesRequest, playerId string) {

	//to upgrade the engine
	//check if engine is already at highest level
	//if not,Then compare the price of engine with the money player has
	//if player has money,check the current level of engine player has
	//then upgrade then engine level
	var player model.Player
	var Car model.Car
	var carCurrentUpgrades model.PlayerCarUpgrades
	var transmissionDetails model.Transmission

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
	err = db.QueryExecutor(query, &transmissionDetails, Car.Class, carCurrentUpgrades.Engine+1)
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

		if player.Cash < int64(transmissionDetails.CashPrice) {

			response.ShowResponse(utils.CASH_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {

			//cut the cash
			//update the player.cash
			player.Cash = player.Cash - int64(transmissionDetails.CashPrice)
			err := db.UpdateRecord(&player, playerId, "player_id").Error
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
		}
	} else {

		if player.Coins < int64(transmissionDetails.CoinPrice) {
			response.ShowResponse(utils.COINS_LIMIT_EXCEEDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		} else {
			//cut the coins
			player.Coins = player.Coins - int64(transmissionDetails.CoinPrice)
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
	query = "UPDATE player_car_stats SET shift_time=? WHERE player_id=? AND car_id=?"
	err = db.RawExecutor(query, (int64(playerCarStats.ShiftTime) + int64(transmissionDetails.ShiftTime)), playerId, upgradeRequest.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//after every upgrade check if the car stats has reached a certain value above which the car level upgrades automatically
	utils.UpgradeCarLevel(playerId, upgradeRequest.CarId)
}
