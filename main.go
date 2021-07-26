package main

import (
	"fmt"
	"retailer-api/Routes"
	"retailer-api/Config"
	"retailer-api/Models"
	"github.com/jinzhu/gorm"
)

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
	Config.DB.AutoMigrate(&Models.Order{},&Models.Product{})

	//setup the routing confg
	r:=Routes.SetupRouter()

	//Create a shared container for concurrency handling
	Models.container := Models.ContainerStruct{}

	//Run the router
	r.Run()
}
