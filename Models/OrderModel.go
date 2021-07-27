package Models

import (
	"retailer-api/Config"
)

var ORD_ID int

func PlaceOrder(order *Order)(err error){
	ORD_ID++
	product_id:=order.product_id
	var product Product
	if err:= Config.DB.Where("id=?",product_id).First(&product).Error ; err!=nil{//product can't be retrieved
		newOrder:=Order{id: string(ORD_ID),product_id: product_id,quantity: order.quantity,status: "order failed"}
		Config.DB.Create(&newOrder)
		order.id= string(ORD_ID)
		order.status="order failed"
		return err
	}else{
		if product.quantity>=order.quantity {//Then place the order
			//decrease the quanity
			new_quantity:=product.quantity-order.quantity
			if err := Config.DB.Model(&Product{}).Where("id=?", product_id).Update("quantity", new_quantity).Error ; err!=nil{
				//updation in quantity failed, then 'failed' status will be inserted in Order table
				newOrder:=Order{id: string(ORD_ID),product_id: product_id,quantity: order.quantity,status: "order failed"}
				Config.DB.Create(&newOrder)
				order.id= string(ORD_ID)
				order.status="order failed"
				return err
			}else{
				//order placed will be inserted in the Order Table
				newOrder:=Order{id: string(ORD_ID),product_id: product_id,quantity: order.quantity,status: "order failed"}
				Config.DB.Create(&newOrder)
				order.id= string(ORD_ID)
				order.status="order placed"
				return nil
			}
		}else{//insufficient quantity is available
			newOrder:=Order{id: string(ORD_ID),product_id: product_id,quantity: order.quantity,status: "order failed"}
			Config.DB.Create(&newOrder)
			order.id= string(ORD_ID)
			order.status="order placed"
			return nil
		}
	}
}

func GetOrder(order *Order,ord_id int)(err error){
	if err:=Config.DB.Where("id=?",ord_id).First(order).Error; err!=nil{
		return err
	}else{
		return nil
	}
}
