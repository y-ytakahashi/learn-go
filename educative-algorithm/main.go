package main

import "fmt"

func SumArray(data []int) int {
	// 配列の長さを取得する
	size := len(data)
	// 合計値を格納する変数を宣言する
	total := 0
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total
}

func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func BinarySearch(data []int, value int) bool {
	var mid int
	low := 0
	high := len(data) - 1

	for low <= high {

		// 配列をここで半分に分割する
		mid = (low + high) / 2
		fmt.Println(mid)

		// 分割した境界値から判定を行う
		if data[mid] == value {

		}

	}
	return false
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(" Sum of all the values in array", SumArray(data))

	fmt.Println("SequentialSearch:", SequentialSearch(data, 5))
	fmt.Println("SequentialSearch:", SequentialSearch(data, 12))
	fmt.Println("BinarySearch:", BinarySearch(data, 1))
}
