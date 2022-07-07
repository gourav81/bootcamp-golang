package main

import (
	"fmt"
	"time"
)

func count(from string, frequency map[string]int) {
	for i := 0; i < len(from); i++ {
		frequency[string(from[i])] += 1
	}
}

func main() {

	strings := []string{"quick", "brown", "fox", "lazy", "dog"}
	frequency := map[string]int{}
	for i := 0; i < len(strings); i++ {
		go count(strings[i], frequency)
	}

	time.Sleep(time.Second)
	fmt.Println(frequency)
}
