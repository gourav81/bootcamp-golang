//Models/Product.go
package Models

import (
	"E3/Config"
	_ "github.com/go-sql-driver/mysql"
	_ "log"
)

func GetAllProduct(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProduct(product *Product, id string, price int, quantity int) (err error) {

	if err = Config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	product.Price = price
	product.Quantity = quantity
	Config.DB.Save(product)
	return nil
}

func DeleteProduct(product *Product, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).Delete(product).Error; err != nil {
		return err
	}
	return nil
}
