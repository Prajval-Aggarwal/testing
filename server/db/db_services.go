package db

import (
	"fmt"
	"main/server/response"

	"gorm.io/gorm"
)

var db *gorm.DB

func Transfer(connection *gorm.DB) {
	db = connection
}

func CreateRecord(data interface{}) error {

	err := db.Create(data).Error
	if err != nil {
		// fmt.Println("gorm error is", gorm.ErrDuplicatedKey.Error())
		// fmt.Println("error is", err.Error())
		return err
	}
	return nil
}

func FindById(data interface{}, id interface{}, columName string) error {
	column := columName + "=?"
	err := db.Where(column, id).First(data).Error
	if err != nil {
		return err
	}
	return nil
}

func FindAll(data interface{}, key interface{}, columnName string) error {
	column := columnName + "=?"
	err := db.Where(column, key).Find(data).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecord(data interface{}, id interface{}, columName string) *gorm.DB {
	column := columName + "=?"
	result := db.Where(column, id).Updates(data)

	return result
}

func QueryExecutor(query string, data interface{}, args ...interface{}) error {

	err := db.Raw(query, args...).Scan(data).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteRecord(data interface{}, id interface{}, columName string) error {
	column := columName + "=?"
	result := db.Where(column, id).Delete(data)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
func RecordExist(tableName string, value string, columnName string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM " + tableName + " WHERE " + columnName + "='" + value + "')"
	db.Raw(query).Scan(&exists)
	return exists
}

func RawExecutor(query string, args ...interface{}) error {
	err := db.Exec(query, args...).Error
	if err != nil {
		return err
	}
	return nil
}

func ResponseQuery(query string, args ...interface{}) response.PlayerResposne {
	playerResposne := &response.PlayerResposne{}
	row := db.Raw(query, args...).Row()
	fmt.Println("card Details is", playerResposne)
	row.Scan(&playerResposne.PlayerId, &playerResposne.PlayerName, &playerResposne.Level, &playerResposne.XP, &playerResposne.Role, &playerResposne.Email, &playerResposne.Coins, &playerResposne.Cash, &playerResposne.CarsOwned, &playerResposne.GaragesOwned, &playerResposne.DistanceTraveled, &playerResposne.ShdWon, &playerResposne.ShdWinRatio, &playerResposne.TdWon, &playerResposne.TdWinRatio)

	return *playerResposne

}
