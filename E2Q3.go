package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var wg sync.WaitGroup

func withdraw(balance *int, amount int) {

	mutex.Lock()
	if *balance < amount {
		fmt.Println("Invalid Transaction")
	} else {
		*balance -= amount
	}
	wg.Done()
	defer mutex.Unlock()
}

func deposit(balance *int, amount int) {

	mutex.Lock()
	*balance += amount
	wg.Done()
	defer mutex.Unlock()
}

func main() {

	wg.Add(3)
	balance := 1000
	go deposit(&balance, 20)
	go withdraw(&balance, 40)
	go withdraw(&balance, 100)
	wg.Wait()
	fmt.Println("Balance")
	fmt.Println(balance)
}
