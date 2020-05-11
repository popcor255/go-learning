package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Go Routines")

	go func() {
		lowerCase()
		wg.Done()
	}()

	go func() {
		upperCase()
		wg.Done()
	}()

	fmt.Println("Waiting to Finish")

	//wg.Wait()
	runtime.Gosched()
	fmt.Println("Terminating Program")
}

func lowerCase() {
	for count := 0; count < 3; count++ {
		for r := 'a'; r < 'z'; r++ {
			fmt.Print(string(r))
		}
	}
}

func upperCase() {
	for count := 0; count < 3; count++ {
		for r := 'A'; r < 'Z'; r++ {
			fmt.Print(string(r))
		}
	}
}
