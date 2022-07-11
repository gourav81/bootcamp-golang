package Controllers

import (
	"E3/Entities"
	"E3/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	_ "log"
	"net/http"
)

func GetProducts(c *gin.Context) {
	var products []Models.Product
	err := Models.GetAllProduct(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}

func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	log.Println(product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var response Entities.CreateProductResponse
		response.Id = product.Id
		response.ProductName = product.ProductName
		response.Price = product.Price
		response.Quantity = product.Quantity
		response.Message = "product successfully added"
		c.JSON(http.StatusOK, response)
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	id := c.Params.ByName("id")

	var requestedProduct Models.Product
	c.BindJSON(&requestedProduct)

	var product Models.Product
	err := Models.UpdateProduct(&product, id, requestedProduct.Price, requestedProduct.Quantity)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
