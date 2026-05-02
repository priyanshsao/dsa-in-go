package main

import "fmt"

func main() {
	exArr := []int{2,5,6,9}
	target := 9

	res := combinationSum(exArr, target)

	fmt.Print(res)
}