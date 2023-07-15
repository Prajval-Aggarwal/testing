package arena

import (
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func ChallengeArenaService(ctx *gin.Context, challengereq request.ChallengeReq, playerId string) {

	//GET THE AREA DETAILS FORM THE DATABASE
	var arenaDetails model.Arena
	err := db.FindById(&arenaDetails, challengereq.ArenaId, "arena_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	//GET THE LEVEL DETAILS FROM THE DATABASE
	var arenaLevelDetails model.ArenaReq
	err = db.FindById(&arenaLevelDetails, arenaDetails.ArenaLevel, "arena_level")
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
	var carCount int64
	query := "select count(*) from owned_cars where player_id=?"
	err = db.QueryExecutor(query, &carCount, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//CHECK FOR THE MINIMUN REQUIRMENTS FOR THE CHALLANGING THE ARENA
	if playerDetails.Level < arenaLevelDetails.PlayerLevel {
		response.ShowResponse("Minimum level required is "+fmt.Sprint(arenaLevelDetails.PlayerLevel), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	} else {
		if carCount < arenaLevelDetails.MinCarReq {
			response.ShowResponse("Minimum amount of cars required is "+fmt.Sprint(arenaLevelDetails.MinCarReq), utils.HTTP_BAD_GATEWAY, utils.FAILURE, nil, ctx)
			return
		}
	}

	//IF YES PRINT YES ELSE GIVE ERROR
	response.ShowResponse(utils.SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func EndChallengeService(ctx *gin.Context, endChallReq request.EndChallengeReq, player1_id string) {
	//playerId is the player who challenged i.e player 1
	//endChallReq.PlayerId is the player who is being challenged i.e player 2

	//player 2 owns that arena and player 1 is trying to own that arena

	//reward concept pending

	//get the win time of that arena
	var arenaDetails model.OwnedBattleArenas
	query := "SELECT * FROM owned_arenas WHERE arena_id=? AND player_id=?"
	err := db.QueryExecutor(query, &arenaDetails, endChallReq.ArenaId, endChallReq.PlayerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if endChallReq.WinTime.Before(arenaDetails.WinTime) {

		//player1 wins
		//give reward and allot arena

		//giving reward left
		wonArena := &model.OwnedBattleArenas{
			PlayerId: player1_id,
			ArenaId:  endChallReq.ArenaId,
			TimeWon:  time.Now().UTC(),
			WinTime:  endChallReq.WinTime,
			Status:   "temporary",
		}

		err := db.CreateRecord(&wonArena)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}

		time.AfterFunc(24*time.Hour, func() {
			var carId string
			query := "SELECT car_id from owned_arenas WHERE arena_id=? AND player_id=?"
			err := db.QueryExecutor(query, &carId, endChallReq.ArenaId, player1_id)

			if err != nil {
				response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
				return
			}
			if carId != "" {
				query := "UPDATE owned_arenas SET status='permanent' WHERE arena_id=? AND player_id = ?"
				err := db.RawExecutor(query, endChallReq.ArenaId, player1_id)
				if err != nil {
					response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
					return
				}
			} else {
				query := "DELETE FROM owned_arenas WHERE arena_id=? AND player_id=?"
				err := db.RawExecutor(query, endChallReq.ArenaId, player1_id)
				if err != nil {
					response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
					return
				}
			}

		})

		response.ShowResponse(utils.WON, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

	} else {
		response.ShowResponse(utils.LOSE, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

	}
}

// func AddCarArenaService(ctx *gin.Context, addCarReq request.AddCarArenaRequest, playerId string) {
// 	var count int64
// 	query := "SELECT count(*) FROM owned_arenas WHERE player_id=? AND car_id=?"
// 	err := db.QueryExecutor(query, &count, playerId, addCarReq.CarId)
// 	if err != nil {
// 		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
// 		return
// 	}
// 	if count != 0 {
// 		response.ShowResponse(utils.CAR_ALREAY_ALLOTED, utils.HTTP_BAD_REQUEST, utils.FAILED_TO_UPDATE, nil, ctx)
// 		return
// 	}

// 	//chek whether the car is bought or not
// 	var exists bool
// 	query = "SELECT EXISTS(SELECT * FROM owned_cars WHERE player_id =? AND car_id=?)"
// 	err = db.QueryExecutor(query, &exists, playerId, addCarReq.CarId)
// 	if err != nil {
// 		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
// 		return
// 	}
// 	if !exists {
// 		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
// 		return
// 	}

// 	// add the car to arena
// 	query = "UPDATE owned_arenas SET car_id =? WHERE player_id =? AND arena_id =?"
// 	err = db.RawExecutor(query, addCarReq.CarId, playerId, addCarReq.ArenaId)
// 	if err != nil {
// 		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
// 		return
// 	}

// 	response.ShowResponse(utils.CAR_ADDED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

// }

func ReplaceArenaService(ctx *gin.Context, replaceReq request.ReplaceReq, playerId string) {

	//CHECK IF THE CAR IS ALREADY ALLOTED TO OTHER ARENA OR NOT
	var count int64
	query := "SELECT count(*) FROM owned_arenas WHERE player_id=? AND car_id=?"
	err := db.QueryExecutor(query, &count, playerId, replaceReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	if count != 0 {
		response.ShowResponse(utils.CAR_ALREAY_ALLOTED, utils.HTTP_BAD_REQUEST, utils.FAILED_TO_UPDATE, nil, ctx)
		return
	}

	//check if the car is bought or not
	var exists bool
	query = "SELECT EXISTS(SELECT * FROM owned_cars WHERE player_id =? AND car_id=?)"
	err = db.QueryExecutor(query, &exists, playerId, replaceReq.CarId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	if !exists {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}
	//Replace the car
	query = "UPDATE owned_arenas SET car_id=? WHERE player_id=? AND arena_id=?"
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

func GetArenaByIdService(ctx *gin.Context, getReq request.GetArenaReq) {

}
