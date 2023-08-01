package auth

import (
	"errors"
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

// GuestLoginService handles guest login requests.
func GuestLoginService(ctx *gin.Context, guestLoginRequest request.GuestLoginRequest) {
	// Check if a token is provided in the request. If not, generate a new guest token and create player records.
	if guestLoginRequest.Token == "" {
		// Generate a new player UUID and access token expiration time (48 hours from now).
		playerUUID := uuid.New().String()
		accessTokenExpirationTime := time.Now().Add(48 * time.Hour)

		// Create a new player record with default values.
		playerRecord := model.Player{
			PlayerId:    playerUUID,
			PlayerName:  guestLoginRequest.PlayerName,
			Level:       1,
			Role:        "player",
			OS:          guestLoginRequest.OS,
			Coins:       10000000,
			Cash:        10000000,
			RepairParts: 0,
			DeviceId:    guestLoginRequest.DeviceId,
		}

		// Create a new access token with player claims and expiration time.
		accessTokenClaims := model.Claims{
			Id:   playerUUID,
			Role: "player",
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(accessTokenExpirationTime),
			},
		}
		accessToken, err := token.GenerateToken(accessTokenClaims)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}

		// Create player and race history records in the database.
		err = db.CreateRecord(&playerRecord)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}
		playerRaceHist := model.PlayerRaceHistory{
			PlayerId:         playerUUID,
			DistanceTraveled: 0,
			ShdWon:           0,
			TotalShdPlayed:   0,
			TdWon:            0,
			TotalTdPlayed:    0,
		}
		err = db.CreateRecord(&playerRaceHist)
		if err != nil {
			response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}

		// Respond with the generated access token.
		response.ShowResponse(utils.LOGIN_SUCCESS, utils.HTTP_OK, utils.SUCCESS, *accessToken, ctx)
	} else {
		// If a token is provided, check if it is valid and not expired.
		newToken, err := token.CheckExpiration(guestLoginRequest.Token)
		if err != nil {
			// If the token is invalid or expired, return an error response.
			response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
			return
		}
		// If the token is valid and not expired, respond with a success message and the new token.
		response.ShowResponse(utils.LOGIN_SUCCESS, utils.HTTP_OK, utils.SUCCESS, *newToken, ctx)
	}
}

// LoginService handles player login requests.
func LoginService(ctx *gin.Context, adminLoginReq request.LoginRequest) {
	// Declare a variable to hold player login details.
	var playerLogin model.Player

	// First, check if a player with the provided email exists.
	err := db.FindById(&playerLogin, adminLoginReq.Email, "email")
	if err != nil {
		// If the player doesn't exist, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Set the expiration time for the access token (48 hours from now).
	accessTokenExpirationTime := time.Now().Add(48 * time.Hour)

	// Create the access token claims with player details and expiration time.
	accessTokenClaims := model.Claims{
		Id:   playerLogin.PlayerId,
		Role: playerLogin.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(accessTokenExpirationTime),
		},
	}

	// Generate the access token.
	accessToken, err := token.GenerateToken(accessTokenClaims)
	if err != nil {
		// If there is an error in generating the access token, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Respond with a success message and the generated access token.
	response.ShowResponse(utils.LOGIN_SUCCESS, utils.HTTP_OK, utils.SUCCESS, *accessToken, ctx)
}

func UpdateEmailService(ctx *gin.Context, email request.UpdateEmailRequest, playerId string) {
	// Check if the provided player ID exists.
	if !db.RecordExist("players", playerId, "player_id") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	var playerDetails model.Player
	// Fetch player details using the player ID.
	err := db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Check if the new email already exists for another player.
	if db.RecordExist("players", email.Email, "email") {
		response.ShowResponse(utils.EMAIL_EXISTS, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Update the email for the player in the database.
	query := "UPDATE players SET email=? WHERE player_id=?"
	err = db.RawExecutor(query, email.Email, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	response.ShowResponse(utils.EMAIL_UPDATED_SUCCESS, utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}

func AdminSignUpService(ctx *gin.Context) {
	password := "admin@$2023"
	email := "admin@gmail.com"
	var adminDetails model.Admin
	adminDetails.Email = email
	hashedPass, err := utils.HashPassword(password)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	adminDetails.Password = *hashedPass

	err = db.CreateRecord(&adminDetails)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

}

// AdminLoginService handles admin login requests.
func AdminLoginService(ctx *gin.Context, adminLoginReq request.LoginRequest) {

	// Check if the admin with the provided email exists.
	var adminDetails model.Admin
	if !db.RecordExist("admins", adminLoginReq.Email, "email") {
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	}

	// Fetch admin details using the admin email.
	err := db.FindById(&adminDetails, adminLoginReq.Email, "email")
	if err != nil {
		// If there is an error in fetching admin details, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Compare the provided password with the stored hashed password.
	err = bcrypt.CompareHashAndPassword([]byte(adminDetails.Password), []byte(adminLoginReq.Password))
	if err != nil {
		// If the password doesn't match, return an unauthorized error response.
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}

	// Set the expiration time for the access token (5 hours from now).
	accessTokenExpirationTime := time.Now().Add(5 * time.Hour)

	// Create the access token claims with admin details and expiration time.
	accessTokenClaims := model.Claims{
		Id:   adminDetails.Id,
		Role: "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(accessTokenExpirationTime),
		},
	}

	// Generate the access token.
	accessToken, err := token.GenerateToken(accessTokenClaims)
	if err != nil {
		// If there is an error in generating the access token, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Respond with a success message and the generated access token.
	response.ShowResponse(utils.LOGIN_SUCCESS, utils.HTTP_OK, utils.SUCCESS, *accessToken, ctx)
}

// ForgotPassService handles admin forgot password requests.
func ForgotPassService(ctx *gin.Context, forgotPassRequest request.ForgotPassRequest) {
	// Set the expiration time for the reset token (5 minutes from now).
	expirationTime := time.Now().Add(time.Minute * 5)

	// Find the admin details using the provided email.
	var adminDetails model.Admin
	err := db.FindById(&adminDetails, forgotPassRequest.Email, "email")
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// If the email is not found, return a not found response.
		response.ShowResponse(utils.NOT_FOUND, utils.HTTP_NOT_FOUND, utils.FAILURE, nil, ctx)
		return
	} else if err != nil {
		// If there is an error in fetching admin details, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Create the reset token payload and generate the token from it.
	resetClaims := model.Claims{
		Id: adminDetails.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	tokenString, err := token.GenerateToken(resetClaims)
	if err != nil {
		// If there is an error in generating the reset token, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Create a reset session for resetting the password.
	resetSession := model.ResetSession{
		Id:    adminDetails.Id,
		Token: *tokenString,
	}
	err = db.CreateRecord(&resetSession)
	if err != nil {
		// If there is an error in creating the reset session, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, ctx, nil)
		return
	}

	// Create the reset password link to be sent to the admin's email address.
	link := "http://localhost:3014/reset-password?token=" + *tokenString

	// Sending the reset password link in the response.
	response.ShowResponse("Reset password link generated successfully", utils.HTTP_OK, utils.SUCCESS, link, ctx)
}

// ResetPasswordService handles reset password requests.
func ResetPasswordService(ctx *gin.Context, tokenString string, password request.UpdatePasswordRequest) {
	// Create a variable to hold the reset session details.
	var resetSession model.ResetSession

	// Decode the token and extract the required data.
	claims, err := token.DecodeToken(tokenString)
	if err != nil {
		// If there is an error in decoding the token, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, err.Error(), ctx)
		return
	}

	// Find the reset session using the extracted admin ID from the token.
	err = db.FindById(&resetSession, claims.Id, "id")
	if err != nil {
		// If there is an error in fetching the reset session, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, err.Error(), ctx)
		return
	}

	// Check if the token in the reset session matches the provided token.
	if resetSession.Token != tokenString {
		// If the tokens do not match, return a forbidden request response.
		response.ShowResponse("Forbidden request", utils.HTTP_FORBIDDEN, utils.FAILURE, err.Error(), ctx)
		return
	}

	// Call the UpdatePasswordService to update the admin's password.
	UpdatePasswordService(ctx, password, claims.Id)

	// After successfully resetting the password, delete the reset session.
	err = db.DeleteRecord(&resetSession, claims.Id, "id")
	if err != nil {
		// If there is an error in deleting the reset session, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, err.Error(), ctx)
		return
	}
}

// UpdatePasswordService handles updating the admin's password.
func UpdatePasswordService(ctx *gin.Context, password request.UpdatePasswordRequest, adminID string) {
	// Fetch the admin details using the provided adminID.
	var adminDetails model.Admin
	err := db.FindById(&adminDetails, adminID, "id")
	if err != nil {
		// If there is an error in fetching admin details, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	// Check the validity of the new password.
	err = utils.IsPassValid(password.Password)
	if err != nil {
		// If the password is invalid, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	//Hashing the new password
	// hashedPass, err := utils.HashPassword(password.Password)
	// if err != nil {
	// 	response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
	// 	return
	// }
	// Update the admin's password with the new one.
	adminDetails.Password = password.Password

	// Update the admin's password in the database.
	err = db.UpdateRecord(&adminDetails, adminID, "id").Error
	if err != nil {
		// If there is an error in updating the password, return an error response.
		response.ShowResponse(err.Error(), utils.HTTP_INTERNAL_SERVER_ERROR, utils.FAILURE, nil, ctx)
		return
	}

	// Respond with a success message.
	response.ShowResponse("Password updated successfully", utils.HTTP_OK, utils.SUCCESS, nil, ctx)
}
