package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_PRICE float32 = 5
var MAX_PANNER_PRICE float32 = 5

func main() {
	var foodChannel = make(chan string) // to find which website has min price of the food in parallel
	var pannerChannel = make(chan string)
	var websites = []string{"walmart.com", "cosmos.com", "wholefood.com"}
	for i := range websites {
		go checkFoodPrice(websites[i], foodChannel)
		go checkPannerPrice(websites[i], pannerChannel)
	}
	sendMessge(foodChannel, pannerChannel)
}

func checkPannerPrice(website string, paneerChannel chan string) {
	for {
		time.Sleep(time.Second + 1)
		var foodPrice = rand.Float32() * 20
		if foodPrice <= MAX_PRICE {
			paneerChannel <- website
			break
		}
	}
}

func checkFoodPrice(website string, foodChannel chan string) {
	for {
		time.Sleep(time.Second + 1)
		var foodPrice = rand.Float32() * 20
		if foodPrice <= MAX_PRICE {
			foodChannel <- website
			break
		}
	}
}

func sendMessge(foodChannel chan string, pannerChannel chan string) {
	select {
	case website := <-foodChannel:
		fmt.Printf("Text: Found deal on food in website : %v\n", website)
	case website := <-pannerChannel:
		fmt.Printf("Email: Found deal on food in website : %v\n", website)
	}
}
