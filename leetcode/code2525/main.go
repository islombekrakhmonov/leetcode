package main

import "fmt"

func main() {
	fmt.Println(categorizeBox(1000, 35, 700, 300))
}

func categorizeBox(length int, width int, height int, mass int) string {
	var category []string
	volume := length * width * height

	if length >= int(1e4) || width >= int(1e4) || height >= int(1e4) || volume >= int(1e9) {
		category = append(category, "Bulky")
	}
	if mass >= 100 {
		category = append(category, "Heavy")
	}

	if len(category) > 1 {
		return "Both"
	} else if len(category) == 0 {
		return "Neither"
	}

	return category[0]
}

/*
The box is "Bulky" if:
Any of the dimensions of the box is greater or equal to 104.
Or, the volume of the box is greater or equal to 109.
If the mass of the box is greater or equal to 100, it is "Heavy".
If the box is both "Bulky" and "Heavy", then its category is "Both".
If the box is neither "Bulky" nor "Heavy", then its category is "Neither".
If the box is "Bulky" but not "Heavy", then its category is "Bulky".
If the box is "Heavy" but not "Bulky", then its category is "Heavy".
Note that the volume of the box is the product of its length, width and height.
*/
