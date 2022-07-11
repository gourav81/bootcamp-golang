package Routes

import (
	"E3/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("products", Controllers.GetProducts)
	r.POST("product", Controllers.CreateProduct)
	r.GET("product/:id", Controllers.GetProductByID)
	r.PATCH("product/:id", Controllers.UpdateProduct)
	r.DELETE("product/:id", Controllers.DeleteProduct)
	r.GET("orders", Controllers.GetOrders)
	r.POST("order", Controllers.CreateOrder)
	r.GET("order/:id", Controllers.GetOrderByID)

	return r
}
