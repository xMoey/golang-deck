package main

import "fmt"

func main() {
	randomNumbers := []int{33, 12, 41, 49, 88, 85, 9, 43, 52, 83}
	for _, r := range randomNumbers {
		if r%2 == 1 {
			fmt.Println("odd", r)
		} else {
			fmt.Println("even", r)
		}
	}
}
