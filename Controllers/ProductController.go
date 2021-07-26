package Controllers

import (
	"net/http"
	"retailer-api/Models"
)

//Product Methods
func AddProduct(c * gin.Context){
	var product Models.Product
	c.BindJSON(&product)
	if err:=Models.AddProduct(&order) ; err!=nil{
		c.JSON(http.StatusOK,order)
	}else{
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func UpdateProduct(c * gin.Context){
	product_id:=c.Params.ByName()
	var product Models.Product
	c.BindJSON(&product)
	if err:= Models.GetProduct(&product,product_id) ; err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{//perform updation
		if err:=Models.UpdateProduct(&product,product_id) ; err!=nil{
			c.AbortWithStatus(http.StatusNotFound)
		}else{
			c.JSON(http.StatusOK,product)
		}
	}
}

func GetProduct(c * gin.Context){
	product_id := c.Params.ByName()
	var product Models.Product
	if err:=Models.GetProducts(&product,product_id) ; err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,product)
	}
}

func GetProducts(c * gin.Context){
	var products []Models.Product
	if err:=Models.GetProducts(&products) ; err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,products)
	}
}