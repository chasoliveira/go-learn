package main

import "fmt"

func main() {
	//ages := []int{6, 5, 83, 5, 3, 18, 20, 83, 20}
	//ages := []int{43, 25, 83, 11, 31, 91, 85, 25, 95, 65}
	//ages := []int{61, 11, 33, 79, 81, 27, 79, 83, 9, 95}
	//ages := []int{69, 67, 59, 45, 59, 37, 65, 39, 85, 21}
	ages := []int{19, 5, 43, 13, 75, 89, 43, 89, 25, 49}

	fmt.Println(TwoOldestAges(ages))
}
func TwoOldestAges(ages []int) [2]int {
	firstOldest := max(ages, 0)
	secondOldest := max(ages, firstOldest)

	return [2]int{secondOldest, firstOldest}
}

func max(numbers []int, except int) int {
	max := 0
	for _, n := range numbers {
		if n != except && n > max {
			max = n
		}
	}
	return max
}
