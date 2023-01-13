package main

import (
	"fmt"
	"reflect"
)

func A() string {
	return "hello"
}

func B() int {
	return 1
}

func C(a float64, b float64) float64 {
	d := a * b
	fmt.Println(d)
	return d
}

func Test(a interface{}) {

}

type Publisher interface {
	publish() error
}

func test(x interface{}) {
	v := reflect.ValueOf(x)

	if v.Kind() != reflect.Func {
		panic("Test requires a function")
	}

	t := v.Type()

	if t.NumIn() != 0 && t.NumOut() != 1 {
		panic("Function type must have no input parameters and a single return value")
	}

	values := v.Call(nil)

	val := values[0].Interface()
	fmt.Println(val)
}

func main() {

	a := Worker1{
		Name:        "user a",
		Description: "Description a",
		FloatNum:    10,
	}

	b := Worker2{
		Name:        "user b",
		Description: "Description b",
		FloatNum:    20,
	}

	c := Worker3{
		Name:        "user c",
		Description: "Description c",
		FloatNum:    30,
	}

	d := Worker4{
		Name:        "user d",
		Description: "Description d",
		FloatNum:    40,
	}

	GetProcessor(a, b, c, d)

	test(A)
	test(B)
	Test(C(2, 2))

}

// w := Worker2A{
// 	Name:     "user a",
// 	Name2:    "user a",
// 	FloatNum: 10,
// }
// sourceA := GenerateA("data string")
// processA := GetProcessorA(w)
// for i := -1; i < WorkersLenght; i++ {
// 	wg.Add(1)
// 	go processA.PostJobA(sourceA)
// 	wg.Wait()

// }
