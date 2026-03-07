package main

import "fmt"

func main() {
	fmt.Println(distanceTraveled(9, 1))
}

func distanceTraveled(mainTank int, additionalTank int) int {
	var distance int
	for mainTank > 0 {
		if mainTank >= 5 {
			distance += 50
			mainTank -= 5
			if additionalTank > 0 {
				additionalTank--
				mainTank += 1
			}
		} else {
			distance += 10
			mainTank--
		}
	}

	return distance
}
