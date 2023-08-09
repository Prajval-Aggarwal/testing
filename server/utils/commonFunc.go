package utils

import (
	"bytes"
	"errors"
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/response"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/go-mail/mail"
	"golang.org/x/crypto/bcrypt"
)

func IsPassValid(password string) error {

	if len(password) < 8 {
		return errors.New("password is too short")

	}
	hasUpperCase := false
	hasLowerCase := false
	hasNumbers := false
	hasSpecial := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpperCase = true
		} else if char >= 'a' && char <= 'z' {
			hasLowerCase = true
		} else if char >= '0' && char <= '9' {
			hasNumbers = true
		} else if char >= '!' && char <= '/' {
			hasSpecial = true
		} else if char >= ':' && char <= '@' {
			hasSpecial = true
		}
	}

	if !hasUpperCase {
		return errors.New("password do not contain upperCase Character")
	}

	if !hasLowerCase {
		return errors.New("password do not contain LowerCase Character")
	}

	if !hasNumbers {
		return errors.New("password do not contain any numbers")
	}

	if !hasSpecial {
		return errors.New("password do not contain any special character")
	}
	return nil
}

func HashPassword(password string) (*string, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(bs)
	return &hashedPassword, nil
}

func AlreadyAtMax(val int) bool {
	return val == 5

}

// func UpgradeCarLevel(playerCarStats *model.PlayerCarsStats, ctx *gin.Context) {

// 	//list of points where car level upgrades automatically
// 	var car model.OwnedCars
// 	query := "select * from owned_cars where player_id=? and car_id=?;"

// 	db.QueryExecutor(query, &car, playerCarStats.PlayerId, playerCarStats.CarId)

// 	upgradePoints := [...]float64{20, 40, 60, 80, 100} //car will get upgraded to higher level whenever Or reaches these points

// 	for i := 0; i < len(upgradePoints); i++ {

// 		if playerCarStats.OR == upgradePoints[i] {

// 			// upgrade the car level
// 			car.Level++
// 			err := db.UpdateRecord(&car, car.Level, "level").Error
// 			if err != nil {
// 				fmt.Println("error in updation")
// 				// response.ShowResponse(STATS_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
// 				response.ShowResponse("server error", HTTP_INTERNAL_SERVER_ERROR, err.Error(), "", ctx)
// 			}
// 			fmt.Println("car level is Upgraded successfully!!")

// 			break
// 		}
// 	}
// }

func UpgradeCarLevel(playerCarStats *model.PlayerCarsStats) {

	//list of points where car level upgrades automatically
	var car model.OwnedCars
	query := "select * from owned_cars where player_id=? and car_id=?;"

	db.QueryExecutor(query, &car, playerCarStats.PlayerId, playerCarStats.CarId)

	if playerCarStats.OVR >= 20 && playerCarStats.OVR < 40 {

		if car.Level == 1 {
			//upgrade the level
			car.Level = 2
			//increase 30% repasir cost
			car.RepairCost += (0.3 * car.RepairCost)
			err := db.UpdateRecord(&car, playerCarStats.CarId, "car_id").Error
			if err != nil {
				fmt.Println("error in level upgrade")
			}
			fmt.Println("CAR LEVEL UPGRADED!!")

		}
	} else if playerCarStats.OVR >= 40 && playerCarStats.OVR < 60 {

		if car.Level == 2 {
			//upgrade the level
			car.Level = 3
			err := db.UpdateRecord(&car, playerCarStats.CarId, "car_id").Error

			if err != nil {
				fmt.Println("error in level upgrade")
			}
			fmt.Println("CAR LEVEL UPGRADED!!")

		}
	} else if playerCarStats.OVR >= 60 && playerCarStats.OVR < 80 {

		if car.Level == 3 {
			//upgrade the level
			car.Level = 4
			err := db.UpdateRecord(&car, playerCarStats.CarId, "car_id").Error

			if err != nil {
				fmt.Println("error in level upgrade")
			}
			fmt.Println("CAR LEVEL UPGRADED!!")

		}
	} else if playerCarStats.OVR >= 80 && playerCarStats.OVR <= 100 {

		if car.Level == 4 {
			//upgrade the level
			car.Level = 5
			err := db.UpdateRecord(&car, playerCarStats.CarId, "car_id").Error

			if err != nil {
				fmt.Println("error in level upgrade")
			}
			fmt.Println("CAR LEVEL UPGRADED!!")
		}
	}

}

func SetPlayerCarDefaults(playerId string, carId string) error {

	//set default car upgrades
	playerCarUpgrades := model.PlayerCarUpgrades{
		PlayerId:     playerId,
		CarId:        carId,
		Engine:       1,
		Turbo:        1,
		Intake:       1,
		Nitrous:      1,
		Body:         1,
		Tires:        1,
		Transmission: 1,
	}
	err := db.CreateRecord(playerCarUpgrades)
	if err != nil {
		//	response.ShowResponse(err.Error(), HTTP_INTERNAL_SERVER_ERROR, FAILURE, nil, ctx)
		return err
	}

	//set default car customizations
	var carDefualtCutomizations []model.DefaultCustomization
	err = db.FindAll(&carDefualtCutomizations, carId, "car_id")
	if err != nil {
		return err
	}
	for _, customise := range carDefualtCutomizations {
		playerCarCustomizations := model.PlayerCarCustomization{
			PlayerId:      playerId,
			CarId:         carId,
			Part:          customise.Part,
			ColorCategory: customise.ColorCategory,
			ColorType:     customise.ColorType,
			ColorCode:     customise.ColorCode,
			ColorName:     customise.ColorName,
			Value:         customise.Value,
		}
		err = db.CreateRecord(playerCarCustomizations)
		if err != nil {
			return err
		}
	}

	//set default car stats
	var currentCarStat model.CarStats
	err = db.FindById(&currentCarStat, carId, "car_id")
	if err != nil {

		return err
	}
	playerCarStats := model.PlayerCarsStats{
		PlayerId:    playerId,
		CarId:       carId,
		Power:       currentCarStat.Power,
		Grip:        currentCarStat.Grip,
		ShiftTime:   currentCarStat.ShiftTime,
		Weight:      currentCarStat.Weight,
		OVR:         currentCarStat.OVR,
		Durability:  currentCarStat.Durability,
		NitrousTime: float64(currentCarStat.NitrousTime),
	}
	err = db.CreateRecord(&playerCarStats)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCarDetails(tableName string, playerId string, carId string, ctx *gin.Context) error {
	query := "DELETE FROM " + tableName + " WHERE car_id =? AND player_id =?"
	err := db.RawExecutor(query, carId, playerId)
	if err != nil {
		response.ShowResponse(err.Error(), HTTP_BAD_REQUEST, FAILURE, nil, ctx)
		return err
	}
	return nil
}

func IsCarEquipped(playerId string, carId string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM owned_cars WHERE player_id =? AND car_id=? AND selected=true)"
	err := db.QueryExecutor(query, &exists, playerId, carId)
	if err != nil {
		return false
	}
	if !exists {
		return false
	}

	return true
}

func IsEmpty(field string) error {

	if field == "" {
		return errors.New(field + " cannot be empty")
	}
	return nil
}

func SendEmaillService(adminDetails model.Admin, link string) error {
	m := mail.NewMessage()
	m.SetHeader("From", "hoodRacing@gmail.com")
	m.SetHeader("Subject", "Reset Password!")

	var body bytes.Buffer
	tmp, err := template.ParseFiles("server/emailTemplate/forgot-password.html")
	if err != nil {
		fmt.Println("sajdsajdvsja", err.Error())
	}

	data := struct {
		Username string
		Link     string
	}{
		Username: adminDetails.Username,
		Link:     link,
	}

	tmp.Execute(&body, data)

	m.SetBody("text/html", body.String())
	if err != nil {
		fmt.Println("basdbbkjsad", err)
	}
	m.SetHeader("To", adminDetails.Email)
	d := mail.NewDialer(os.Getenv("SMTP_HOST"), 587, os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"))
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
