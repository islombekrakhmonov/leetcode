package main

import "fmt"

func main() {
	fmt.Println(convertTemperature(36.5))
}

func convertTemperature(celsius float64) []float64 {
	var output []float64
    kelvin := celsius + 273.15
	fahrenheit := celsius * 1.80 + 32.00
	output = append(output, kelvin,fahrenheit)

	return output
}