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
		err := db.AutoMigrate(&model.Player{}, &model.OwnedCars{}, &model.Car{}, &model.CarStats{}, &model.Garage{}, &model.GarageCarList{}, &model.GarageUpgrades{}, &model.OwnedGarage{}, &model.PlayerCarUpgrades{}, &model.PlayerCarsStats{}, &model.Upgrades{}, &model.LostRewards{}, &model.WinRewards{}, &model.RatingMulti{})
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
	if dbVersion.Version < 3 {
		err := db.AutoMigrate(&model.CarCustomization{}, &model.DefaultCustomization{}, &model.PlayerCarCustomization{})
		if err != nil {
			panic(err)
		}
		db.Where("version=?", dbVersion.Version).Updates(&model.DbVersion{
			Version: 3,
		})
		dbVersion.Version = 3
	}
	if dbVersion.Version < 4 {
		err := db.AutoMigrate(&model.PlayerCarCustomization{})
		if err != nil {
			panic(err)
		}
		db.Where("version=?", dbVersion.Version).Updates(&model.DbVersion{
			Version: 4,
		})
		dbVersion.Version = 4
	}
	if dbVersion.Version < 5 {
		err := db.AutoMigrate(&model.PlayerRaceHistory{}, &model.OwnedBattleArenas{})
		if err != nil {
			panic(err)
		}
		db.Where("version=?", dbVersion.Version).Updates(&model.DbVersion{
			Version: 5,
		})
		dbVersion.Version = 5
	}
	if dbVersion.Version < 6 {
		err := db.AutoMigrate(&model.ResetSession{})
		if err != nil {
			panic(err)
		}
		db.Where("version=?", dbVersion.Version).Updates(&model.DbVersion{
			Version: 6,
		})
		dbVersion.Version = 6
	}

}
