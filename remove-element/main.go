package main

import "fmt"

func removeLeftRight(arr []interface{}, position string) []interface{} {
	if position == "left" {
		return arr[1:]
	} else if position == "right" {
		return arr[:len(arr)-1]
	} else {
		return nil
	}
}

func main() {
	data1 := []interface{}{1, 2, 3, 4, 5}
	data2 := []interface{}{"Edo", "Budi", "Joko", "Tono"}

	newData1 := removeLeftRight(data1, "left")
	newData2 := removeLeftRight(data1, "right")
	newData3 := removeLeftRight(data2, "left")
	newData4 := removeLeftRight(data2, "right")
	newData5 := removeLeftRight(data2, "test")

	fmt.Println(newData1) // Output: [2 3 4 5]
	fmt.Println(newData2) // Output: [1 2 3 4]
	fmt.Println(newData3) // Output: [Budi Joko Tono]
	fmt.Println(newData4) // Output: [Edo Budi Joko]
	fmt.Println(newData5) // Output: []
}
