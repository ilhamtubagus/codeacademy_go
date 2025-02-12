package main

import "fmt"

func accessSlice(slice []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from panic:", r)
		} else {
			fmt.Println(r)
		}
	}()

	fmt.Println(slice[index])
}

func main() {
	nums := []int{1, 2, 3}

	accessSlice(nums, 5)

	fmt.Println("Accessing a valid index:", nums[1])
}
