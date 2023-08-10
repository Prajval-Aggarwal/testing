package arena

import (
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func EndChallengeService(ctx *gin.Context, endChallReq request.EndChallengeReq, playerId string) {

	//playerId is the first player who challenged
	//endChallReq.playerId is the second player who is being challenged

	if endChallReq.ArenaId == "" {

		var opponentTime time.Time
		query := "SELECT time_win FROM race_rewards WHERE player_id=? AND arena_id=?"
		err := db.QueryExecutor(query, &opponentTime, endChallReq.PlayerId, endChallReq.ArenaId)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}
		var rewards model.Rewards
		win := false

		query = "SELECT * FROM rewards WHERE id=? AND status=?"
		if endChallReq.WinTime.Compare(opponentTime) == -1 {

			win = true
			//check the type off the race and allot the rewards to the player

			err = db.QueryExecutor(query, endChallReq.RaceId, "win")
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
			err = EarnedRewards(playerId, ctx, rewards)
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}

			//pending

			// playerRaceRecord := model.ArenaRaceRecord{
			// 	PlayerId: playerId,
			// 	ArenaId:  endChallReq.ArenaId,
			// 	Time:     endChallReq.WinTime,
			// 	Result:   "",
			// }

			response.ShowResponse(utils.WON, utils.HTTP_OK, utils.SUCCESS, rewards, ctx)
			//player wins
		} else {
			//player looses
			win = false
			err = db.QueryExecutor(query, endChallReq.RaceId, "lost")
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}

			//get player details
			var playerDetails model.Player
			err = db.FindById(&playerDetails, playerId, "player_id")
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
			var carCount uint64
			query := "select count(*) from owned_cars where player_id=?"
			err = db.QueryExecutor(query, &carCount, playerId)
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
			//give rewrads to player
			err = EarnedRewards(playerId, ctx, rewards)
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
			response.ShowResponse(utils.LOSE, utils.HTTP_OK, utils.SUCCESS, rewards, ctx)
		}

		//update player race history

		err = UpdatePlayerRaceHistory(playerId, ctx, endChallReq, win)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}

	} else {
		//player is taking challenge in arena
		var opponentTime time.Time
		query := "SELECT time_win FROM race_rewards WHERE player_id=? AND arena_id=?"
		err := db.QueryExecutor(query, &opponentTime, endChallReq.PlayerId, endChallReq.ArenaId)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}

		if endChallReq.WinTime.Compare(opponentTime) == -1 {
			//player wins the a series in arena
			//add the count to arenaRaceWins
			var exists bool
			query := "SELECT EXISTS (SELECT * FROM arena_seriess WHERE arena_id=? AND player_id=?)"
			err := db.QueryExecutor(query, &exists, endChallReq.ArenaId, playerId)
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
			if exists {
				//increment the number of wins
				query := "UPDATE arena_seriess SET win_streak=win_streak+1 WHERE  arena_id=? AND player_id=?"
				err := db.RawExecutor(query, endChallReq.ArenaId, playerId)
				if err != nil {
					response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
					return
				}

				//check if the player is eligible for arena reward or not
				var arenaSeries model.ArenaSeries
				query = "SELECT * FROM arena_series WHERE arena_id=? AND player_id=?"
				err = db.QueryExecutor(query, &arenaSeries, endChallReq.ArenaId, playerId)
				if err != nil {
					response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
					return
				}

				var arenaDetails model.Arena
				err = db.FindById(&arenaDetails, endChallReq.ArenaId, "arena_id")
				if err != nil {
					response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
					return
				}

				//change variable name
				var test model.RaceTypes
				query = "SELECT * FROM race_types WHERE difficulty=? AND race_name='arena'"
				err = db.QueryExecutor(query, &test, arenaDetails.ArenaLevel)
				if err != nil {
					response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
					return
				}

				if arenaSeries.WinStreak == test.RaceSeries {
					//player won the arena

					playerArena := model.OwnedBattleArenas{
						PlayerId: playerId,
						ArenaId:  endChallReq.ArenaId,
						WinTime:  endChallReq.WinTime,
					}
					err := db.CreateRecord(&playerArena)
					if err != nil {
						response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
						return
					}

					var reward1, reward2 model.Rewards
					query = "SELECT * FROM rewards WHERE id=? AND status=?"
					err = db.QueryExecutor(query, reward1, test.RaceId, "win")
					if err != nil {
						response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
						return
					}
					err = EarnedRewards(playerId, ctx, reward1)
					if err != nil {
						response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
						return
					}

					err = db.QueryExecutor(query, reward2, endChallReq.RaceId, "win")
					if err != nil {
						response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
						return
					}
					err = EarnedRewards(playerId, ctx, reward2)
					if err != nil {
						response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
						return
					}

					totalRewards := struct {
						Reward1 model.Rewards
						Reward2 model.Rewards
					}{
						Reward1: reward1,
						Reward2: reward2,
					}
					//give both rewards arena and takedown
					response.ShowResponse(utils.WON, utils.HTTP_OK, utils.SUCCESS, totalRewards, ctx)

					//add a 24 hour timer after the arena is won
					///if after the 24 hour there is no entery in carSlots table then the arebna will be given back to the AI
					time.AfterFunc(24*time.Hour, func() {
						count := 0
						query := "SELECT count(*) FROM car_slots WHERE player_id=? AND arena_id=?"
						err = db.QueryExecutor(query, &count, playerId, endChallReq.ArenaId)
						if err != nil {
							response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
							return
						}
						if count == 0 {
							// give the garage back to AI
							query := "UPDATE owned_battle_arenas SET player_id=? WHERE arena_id=? AND player_id=?"
							err = db.RawExecutor(query, utils.AI, endChallReq.ArenaId, playerId)
							if err != nil {
								response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
								return
							}
						}
					})

				} else {

					//player won the challnege but not the arena
					var reward model.Rewards
					query = "SELECT * FROM rewards WHERE id=? AND status=?"
					err = db.QueryExecutor(query, endChallReq.RaceId, "win")
					if err != nil {
						response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
						return
					}
					err = EarnedRewards(playerId, ctx, reward)
					if err != nil {
						response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
						return
					}
					response.ShowResponse(utils.WON, utils.HTTP_OK, utils.SUCCESS, reward, ctx)

				}

			} else {
				//create a record and set the initail win to 1
				arenaSeriesRecord := model.ArenaSeries{
					ArenaId:   endChallReq.ArenaId,
					PlayerId:  playerId,
					WinStreak: 1,
				}
				err := db.CreateRecord(&arenaSeriesRecord)
				if err != nil {
					response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
					return
				}
			}

		} else {
			//player Lost
			var reward model.Rewards
			query = "SELECT * FROM rewards WHERE id=? AND status=?"
			err = db.QueryExecutor(query, endChallReq.RaceId, "lost")
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
			err = EarnedRewards(playerId, ctx, reward)
			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
				return
			}
			response.ShowResponse(utils.WON, utils.HTTP_OK, utils.SUCCESS, reward, ctx)
		}

	}
}

func UpdatePlayerRaceHistory(playerId string, ctx *gin.Context, endChallReq request.EndChallengeReq, win bool) error {

	//get player race history
	var playerRaceHistory model.PlayerRaceHistory
	err := db.FindById(&playerRaceHistory, playerId, "player_id")
	if err != nil {
		return err
	}

	//get the details of the race type
	var raceType model.RaceTypes
	err = db.FindById(&raceType, endChallReq.RaceId, "race_id")
	if err != nil {
		return err
	}

	//update the details
	playerRaceHistory.DistanceTraveled += raceType.RaceLength
	if raceType.RaceName == "showdowns" {
		playerRaceHistory.TotalShdPlayed += 1
		if win {
			playerRaceHistory.ShdWon += 1
		}
	}
	if raceType.RaceName == "takedowns" {
		playerRaceHistory.TotalTdPlayed += 1
		if win {
			playerRaceHistory.TdWon += 1
		}
	}

	err = db.UpdateRecord(&playerRaceHistory, playerId, "player_id").Error
	if err != nil {
		return err
	}
	return nil
}

func EarnedRewards(playerId string, ctx *gin.Context, rewards model.Rewards) error {

	//get player details
	var playerDetails model.Player
	err := db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		return err
	}
	//begin transaction
	tx := db.BeginTransaction()
	if tx.Error != nil {
		return err
	}

	playerDetails.Coins += rewards.Coins
	playerDetails.Cash += rewards.Cash
	playerDetails.RepairParts += rewards.RepairParts
	playerDetails.XP += rewards.XPGained

	err = db.UpdateRecord(&playerDetails, playerId, "player_id").Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

func AddCarToSlotService(ctx *gin.Context, addCarReq request.AddCarArenaRequest, playerId string) {
	// Check if the car is bought by the player
	query := "SELECT EXISTS(SELECT * FROM owned_cars WHERE player_id = ? AND car_id = ?)"
	if !utils.IsExisting(query, playerId, addCarReq.CarId) {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Check if the player owns the arena
	query = "SELECT EXISTS(SELECT * FROM owned_battle_arenas WHERE player_id = ? AND car_id = ?)"
	if !utils.IsExisting(query, playerId, addCarReq.CarId) {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Check that it should not add more cars than required slots for the arena
	var arenaDetails model.Arena
	err := db.FindById(&arenaDetails, addCarReq.ArenaId, "arena_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	var carCount uint64
	query = "SELECT COUNT(*) FROM car_slots WHERE player_id = ? AND arena_id = ?"
	err = db.QueryExecutor(query, &carCount, playerId, addCarReq.ArenaId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Check the slot limit for the arena level and ensure it's not exceeded
	var maxSlots uint64
	switch arenaDetails.ArenaLevel {
	case uint64(utils.EASY):
		maxSlots = utils.EASY_ARENA_SLOT
	case uint64(utils.MEDIUM):
		maxSlots = utils.MEDIUM_ARENA_SLOT
	case uint64(utils.HARD):
		maxSlots = utils.HARD_ARENA_SLOT
	default:
		response.ShowResponse("Invalid arena level", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	if carCount == maxSlots {
		response.ShowResponse(utils.NO_CARS_ADDED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Create a record in the car_slots table
	carSlot := model.CarSlots{
		PlayerId: playerId,
		ArenaId:  addCarReq.ArenaId,
		CardId:   addCarReq.CarId,
	}

	err = db.CreateRecord(&carSlot)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.CAR_ADDED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func ReplaceCarService(ctx *gin.Context, replaceReq request.AddCarArenaRequest, playerId string) {
	// Check if the car is bought by the player and owned by the player
	query := "SELECT EXISTS(SELECT * FROM owned_cars WHERE player_id = ? AND car_id = ?)"
	if !utils.IsExisting(query, playerId, replaceReq.CarId) {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Check if the player owns the arena and the car is owned by the player
	query = "SELECT EXISTS(SELECT * FROM owned_battle_arenas WHERE player_id = ? AND car_id = ?)"
	if !utils.IsExisting(query, playerId, replaceReq.CarId) {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Check if the car is already allotted to any slot
	query = "SELECT COUNT(*) FROM car_slots WHERE player_id = ? AND car_id = ?"
	var count int64
	err := db.QueryExecutor(query, &count, playerId, replaceReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if count != 0 {
		response.ShowResponse(utils.CAR_ALREADY_ALLOTTED, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Replace the car in the slot
	query = "UPDATE car_slots SET car_id = ? WHERE player_id = ? AND arena_id = ?"
	err = db.RawExecutor(query, replaceReq.CarId, playerId, replaceReq.ArenaId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.CAR_REPLACED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func GetArenaService(ctx *gin.Context) {
	var getArenaResposne []response.ArenaResp
	query := "SELECT arena_id,arena_name,level FROM arenas"
	err := db.QueryExecutor(query, &getArenaResposne)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, getArenaResposne, ctx)
}
