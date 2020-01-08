package main

func main() {

}

func distanceCalculation(consumption, fuel int) int {
	const inaccuracy = 10
	const km = 100

	distance := km * fuel / consumption
	distance -= inaccuracy * distance / km
	return distance
}
