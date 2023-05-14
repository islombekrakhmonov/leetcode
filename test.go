package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"unsafe"
)

func main() {
	first := []int{1,2,3}
	second := []int{2,5,6}
	
	fmt.Println(powerOfFour(64))
	fmt.Println(connectArrays(first,second))


	a := "hello"
	b:= a[:3]
	fmt.Println(&a != &b)

	a = "bro hello"

	fmt.Println(&a != &b)



	name := "Hello"
	t := name[0:3]
   
	fmt.Printf("Add ress of name: %08x\n", stringAddr(name))
	fmt.Printf("Add ress of t: %08x\n", stringAddr(t))
	fmt.Println(stringAddr(name) == stringAddr(t)) // Gues the output of this line
   
	name += "hello world 😀"
   
	fmt.Printf("Add ress of name: %08x\n", stringAddr(name))
	fmt.Printf("Add ress of t: %08x\n", stringAddr(t))
	fmt.Println(stringAddr(name) == stringAddr(t)) // Gues the output of this line
   


}

func powerOfFour(n int)bool{
	res := math.Log(float64(n)) / math.Log(float64(4))
	if res == math.Floor(res) {
	   fmt.Printf("%d is the power of 4.\n", n)
	   return true
	} else {
	   fmt.Printf("%d is not the power of 4.\n", n)
	   return false
	}
}


func connectArrays(first []int, second[]int)[]int{
	var output []int
	output = append(first,second...)
	// output = append(output, second...)
	pop := output
	sort.Ints(output)
	

	for i := 0; i <= len(pop)-1; i++ {
		for j := 0; j < len(pop)-1-i; j++ {
		   if pop[j] > pop[j+1] {
			  pop[j], pop[j+1] = pop[j+1], pop[j]
			  fmt.Println(pop[j], pop[j+1])
		   }
		}
	 }
	 fmt.Println("Hi")
	 fmt.Println(pop)



	return output
}

func stringAddr(s string) uintptr {
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
   }