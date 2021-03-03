package main

import "fmt"

func SmallestIntegerFinder(numbers []int) int {
	var smallest_int int = numbers[0]

	for _, value := range numbers {
		if smallest_int > value {
			smallest_int = value
		}
	}

	return smallest_int
}

func main() {
	numbers := []int{10, 2, 1, 4, 5, -6}
	var smallest int = SmallestIntegerFinder(numbers)
	fmt.Println(fmt.Sprint(smallest))
}
