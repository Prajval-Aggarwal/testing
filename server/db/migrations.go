package db

import (
	"fmt"
	"main/server/model"

	"gorm.io/gorm"
)

func AutoMigrateDatabase(db *gorm.DB) {

	var dbVersion model.DbVersion
	err := db.First(&dbVersion).Error
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("db version is:", dbVersion.Version)
	if dbVersion.Version < 1 {
		err := db.AutoMigrate(&model.Player{}, &model.OwnedCars{}, &model.Car{}, &model.CarStats{}, &model.Garage{}, &model.GarageCarList{}, &model.GarageUpgrades{}, &model.OwnedGarage{}, &model.Engine{}, &model.Turbo{}, &model.Intake{}, &model.Body{}, &model.Nitrous{}, &model.Tires{}, &model.Transmission{}, &model.PlayerCarUpgrades{}, &model.PlayerCarsStats{})
		if err != nil {
			panic(err)
		}
		db.Create(&model.DbVersion{
			Version: 1,
		})
		dbVersion.Version = 1
	}
	if dbVersion.Version < 2 {
		err := db.AutoMigrate(&model.PlayerCarUpgrades{}, &model.PlayerCarsStats{})
		if err != nil {
			panic(err)
		}
		db.Where("version=?", dbVersion.Version).Updates(&model.DbVersion{
			Version: 2,
		})
		dbVersion.Version = 2
	}

}
