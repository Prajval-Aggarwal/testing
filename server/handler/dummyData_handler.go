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
	// fmt.Println("jhj")
	// // car  dummy data
	// addtoDb("server/dummyData/car.json", &[]model.Car{})
	// addtoDb("server/dummyData/carStats.json", &[]model.CarStats{})

	// // //garage dummy data
	// addtoDb("server/dummyData/garage.json", &[]model.Garage{})
	// addtoDb("server/dummyData/garageUpgrades.json", &[]model.GarageUpgrades{})

	// // //car upgrades dummy data
	// addtoDb("server/dummyData/engineUpgrades.json", &[]model.Engine{})
	// addtoDb("server/dummyData/turboUpgrades.json", &[]model.Turbo{})
	// addtoDb("server/dummyData/intakeUpgrades.json", &[]model.Intake{})
	// addtoDb("server/dummyData/bodyUpgrades.json", &[]model.Body{})
	// addtoDb("server/dummyData/tiresUpgrades.json", &[]model.Tires{})
	// addtoDb("server/dummyData/transmissionUpgrades.json", &[]model.Transmission{})
	// addtoDb("server/dummyData/nitrousUpgrades.json", &[]model.Nitrous{})

	// //car custoization
	// addtoDb("server/dummyData/customization.json", &[]model.CarCustomization{})
	// addtoDb("server/dummyData/defaultCustomization.json", &[]model.DefaultCustomization{})

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
	case *[]model.Engine:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("engine upgrades data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.Turbo:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("turbo upgrades data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.Intake:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("intake upgrades data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.Nitrous:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("nitrous upgrades data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.Body:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("body upgrades data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.Tires:
		// Handle other struct types similarly
		for _, item := range *slice {
			fmt.Println("tires upgrades data:", item)
			db.CreateRecord(&item)
		}
	case *[]model.Transmission:
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

	default:
		log.Fatal("Invalid modelType provided")
	}

}
