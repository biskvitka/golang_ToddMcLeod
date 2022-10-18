package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{ 232,54, 2, 543, 654}
	sort.Ints(arr)
	fmt.Println(arr)

	arrStr := []string{ "James", "fasd", "ab", "fdasfdas" }
	sort.Strings(arrStr)
	fmt.Println(arrStr)
}