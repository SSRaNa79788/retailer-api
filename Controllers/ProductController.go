package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"retailer-api/Models/Product"
)

//Product Methods
func AddProduct(c * gin.Context){
	var product Product.Product
	c.BindJSON(&product)
	if err:= Product.AddProduct(&product) ; err!=nil{
		c.JSON(http.StatusOK,product)
	}else{
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func UpdateProduct(c * gin.Context){
	product_id:=c.Params.ByName("id")
	var product Product.Product
	c.BindJSON(&product)
	if err:= Product.GetProduct(&product,product_id) ; err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{//perform updation
		if err:= Product.UpdateProduct(&product,product_id) ; err!=nil{
			c.AbortWithStatus(http.StatusNotFound)
		}else{
			c.JSON(http.StatusOK,product)
		}
	}
}

func GetProduct(c * gin.Context){
	product_id := c.Params.ByName("id")
	var product Product.Product
	if err:= Product.GetProduct(&product,product_id) ; err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,product)
	}
}

func GetProducts(c * gin.Context){
	var products []Product.Product
	if err:= Product.GetProducts(&products) ; err!=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK,products)
	}
}