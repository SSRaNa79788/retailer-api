package Models

import "sync"

type ContainerStruct struct{
	mapp map[string]int
	mu sync.Mutex
}
var container ContainerStruct

func CanAcquireLock(product_id string)bool{
	if _,ok := container.mapp[product_id] ; ok==true{
		return false
	}else{
		return true
	}
}

func LockContainer(product_id string){
	container.mu.Lock()
	container.mapp[product_id]=1
	container.mu.Unlock()
}

func UnlockContainer(product_id string){
	container.mu.Lock()
	delete(container.mapp,product_id)
	container.mu.Unlock()
}


