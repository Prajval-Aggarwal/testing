package player

import (
	"main/server/db"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

func GetPlayerDetails(ctx *gin.Context, playerId string) {

	var playerResponse response.PlayerResposne
	query := `SELECT
		p.player_id,
		p.player_name,
		p.level,
		p.xp,
		p.role,
		p.email,
		p.coins,
		p.cash,
		p.repair_parts
    COUNT(oc.car_id) AS CarsOwned,
    COUNT(og.garage_id) AS GaragesOwned,
    prh.distance_traveled,
    prh.shd_won AS ShowDownWon,
		CASE
			WHEN prh.total_shd_played > 0 THEN prh.shd_won / prh.total_shd_played
			ELSE 0
		END AS ShowDownWinRatio,
		prh.td_won AS TakeDownWon,
		CASE
			WHEN prh.total_td_played > 0 THEN prh.td_won / prh.total_td_played
			ELSE 0
		END AS TakeDownWinRatio
	FROM players p
	LEFT JOIN owned_cars oc ON oc.player_id = p.player_id
	LEFT JOIN owned_garages og ON og.player_id = p.player_id
	LEFT JOIN player_race_histories prh ON prh.player_id = p.player_id
	WHERE p.player_id = ?
	GROUP BY
		p.player_id,
		p.player_name,
		p.level,
		p.role,
		p.email,
		p.coins,
		p.cash,
		prh.distance_traveled,
		prh.shd_won,
		prh.total_shd_played,
		prh.td_won,
		prh.total_td_played;`

	playerResponse = db.ResponseQuery(query, playerId)

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, playerResponse, ctx)
}
