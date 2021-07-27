package Customer

import (
	"retailer-api/Config"
	"time"
)

func GetCustomer(customer *Customer,CustID string)(err error){
	if err:= Config.DB.Where("id=?",CustID).First(customer).Error ; err!=nil{
		return err
	}else{
		return nil
	}
}

func AddCustomer(customer *Customer){
	Config.DB.Create(customer)
}

//function to check the customer corresponding to CustID is again trying to place the order within the coolong period or not
func ElligibleCustomer(CustID string)bool{
	var customer Customer
	if err := GetCustomer(&customer,CustID); err!=nil{
		return false
	}else if time.Now().Sub(customer.LastOrderTime) <= 5*time.Minute{
		return false
	}else{
		return true
	}
}
