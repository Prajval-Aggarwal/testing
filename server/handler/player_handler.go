package handler

import (
	"main/server/response"
	"main/server/services/player"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Get player details
// @Description Equip a car for a player
// @Tags Player
// @Accept json
// @Produce json
// @Param Authorization header string true "Access token"
// @Success 200 {object} response.Success "Car equipped successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure  401 {object} response.Success "Unauthorised"
// @Router /player-details [get]
func GetPlayerDetailsHandler(ctx *gin.Context) {
	playerId, exists := ctx.Get(utils.PLAYER_ID)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	player.GetPlayerDetails(ctx, playerId.(string))

}
