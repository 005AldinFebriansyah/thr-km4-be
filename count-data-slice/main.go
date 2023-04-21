package main

func howManyElements(data []interface{}) int {
	return len(data)
}

func main() {
	data1 := []interface{}{1, 2, 3, 4, 5}
	data2 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data3 := []interface{}{1, 1, 1, 5, 5, 5}
	data4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	data5 := []interface{}{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	data6 := []interface{}{true, false, true, false, true, false}

	println(howManyElements(data1)) // Output: 5
	println(howManyElements(data2)) // Output: 10
	println(howManyElements(data3)) // Output: 6
	println(howManyElements(data4)) // Output: 4
	println(howManyElements(data5)) // Output: 8
	println(howManyElements(data6)) // Output: 6
}
