package main

import "fmt"

func main() {
	// 删除切片中的元素
	sliceName1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 删除第一个元素
	sliceName1 = sliceName1[1:]
	fmt.Println(sliceName1)

	// 删除最后一个元素
	sliceName1 = sliceName1[:len(sliceName1)-1]
	fmt.Println(sliceName1)

	// 删除指定位置的元素
	index := 3
	sliceName1 = append(sliceName1[:index], sliceName1[index+1:]...)
	fmt.Println(sliceName1)

	// 删除指定元素块
	startindex := 2
	endindex := 5
	sliceName1 = append(sliceName1[:startindex], sliceName1[endindex+1:]...)
	fmt.Println(sliceName1)

	// 清空切片
	sliceName1 = nil
	fmt.Println(sliceName1)

}
