package Models

import "sync"

type ContainerStruct struct{
	mapp map[string]int
	mu sync.Mutex
}
var container ContainerStruct

func CanAcquireLock(product_id string)bool{
	container.mu.Lock()
	if _,ok := container.mapp[product_id] ; ok==true{
		container.mu.Unlock()
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
	delete(container.mapp,product_id)
	container.mu.Unlock()
}


