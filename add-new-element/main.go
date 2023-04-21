package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		newDataSlice := []int{newData}
		data = append(newDataSlice, data...)
	} else if position == "down" {
		data = append(data, newData)
	} else {
		return nil
	}
	return data
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	newArr := AddElement(arr, 6, "up")
	fmt.Println(newArr) // [6 1 2 3 4 5]

	newArr = AddElement(arr, 6, "down")
	fmt.Println(newArr) // [1 2 3 4 5 6]
}
