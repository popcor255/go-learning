package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Heap struct {
	heapArray []int
	size      int
	maxsize   int
}

func newMinHeap(arr []int) *Heap {
	heap := &Heap{
		heapArray: arr,
		size:      0,
		maxsize:   len(arr),
	}
	return heap
}

func main() {
	heap := newMinHeap(make_random_int_array())
	heap.heap_sort()
	fmt.Println(heap)
}

func make_random_int_array() (result []int) {
	arr_size := random_int(1, 100)
	for i := 0; i < arr_size; i++ {
		num := random_int(1, 100)
		result = append(result, num)
	}
	return
}

func random_int(min, max int) (result int) {
	rand.Seed(time.Now().UnixNano())
	result = rand.Intn(max-min) + min
	return
}

func (heap *Heap) heap_sort() {
	array := heap.heapArray
	heap.build_heap(array)

	for length := len(array); length > 1; length-- {
		heap.remove_top(array, length)
	}
}

func (heap *Heap) build_heap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		heap.Heapify(array, i, len(array))
	}
}

func (heap *Heap) remove_top(array []int, length int) {
	var lastIndex = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	heap.Heapify(array, 0, lastIndex)
}

func (heap *Heap) Heapify(array []int, root, length int) {
	var max = root
	var l, r = heap.Left(array, root), heap.Right(array, root)

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		heap.Heapify(array, max, length)
	}
}

func (*Heap) Left(array []int, root int) int {
	return (root * 2) + 1
}

func (*Heap) Right(array []int, root int) int {
	return (root * 2) + 2
}
