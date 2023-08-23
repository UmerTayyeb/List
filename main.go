package main

import "fmt"

func main() {
	var num []int

	for len(num) < 5 {
		var input int
		fmt.Println("Write a number:")
		fmt.Scanln(&input)

		found := false
		for j := 0; j < len(num); j++ {
			if input == num[j] {
				fmt.Printf("This number %d is already in the list.\n", input)
				found = true
				break
			}
		}

		if !found {
			num = append(num, input)
		}
	}

	for index, value := range num {
		fmt.Printf("Index: %d \t Value: %d\n", index, value)
	}
}
