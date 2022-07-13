package main

import (
	"E3/Config"
	"E3/Controllers"
	"E3/Models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateProduct(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	r := SetUpRouter()

	r.POST("/product", Controllers.CreateProduct)
	product := Models.Product{
		ProductName: "T",
		Price:       20,
		Quantity:    30,
	}
	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateOrder(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Order{})
	r := SetUpRouter()

	r.POST("/order", Controllers.CreateOrder)
	order := Models.Order{
		CustomerId: "customer123",
		ProductId:  "product123",
		Quantity:   30,
		Status:     "order placed",
	}
	jsonValue, _ := json.Marshal(order)
	req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProduct(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Product{})
	r := SetUpRouter()
	r.GET("/products", Controllers.GetProducts)
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var products []Models.Product
	json.Unmarshal(w.Body.Bytes(), &products)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, products)
}

func TestGetOrder(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Order{})
	r := SetUpRouter()
	r.GET("/orders", Controllers.GetOrders)
	req, _ := http.NewRequest("GET", "/orders", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var orders []Models.Order
	json.Unmarshal(w.Body.Bytes(), &orders)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, orders)
}

func TestGetProductByID(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Product{})
	r := SetUpRouter()
	r.GET("/product/:id", Controllers.GetProductByID)
	req, _ := http.NewRequest("GET", "/product/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetOrderByID(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Order{})
	r := SetUpRouter()
	r.GET("/order/:id", Controllers.GetOrderByID)
	req, _ := http.NewRequest("GET", "/order/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateProduct(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})

	r := SetUpRouter()
	r.PATCH("/product/:id", Controllers.UpdateProduct)
	product := Models.Product{
		Price:    30,
		Quantity: 50,
	}
	jsonValue, _ := json.Marshal(product)
	req, _ := http.NewRequest("PATCH", "/product/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteProduct(t *testing.T) {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Product{})
	r := SetUpRouter()
	r.DELETE("/product/:id", Controllers.DeleteProduct)
	req, _ := http.NewRequest("DELETE", "/product/3", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
