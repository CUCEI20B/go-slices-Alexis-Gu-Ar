package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]int, 0, n)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		s = append(s, num)
	}

	var sum int
	for _, num := range s {
		sum += num
	}

	fmt.Print(sum)
}
