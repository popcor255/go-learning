package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	unordered_arr := make_random_int_array()
	unordered_arr_string := fmt.Sprint(unordered_arr)
	fmt.Println("original array => " + unordered_arr_string)
	ordered_arr := bubble_sort(unordered_arr)
	ordered_arr_string := fmt.Sprint(ordered_arr)
	fmt.Println("sorted array => " + ordered_arr_string)
}

func bubble_sort(arr []int) (result []int) {
	result = arr
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1; j++ {
			if result[j+1] < result[j] {
				var temp = result[j+1]
				result[j+1] = result[j]
				result[j] = temp
			}
		}
	}

	return
}

func make_random_int_array() (result []int) {
	arr_size := get_rng_int(1, 100)
	for i := 0; i < arr_size; i++ {
		num := get_rng_int(1, 100)
		result = append(result, num)
	}
	return
}

func get_rng_int(min, max int) (result int) {
	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(max-min) + min
	return
}
