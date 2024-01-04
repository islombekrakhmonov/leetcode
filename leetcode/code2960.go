package main

import "fmt"

func main() {
	fmt.Println(countTestedDevices([]int{1,1,2,1,3}))
}

func countTestedDevices(batteryPercentages []int) int {
	var count int
	for i:=0;i<len(batteryPercentages); i++{
		if batteryPercentages[i] > 0 {
			count++
			if i != len(batteryPercentages) -1 {
				for j:=i+1; j<len(batteryPercentages); j++{
					if batteryPercentages[j] > 0 {
						batteryPercentages[j] -=1 
					}
				}
			}
		}  
	}
    
	return count
}