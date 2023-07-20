package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/server/db"
	"main/server/model"

	"github.com/gin-gonic/gin"
)

func ReadJSONFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func AddDummyDataHandler(ctx *gin.Context) {
	fmt.Println("jhj")
	//car  dummy data
	addtoDb("server/dummyData/car.json", &[]model.Car{})
	addtoDb("server/dummyData/carStats.json", &[]model.CarStats{})

	// //garage dummy data
	addtoDb("server/dummyData/garage.json", &[]model.Garage{})
	addtoDb("server/dummyData/garageUpgrades.json", &[]model.GarageUpgrades{})

	// //car upgrades dummy data
	// addtoDb("server/dummyData/engineUpgrades.json", &[]model.Engine{})
	// addtoDb("server/dummyData/turboUpgrades.json", &[]model.Turbo{})
	// addtoDb("server/dummyData/intakeUpgrades.json", &[]model.Intake{})
	// addtoDb("server/dummyData/bodyUpgrades.json", &[]model.Body{})
	// addtoDb("server/dummyData/tiresUpgrades.json", &[]model.Tires{})
	// addtoDb("server/dummyData/transmissionUpgrades.json", &[]model.Transmission{})
	// addtoDb("server/dummyData/nitrousUpgrades.json", &[]model.Nitrous{})

	//car customization
	addtoDb("server/dummyData/customization.json", &[]model.CarCustomization{})
	addtoDb("server/dummyData/defaultCustomization.json", &[]model.DefaultCustomization{})

	//arena
	addtoDb("server/dummyData/arena.json", &[]model.Arena{})
	addtoDb("server/dummyData/arenaLevel.json", &[]model.ArenaReq{})

	addtoDb("server/dummyData/upgrades.json", &[]model.Upgrades{})
	addtoDb("server/dummyData/looseRewards.json", &[]model.LostRewards{})
	addtoDb("server/dummyData/winRewards.json", &[]model.WinRewards{})
	addtoDb("server/dummyData/classMultiplier.json", &[]model.RatingMulti{})

}

func addtoDb(filePath string, modelType interface{}) {

	carData, err := ReadJSONFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(carData, &modelType)
	if err != nil {
		log.Fatal(err)
	}

	switch slice := modelType.(type) {
	case *[]model.Car:
		for _, item := range *slice {
			fmt.Println("Car data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.CarStats:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("car stats data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.Garage:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("garage data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.GarageUpgrades:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("garage upgrades data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.Upgrades:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("transmission upgrades data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.CarCustomization:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("customization data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.DefaultCustomization:
		for _, item := range *slice {
			fmt.Println("default customization data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.Arena:
		for _, item := range *slice {
			fmt.Println("arena data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.ArenaReq:
		for _, item := range *slice {
			fmt.Println("arena level data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.WinRewards:
		for _, item := range *slice {
			fmt.Println("win rewards data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.LostRewards:
		for _, item := range *slice {
			fmt.Println("lost rewards data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.RatingMulti:
		for _, item := range *slice {
			fmt.Println("rating data:", item)
			db.CreateRecord(&item)
		}
	default:
		log.Fatal("Invalid modelType provided")
	}

}
