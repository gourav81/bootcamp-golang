package Models

import (
	"E3/Config"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
)

func CreateOrder(order *Order) (err error) {

	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOrder(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderById(order *Order, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}
