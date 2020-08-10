package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var numOfElements int
	fmt.Println("Enter number of elements in array: ")
	fmt.Scanln(&numOfElements)

	var array []int = createArray(numOfElements)
	fmt.Println("-----------------------------")
	fmt.Println("------- Array created -------")
	fmt.Println("------ Sorting started ------")
	selectionSort(array)
}

func createArray(numOfElements int) []int {
	var array = make([]int, numOfElements)
	for i := 0; i < numOfElements; i++ {
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

func selectionSort(array []int) {
	var sortedArray []int = array
	var length = len(sortedArray)
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
	var elapsed = time.Since(startTime)
	fmt.Println("-----------------------------")
	fmt.Println("Selection sort")
	fmt.Println("Time elapsed:", elapsed.Nanoseconds(), "ns")
	fmt.Println("-----------------------------")
}
