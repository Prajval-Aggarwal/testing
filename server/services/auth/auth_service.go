package auth

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"main/server/db"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/services/token"
	"main/server/utils"

	"github.com/google/uuid"
)

func GuestLoginService(ctx *gin.Context, guestLoginReuqest request.GuestLoginRequest) {
	//generate a token

	//Generate access and refresh token

	if guestLoginReuqest.Token == "" {
		accessTokenexpirationTime := time.Now().Add(5 * time.Hour) //5 minute expiration time for access token

		fmt.Println("accessTokenExpiration time is", accessTokenexpirationTime)
		playeruuid := uuid.New()
		fmt.Println("unique player id is", playeruuid)

		playerRecord := model.Player{
			PlayerId:   playeruuid,
			PlayerName: guestLoginReuqest.PlayerName,
			Role:       "player",
			OS:         int64(guestLoginReuqest.OS),
			Coins:      10000000,
			Cash:       10000000,
			DeviceId:   guestLoginReuqest.DeviceId,
		}

		accessTokenClaims := model.Claims{
			Id:   playeruuid,
			Role: "player",
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(accessTokenexpirationTime),
			},
		}
		accessToken, err := token.GenerateToken(accessTokenClaims)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}

		//create a record in database
		err = db.CreateRecord(&playerRecord)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}

		fmt.Println("access token generated:", *accessToken)

		response.ShowResponse(utils.LOGIN_SUCCESS, utils.HTTP_OK, utils.SUCCESS, *accessToken, ctx)
	} else {
		fmt.Println("token form request is", guestLoginReuqest.Token)
		newToken, err := token.CheckExpiration(guestLoginReuqest.Token)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}
		response.ShowResponse(utils.LOGIN_SUCCESS, utils.HTTP_OK, utils.SUCCESS, *newToken, ctx)

	}

}

func LoginService(ctx *gin.Context, loginDetails request.LoginRequest) {
	var playerLogin model.Player

	//first check if the player with that email exists or not
	err := db.FindById(&playerLogin, loginDetails.Email, "email")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)

		return
	}

	accessTokenexpirationTime := time.Now().Add(5 * time.Hour) //5 minute expiration time for access token

	accessTokenClaims := model.Claims{
		Id:   playerLogin.PlayerId,
		Role: playerLogin.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(accessTokenexpirationTime),
		},
	}

	//Generate access and refresh token
	accessToken, err := token.GenerateToken(accessTokenClaims)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.LOGIN_SUCCESS, utils.HTTP_OK, utils.SUCCESS, *accessToken, ctx)

}

func UpdateEmailService(ctx *gin.Context, email request.UpdateEmailRequest, playerId string) {

	var playerDetails model.Player
	//check if the password is valid or not

	//token Decoding

	if db.RecordExist("players", email.Email, "email") {
		response.ShowResponse("Email is already attached to another player", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	if !db.RecordExist("players", playerId, "player_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}
	err := db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return

	}
	//check if the record exists of not

	query := "UPDATE players SET email=? WHERE player_id=?"
	err = db.RawExecutor(query, email.Email, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return

	}

	response.ShowResponse(utils.EMAIL_UPDATED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}
