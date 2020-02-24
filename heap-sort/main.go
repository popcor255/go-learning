package main

import (
	"fmt"
	"math/rand"
	"time"
)

type minheap struct {
    heapArray []int
    size      int
    maxsize   int
}

func newMinHeap(maxsize int) *minheap {
    minheap := &minheap{
        heapArray: []int{},
        size:      0,
        maxsize:   maxsize,
    }	
	return minheap
}

func main(){
	heap := make_random_int_array();
	fmt.Println(arr);
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