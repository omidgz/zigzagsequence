package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findZigZagSequence(arr []int32) []int32 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	for i, j := len(arr)/2, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func main() {
	var str string
	fmt.Scanln(&str)
	t, _ := strconv.Atoi(str)
	for i := 0; i < t; i++ {
		fmt.Scanln(&str)
		n, _ := strconv.Atoi(str)
		var arr []int32
		for e := 0; e < n; e++ {
			fmt.Scan(&str)
			x, _ := strconv.Atoi(str)
			arr = append(arr, int32(x))
		}
		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(findZigZagSequence(arr)), "", "", -1), "[]"))
	}
}
