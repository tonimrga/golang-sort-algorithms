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
	fmt.Println(array)
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
