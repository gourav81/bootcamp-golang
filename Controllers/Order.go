package Controllers

import (
	"E3/Entities"
	"E3/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var t = time.Now().Second()

func GetOrders(c *gin.Context) {
	var orders []Models.Order
	err := Models.GetAllOrder(&orders)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}

func CreateOrder(c *gin.Context) {
	if time.Now().Second()-t < 300 {
		c.JSON(http.StatusOK, Entities.CreateProductResponseOnCoolDown{Message: " Couldn't complete the order, transaction in cooldown period of 5 min"})
		return
	}
	var order Models.Order
	c.BindJSON(&order)
	order.Status = "order placed"

	err := Models.CreateOrder(&order)
	t = time.Now().Second()
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		c.JSON(http.StatusOK, order)
	}
}

func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOrderById(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
