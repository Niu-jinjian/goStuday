package main

import "fmt"

func main() {
	array := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(array))

	slice01 := []int{}
	slice01 = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice01)

	slice02 := make([]int, 10)
	fmt.Println(slice02)

	slice03 := []int{}
	fmt.Println(slice03)

	slice04 := make([]int, 10, 20)
	fmt.Println(slice04)

	slice05 := []int{}
	slice05 = append(slice05, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(slice05)

	arrayName := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	sliceName6 := arrayName[2:6]
	sliceName7 := arrayName[0:5:6]

	fmt.Println(sliceName6)
	fmt.Println(sliceName7)
	fmt.Println(len(sliceName7))
	fmt.Println(cap(sliceName7))
}
