package Order

import (
	"retailer-api/Config"
	"retailer-api/Models/Customer"
	"sync"
	"time"
)

type OrderIDHandler struct{
	ORD_ID int
	mu sync.Mutex
}

var Orderidhandler OrderIDHandler


func HandleFailedOrder(order *Order){
	Orderidhandler.mu.Lock()
	product_id:=order.Product_id
	Orderidhandler.ORD_ID++
	ORD_ID := Orderidhandler.ORD_ID
	newOrder:= Order{Id: string(ORD_ID),Product_id: product_id,Quantity: order.Quantity,Status: "order failed"}
	Config.DB.Create(&newOrder)
	order.Id= string(ORD_ID)
	order.Status="order failed"
	Orderidhandler.mu.Lock()
}
func HandleSuccessfulOrder(order *Order){
	Orderidhandler.mu.Lock()
	product_id:=order.Product_id
	Orderidhandler.ORD_ID++
	ORD_ID := Orderidhandler.ORD_ID
	newOrder:= Order{Id: string(ORD_ID),Product_id: product_id,Quantity: order.Quantity,Status: "order placed"}

	//Create a new order
	Config.DB.Create(&newOrder)
	//Now update the Customer Table that This customer's order is placed successfully at time.Now()
	customer := Customer.Customer{Id:order.Cust_id,LastOrderTime: time.Now()}
	Customer.AddCustomer(&customer)

	order.Id= string(ORD_ID)
	order.Status="order placed"
	Orderidhandler.mu.Lock()
}
