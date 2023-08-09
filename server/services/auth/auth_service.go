package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

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
		playeruuid := uuid.New().String()
		fmt.Println("unique player id is", playeruuid)

		playerRecord := model.Player{
			PlayerId:   playeruuid,
			PlayerName: guestLoginReuqest.PlayerName,
			Level:      1,
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
		playerRaceHist := model.PlayerRaceHistory{
			PlayerId:         playeruuid,
			DistanceTraveled: 0,
			ShdWon:           0,
			TotalShdPlayed:   0,
			TdWon:            0,
			TotalTdPlayed:    0,
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

		err = db.CreateRecord(&playerRaceHist)
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
		response.ShowResponse(utils.EMAIL_EXISTS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
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

func AdminSignUpService(ctx *gin.Context) {
	username := "Davinder"
	password := "hood@123"
	email := "davinder@yopmail.com"
	var adminDetails model.Admin
	adminDetails.Email = email
	hashedPass, err := utils.HashPassword(password)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	adminDetails.Password = *hashedPass
	adminDetails.Username = username
	err = db.CreateRecord(&adminDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

}

func AdminLoginService(ctx *gin.Context, adminLoginReq request.LoginRequest) {
	var adminDetails model.Admin
	if !db.RecordExist("admins", adminLoginReq.Email, "email") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}
	err := db.FindById(&adminDetails, adminLoginReq.Email, "email")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(adminDetails.Password), []byte(adminLoginReq.Password))
	if err != nil {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	accessTokenexpirationTime := time.Now().Add(5 * time.Hour) //5 hour expiration time for access token

	accessTokenClaims := model.Claims{
		Id:   adminDetails.Id,
		Role: "admin",
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

	//create a record in session table for the admin
	session := model.Session{
		Token: *accessToken,
	}
	err = db.CreateRecord(&session)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, ctx, nil)
		return
	}
	response.ShowResponse(utils.LOGIN_SUCCESS, utils.HTTP_OK, utils.SUCCESS, *accessToken, ctx)

}

func ForgotPassService(ctx *gin.Context, forgotPassRequest request.ForgotPassRequest) {
	expirationTime := time.Now().Add(time.Minute * 5)
	var adminDetails model.Admin

	// finding the player email
	err := db.FindById(&adminDetails, forgotPassRequest.Email, "email")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	} else if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//Creating reset token payload and generating token form it
	resetClaims := model.Claims{
		Id: adminDetails.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	tokenString, err := token.GenerateToken(resetClaims)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Creating reset session for reseting the password
	resetSession := model.ResetSession{
		Id:    adminDetails.Id,
		Token: *tokenString,
	}
	err = db.CreateRecord(&resetSession)
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		query := "UPDATE reset_sessions SET token = ? WHERE id=?"
		err = db.RawExecutor(query, *tokenString, adminDetails.Id)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
			return
		}
	}
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	//creating link
	link := ctx.Request.Header.Get("Origin") + "/reset-password?token=" + *tokenString

	//Sending mail on admin's email address
	err = utils.SendEmaillService(adminDetails, link)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)

		return
	}
	response.ShowResponse(utils.LINK_GENERATED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, link, ctx)

}
func ResetPasswordService(ctx *gin.Context, tokenString string, password request.UpdatePasswordRequest) {
	var resetSession model.ResetSession
	//Decoding the token and extracting require data
	claims, err := token.DecodeToken(tokenString)
	if errors.Is(err, fmt.Errorf("invalid token")) {

		//delete session in resset session
		err = db.DeleteRecord(&resetSession, claims.Id, "id")
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, err.Error(), ctx)
			return
		}
	}
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = db.FindById(&resetSession, claims.Id, "id")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.ShowResponse("Invalid session", utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	} else if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	if resetSession.Token != tokenString {
		response.ShowResponse(utils.FORBIDDEN_REQUEST, utils.HTTP_FORBIDDEN, utils.FAILURE, nil, ctx)
		return
	}
	// Reusing he updatepasswordservice
	UpdatePasswordService(ctx, password, claims.Id)

	//After sucessfully reseting the password deleteing reset session of the player
	err = db.DeleteRecord(&resetSession, claims.Id, "id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, err.Error(), ctx)
		return
	}

	//truncate the sessions table
	query := "TRUNCATE TABLE sessions"
	err = db.RawExecutor(query)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, err.Error(), ctx)
		return
	}

}

func UpdatePasswordService(ctx *gin.Context, password request.UpdatePasswordRequest, playerID string) {

	var adminDetails model.Admin
	//Finding the admin
	err := db.FindById(&adminDetails, playerID, "id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Password validity check
	err = utils.IsPassValid(password.Password)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//Hashing the new password
	hashedPass, err := utils.HashPassword(password.Password)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	adminDetails.Password = *hashedPass

	//Updating players new password
	err = db.UpdateRecord(&adminDetails, playerID, "id").Error
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	response.ShowResponse(utils.PASSWORD_UPDATE_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func GetAdminService(ctx *gin.Context) {
	var admins = []model.Admin{}
	var dataresp response.DataResponse
	query := "SELECT * FROM admins"

	err := db.QueryExecutor(query, &admins)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	var totalCount int
	countQuery := "SELECT COUNT(*) FROM admins"
	err = db.QueryExecutor(countQuery, &totalCount)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}
	dataresp.TotalCount = totalCount
	dataresp.Data = admins

	response.ShowResponse(utils.DATA_FETCH_SUCCESS, utils.HTTP_OK, utils.SUCCESS, dataresp, ctx)
}
