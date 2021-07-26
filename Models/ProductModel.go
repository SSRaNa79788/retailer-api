package Models

import "retailer-api/Config"

func AddProduct(order *Order)(err error){
	if err:= Config.DB.Create(&order).Error ; err!=nil{
		return err
	}else{
		return nil
	}
}

func UpdateProduct(order *Order)(err error){
	if err := Config.DB.Save(&order).Error ; err!=nil{
		return err
	}else{
		return nil
	}
}

func GetProduct(order *Order, product_id int)(err error){
	if err := Config.DB.Where("id=?",product_id).First(order).Error ; err!=nil{
		return err
	}else{
		return nil
	}
}
func GetProducts(order []*Order)(err error){
	if err := Config.DB.Find(order).Error ; err!=nil{
		return err
	}else{
		return nil
	}
}

