package Product

import (
	"retailer-api/Config"
)

func AddProduct(product *Product)(err error){
	if err:= Config.DB.Create(&product).Error ; err!=nil{
		return err
	}else{
		return nil
	}
}

func UpdateProduct(product *Product,product_id string)(err error){
	if err := Config.DB.Where("id=?",product_id).Save(&product).Error ; err!=nil{
		return err
	}else{
		return nil
	}
}

func GetProduct(product *Product, product_id string)(err error){
	if err := Config.DB.Where("id=?",product_id).First(product).Error ; err!=nil{
		return err
	}else{
		return nil
	}
}
func GetProducts(product *[]Product)(err error){
	if err := Config.DB.Find(product).Error ; err!=nil{
		return err
	}else{
		return nil
	}
}

