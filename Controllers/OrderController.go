package Controllers

import (
	"net/http"
	"retailer-api/Models"
)


//Order methods
func PlaceOrder(c * gin.Context){
	var order Models.Order
	c.BindJSON(&order)
	if err := Models.PlaceOrder(&order) ; err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,order)
	}

}

func GetOrder(c * gin.Context){
	ord_id:=c.Params.ByName("id")
	var order Models.Order
	if err:=Models.GetOrder(&order,ord_id); err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,order)
	}
}

