package Order

import (
	"retailer-api/Config"
	"retailer-api/Models"
	"retailer-api/Models/Customer"
	"retailer-api/Models/Product"
)

func PlaceOrder(order *Order)(err error){

	product_id:=order.Product_id
	if Models.TryToAcquireLock(product_id)==false{
		HandleFailedOrder(order)
	}

	defer Models.RemoveProductFromContainer(product_id)

	//check if customer is place=ing order within the cooling-down  period
	if Customer.ElligibleCustomer(order.Cust_id)==false{
		HandleFailedOrder(order)
		return nil
	}

	//Now trying to place the order---------

	//1st retrieve the product details inorder to know the availability
	var product Product.Product
	if err= Config.DB.Where("id=?",product_id).First(&product).Error ; err!=nil{//product can't be retrieved
		HandleFailedOrder(order)
		return err
	}else{
		if product.Quantity >=order.Quantity { //Then place the order and try to decrease the quanity
			new_quantity:=product.Quantity -order.Quantity
			if err = Config.DB.Model(&Product.Product{}).Where("id=?", product_id).Update("quantity", new_quantity).Error ; err!=nil{
				//updation in quantity failed, then 'failed' status will be inserted in Order table
				HandleFailedOrder(order)
				return err
			}else{ //order placed will be inserted in the Order Table
				HandleSuccessfulOrder(order)
				return nil
			}
		}else{ //insufficient quantity is available
			HandleFailedOrder(order)
			return nil
		}
	}
}

func GetOrder(order *Order,ord_id string)(err error){
	if err:=Config.DB.Where("id=?",ord_id).First(order).Error; err!=nil{
		return err
	}else{
		return nil
	}
}
