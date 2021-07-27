package Models

import (
	"sync"
	"time"
)

type ContainerStruct struct{
	mapp map[string]int
	mu sync.Mutex
}
var Container = ContainerStruct{}

func CanAcquireLock(product_id string)bool{
	Container.mu.Lock()
	if _,ok := Container.mapp[product_id] ; ok==true{
		Container.mu.Unlock()
		return false
	}else{
		Container.mapp[product_id]=1
		Container.mu.Unlock()
		return true
	}
}

func RemoveProductFromContainer(product_id string){
	Container.mu.Lock()
	delete(Container.mapp,product_id)
	Container.mu.Unlock()
}

func TryToAcquireLock(product_id string) bool{
	for i:=0;i<100;i++{
		time.Sleep(time.Millisecond)
		if CanAcquireLock(product_id)==false {
			continue
		} else {
			return true
		}
	}
	return false
}



