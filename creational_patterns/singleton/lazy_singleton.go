package singleton

// package main

import (
	"fmt"
	"sync"
)

var (
	lock = &sync.Mutex{}
	// once = &sync.Once{}
)

type lazySingleton struct {
}

var lazySingletonInstance *lazySingleton

func GetLazySingletonInstance() *lazySingleton {

	// if lazySingletonInstance == nil {
	// 	once.Do(func() {
	// 		fmt.Println("Creating lazy_single instance now.")
	// 		lazySingletonInstance = &lazySingleton{}
	// 	})
	// } else {
	// 	fmt.Println("Lazy_single instance already created.")
	// }
	// return lazySingletonInstance

	if lazySingletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if lazySingletonInstance == nil {
			fmt.Println("Creating lazy_single instance now.")
			lazySingletonInstance = &lazySingleton{}
		} else {
			fmt.Println("Lazy_single instance already created.")
		}
	} else {
		fmt.Println("Lazy_single instance already created.")
	}
	return lazySingletonInstance
}

// func main() {
// 	for i := 0; i < 30; i++ {
// 		go GetLazySingletonInstance()
// 	}
// 	fmt.Scanln()
// }
