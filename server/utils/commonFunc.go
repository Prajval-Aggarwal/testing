package utils

import (
	"errors"
	"fmt"
	"main/server/db"
	"main/server/model"
	"main/server/response"
	"math"
	"time"

	"github.com/gin-gonic/gin"
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

func SetPlayerCarDefaults(playerId string, carId string) error {

	//set default car upgrades
	playerCarUpgrades := model.PlayerCarUpgrades{
		PlayerId:     playerId,
		CarId:        carId,
		Engine:       0,
		Turbo:        0,
		Intake:       0,
		Nitrous:      1,
		Body:         0,
		Tires:        0,
		Transmission: 0,
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

func CalculateOVR(classOr, power, grip, weight float64) float64 {
	x := (classOr * (0.7*float64(power) + (0.6 * float64(grip)))) - 0.02*float64(weight)
	return x
}

func UpgradeData(playerId string, carId string) (*model.Player, *model.PlayerCarsStats, *model.PlayerCarUpgrades, string, uint64, *model.RatingMulti, error) {
	var playerDetails model.Player
	var playerCarStats model.PlayerCarsStats
	var carClassDetails string

	var playerCarUpgrades model.PlayerCarUpgrades
	var maxUpgradeLevel uint64
	var classRating model.RatingMulti
	//check if the car is owned or not
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM owned_cars WHERE player_id =? AND car_id=?)"
	err := db.QueryExecutor(query, &exists, playerId, carId)
	if err != nil {
		return nil, nil, nil, "", 0, nil, err
	}
	if !exists {
		return nil, nil, nil, "", 0, nil, errors.New(NOT_FOUND)
	}
	err = db.FindById(&playerDetails, playerId, "player_id")
	if err != nil {
		return nil, nil, nil, "", 0, nil, err
	}

	query = "SELECT * FROM player_cars_stats WHERE player_id=? AND car_id=?"
	err = db.QueryExecutor(query, &playerCarStats, playerId, carId)
	if err != nil {
		return nil, nil, nil, "", 0, nil, err
	}

	query = "SELECT class FROM cars WHERE car_id=?"
	err = db.QueryExecutor(query, &carClassDetails, carId)
	if err != nil {
		return nil, nil, nil, "", 0, nil, err
	}

	query = "SELECT * FROM rating_multis WHERE class=?"
	err = db.QueryExecutor(query, &classRating, carClassDetails)
	if err != nil {
		return nil, nil, nil, "", 0, nil, err
	}

	query = "SELECT * FROM player_car_upgrades WHERE player_id=? AND car_id=?"
	err = db.QueryExecutor(query, &playerCarUpgrades, playerId, carId)
	if err != nil {
		return nil, nil, nil, "", 0, nil, err
	}

	query = "SELECT upgrade_level FROM upgrades WHERE class =? ORDER BY upgrade_level DESC LIMIT 1;"
	err = db.QueryExecutor(query, &maxUpgradeLevel, carClassDetails)
	if err != nil {
		return nil, nil, nil, "", 0, nil, err
	}

	return &playerDetails, &playerCarStats, &playerCarUpgrades, carClassDetails, maxUpgradeLevel, &classRating, nil
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func IsExisting(query string, values ...interface{}) bool {
	var exists bool
	err := db.QueryExecutor(query, &exists, values...)
	if err != nil {
		return false
	}
	return exists
}

func TimeConversion(stringFormat string) *time.Time {
	timeFormat, err := time.Parse("15:04:05.9999999", stringFormat)
	if err != nil {
		fmt.Println("error in parsing the string format of time")
		return nil
	}
	return &timeFormat
}
