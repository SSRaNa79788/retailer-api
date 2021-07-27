package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"retailer-api/Models/Order"
)


//Order methods
func PlaceOrder(c * gin.Context){
	var order Order.Order
	c.BindJSON(&order)
	if err := Order.PlaceOrder(&order) ; err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,order)
	}

}

func GetOrder(c * gin.Context){
	ord_id:=c.Params.ByName("id")
	var order Order.Order
	if err:= Order.GetOrder(&order,ord_id); err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,order)
	}
}

