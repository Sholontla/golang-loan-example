package main

import "fmt"

func Data() []int {
	data := make([]int, 0, 100)
	for n := 0; n < 100; n++ {
		data = append(data, n)
	}
	return data
}

func Test() {

	test := Data()
	for d := range test {
		fmt.Println(d)
	}
}
