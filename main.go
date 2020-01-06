package main

import "fmt"

func main() {
	consumption := 20 // 20 liter per 100km
	fuel := 15 //exist liter
	fmt.Println(distanceCalculation(consumption, fuel))




}
func distanceCalculation(consumption, fuel int)  int {

	distance := 100*fuel/consumption
	distance -= 10*distance/100 //not to get stuck on the road

	return distance
}