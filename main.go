package main

import (
	"fmt"
	"reflect"
	"sync"
)

func calculateSum(numbers []int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	sums := 0
	for _, num := range numbers {
		sums += num
	}

	resultChan <- sums
}

func validateType(value interface{}, targetType string) bool {
	valueType := reflect.TypeOf(value).Kind().String()
	return valueType == targetType
}

func main() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{10, 8, 6, 9, 7}

	resultChan := make(chan int, 2)

	var wg sync.WaitGroup

	// Validasi numbers1
	if !validateType(numbers1, "slice" ) {
		fmt.Println("Error: numbers1 should be of type 'slice'")
		return
	}

	// Validasi numbers2
	if !validateType(numbers2, "slice") {
		fmt.Println("Error: numbers2 should be of type 'slice'")
		return
	}

	wg.Add(2)
	go calculateSum(numbers1, resultChan, &wg)
	go calculateSum(numbers2, resultChan, &wg)

	wg.Wait()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	total1 := <-resultChan
	total2 := <-resultChan

	fmt.Println("Jumlah deret bilangan pertama:", total1)
	fmt.Println("Jumlah deret bilangan kedua:", total2)
	fmt.Println("Total:", total1+total2)
}
