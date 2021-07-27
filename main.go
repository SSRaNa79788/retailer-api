package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"retailer-api/Config"
	"retailer-api/Models/Order"
	"retailer-api/Models/Product"
	"retailer-api/Routes"
)
var err error

func main(){
	//create database conection
	Config.DB ,err = gorm.Open("mysql",Config.DbURL(Config.BuildDBConfig()))

	//check for error in building database connection
	if err!=nil{
		fmt.Println("status: ",err)
	}

	//at last, Close the DB connection
	defer Config.DB.Close()

	//Do automigration
	Config.DB.AutoMigrate(&Order.Order{},&Product.Product{})

	//setup the routing confg
	r:=Routes.SetupRouter()

	//Create a shared container for concurrency handling
	//Models.Container := Models.ContainerStruct{}

	Order.Orderidhandler = Order.OrderIDHandler{ORD_ID: 1}

	//Run the router
	r.Run()
}
