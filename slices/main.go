package main

import (
	"fmt"
)

func main() {
	fmt.Println("running")
	arr := []int{1, 2, 3, 4}
	var temp []int

	for i := 0; i < len(arr); i++ {
		arr[i] = 3
		//shares a backing array
		temp = arr[1:2]
		temp[0] = 2
	}

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		arr[i] = 5
		//shares backing array
		//chaning capacity gives a copy
		temp = arr[1:2:2]
		temp = append(temp, 2)
	}

	fmt.Println(arr)
	fmt.Println(temp)

	//only make copy and send copy of arr unless within scope
	//append is a coding smell
	for i := 0; i < len(arr); i++ {
		arr[i] = 5
		//shares backing array
		//chaning capacity gives a copy
		copy(temp, arr[1:2])
		temp = append(temp, 2)
	}

	fmt.Println(arr)
	fmt.Println(temp)


	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{}

	fmt.Println(cap(arr2))

	for right := len(arr1) - 1; right >= 0; right-- {
		arr2 = append(arr2, arr1[right])
		fmt.Println(cap(arr2))
	}

	fmt.Println(arr1)
	fmt.Println(arr2)

}
