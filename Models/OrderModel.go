package Models


func PlaceOrder(order *Order,){
	return
}

func GetOrder(order *Order,ord_id int)(err error){
	if err:=Config.DB.Where("id=?",ord_id).First(order).Error; err!=nil{
		return err
	}else{
		return nil
	}
}
