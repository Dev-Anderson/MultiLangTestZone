package main

import "fmt"

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	//a lasanha leva 40 minutos
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayer int) int {
	//cada camada leva 2 minutos
	var timeLayers = 2
	return numberOfLayer * timeLayers
}

func Elapsedtime(numberOfLayers, actualMinutesInOven int) int {
	timePrepration := PreparationTime(numberOfLayers)
	return timePrepration + actualMinutesInOven
}

func main() {
	fmt.Println(RemainingOvenTime(35))
	fmt.Println(PreparationTime(4))
	fmt.Println(Elapsedtime(3, 20))
}
