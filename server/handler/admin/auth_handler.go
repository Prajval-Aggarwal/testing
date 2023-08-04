package handler

import (
	"fmt"
	"main/server/request"
	"main/server/response"
	"main/server/services/auth"
	"main/server/utils"

	"github.com/gin-gonic/gin"
)

// GuestLoginService handles guest login and token generation.
//
// @Summary Guest Login
// @Description Perform guest login and generate access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param guestLoginRequest body request.GuestLoginRequest true "Guest Login Request"
// @Success 200 {object} response.Success "Login successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /guest-login [post]
func GuestLoginHandler(ctx *gin.Context) {
	var guestLoginReq request.GuestLoginRequest
	err := utils.RequestDecoding(ctx, &guestLoginReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	err = guestLoginReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	auth.GuestLoginService(ctx, guestLoginReq)
}

// just for testing not required...
func AdminSignUpHandler(ctx *gin.Context) {
	auth.AdminSignUpService(ctx)
}

// @Description	Forgot password
// @Accept			json
// @Produce		json
// @Success		200			{object}	response.Success
// @Failure		400			{object}	response.Success
// @Param			adminEmail	body		request.ForgotPassRequest	true	"Admin registered email"
// @Tags			Authentication
// @Router			/forgot-password [post]
func ForgotPasswordHandler(ctx *gin.Context) {
	var forgotRequest request.ForgotPassRequest
	err := utils.RequestDecoding(ctx, &forgotRequest)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = forgotRequest.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)

		return
	}

	auth.ForgotPassService(ctx, forgotRequest)
}

// @Description	Reset password
// @Accept			json
// @Produce		json
// @Success		200			{object}	response.Success
// @Failure		400			{object}	response.Success
// @Param			NewPassword	body		request.UpdatePasswordRequest	true	"Admins new password"
// @Tags			Authentication
// @Router			/reset-password [post]
func ResetPasswordHandler(ctx *gin.Context) {
	tokenString := ctx.Request.URL.Query().Get("token")
	var password request.UpdatePasswordRequest

	err := utils.RequestDecoding(ctx, &password)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = password.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	auth.ResetPasswordService(ctx, tokenString, password)

}

// LoginService handles user login and token generation.
//
// @Summary User Login
// @Description Perform user login and generate access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginDetails body request.LoginRequest true "Login Details"
// @Success 200 {object} response.Success "Login successful"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /login [post]
func LoginHandler(ctx *gin.Context) {
	var loginReq request.LoginRequest
	err := utils.RequestDecoding(ctx, &loginReq)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	err = loginReq.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}

	if loginReq.Password != "" {
		fmt.Println("admin login")
		auth.AdminLoginService(ctx, loginReq)
	} else {
		fmt.Println("player login ")
		auth.LoginService(ctx, loginReq)
	}

}

// UpdateEmailService updates the email of a player.
//
// @Summary Update Email
// @Description Update the email address of a player
// @Tags Player
// @Accept json
// @Produce json
// @Param email body request.UpdateEmailRequest true "Update Email Request"
// @Param Authorization header string true "Access Token"
// @Success 200 {object} response.Success "Email updated successfully"
// @Failure 400 {object} response.Success "Bad request"
// @Failure 401 {object} response.Success "Unauthorized"
// @Failure 404 {object} response.Success "Player not found"
// @Failure 500 {object} response.Success "Internal server error"
// @Router /update-email [put]
func UpdateEmailHandler(ctx *gin.Context) {
	var email request.UpdateEmailRequest
	playerId, exists := ctx.Get("playerId")
	fmt.Println("player id is:", playerId, exists)
	if !exists {
		response.ShowResponse(utils.UNAUTHORIZED, utils.HTTP_UNAUTHORIZED, utils.FAILURE, nil, ctx)
		return
	}
	err := utils.RequestDecoding(ctx, &email)
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	err = email.Validate()
	if err != nil {
		response.ShowResponse(err.Error(), utils.HTTP_BAD_REQUEST, "Failure", nil, ctx)
		return
	}
	auth.UpdateEmailService(ctx, email, playerId.(string))
}
