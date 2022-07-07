package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func giveRating(rating int, ratingSum *int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	*ratingSum += rating
	//fmt.Print(*ratingSum)
}

func main() {

	var wg sync.WaitGroup
	var rating [50]int
	for i := 0; i < 50; i++ {
		rating[i] = rand.Intn(50)
	}
	var ratingSum int
	for i := 0; i < 50; i++ {
		wg.Add(1)
		rat := rating[i]
		go func() {
			defer wg.Done()
			giveRating(rat, &ratingSum)
		}()
	}
	wg.Wait()
	fmt.Print("Average Rating\n")
	fmt.Print(ratingSum / 50)
}
