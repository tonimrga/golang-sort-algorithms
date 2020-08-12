package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var size int
	fmt.Println("Enter number of elements in array: ")
	fmt.Scanln(&size)

	var array []int = createArray(size)
	fmt.Println("-----------------git ------------")
	fmt.Println("------- Array created -------")
	fmt.Println("------ Sorting started ------")
	fmt.Println("-----------------------------")

	bubbleSort(array)
	insertionSort(array)
	selectionSort(array)

	fmt.Println("-----------------------------")
	fmt.Println("------- Sorting ended -------")
	fmt.Println("-----------------------------")

}

func createArray(size int) []int {
	var array = make([]int, size)
	for i := 0; i < size; i++ {
		array = append(array[:i], i+1)
	}
	var shuffledArray = shuffleArray(array)
	return shuffledArray
}

func shuffleArray(array []int) []int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
	return array
}

func bubbleSort(array []int) {
	fmt.Println("Bubble sort starting")
	var sortedArray []int = array
	var length int = len(sortedArray)
	var startTime = time.Now()
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if sortedArray[j] > sortedArray[j+1] {
				sortedArray[j], sortedArray[j+1] = sortedArray[j+1], sortedArray[j]
			}
		}
	}
	var elapsed = time.Since(startTime).Nanoseconds()
	fmt.Printf("Bubble sort: %.3f ns\n", float64(elapsed))
}

func insertionSort(array []int) {
	fmt.Println("Insertion sort starting")
	var sortedArray []int = array
	var length int = len(sortedArray)
	var startTime = time.Now()
	for i := 1; i < length; i++ {
		var key int = sortedArray[i]
		var j int = i - 1
		for j >= 0 && sortedArray[j] > key {
			sortedArray[j+1] = sortedArray[j]
			j--
		}
		sortedArray[j+1] = key
	}
	var elapsed = time.Since(startTime).Nanoseconds()
	fmt.Printf("Insertion sort: %.3f ns\n", float64(elapsed))
}

func selectionSort(array []int) {
	fmt.Println("Selection sort starting")
	var sortedArray []int = array
	var length int = len(sortedArray)
	var startTime = time.Now()
	for i := 0; i < length-1; i++ {
		var min int = i
		for j := i + 1; j < length; j++ {
			if sortedArray[j] < sortedArray[min] {
				min = j
			}
		}
		sortedArray[i], sortedArray[min] = sortedArray[min], sortedArray[i]
	}
	var elapsed = time.Since(startTime).Nanoseconds()
	fmt.Printf("Selection sort: %.3f ns\n", float64(elapsed))
}
