package main

import "fmt"

type tester interface {
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

type rangeTest struct {
	min int
	max int
}

func (rt rangeTest) test(i int) bool {
	return rt.min <= i && i <= rt.max
}

type divTest int

func (dt divTest) test(i int) bool {
	return i%int(dt) == 0
}

type testerfunc func(int) bool

func (tf testerfunc) test(i int) bool {
	return tf(i)
}

func main() {
	result := runTests(10, []tester{
		testerfunc(func(i int) bool { return i%2 == 0 }),
		testerfunc(func(i int) bool { return i <= 20 }),
	})
	fmt.Println(result)
}

//func main() {
//	result := runTests(10, []tester{
//		rangeTest{min: 5, max: 20},
//		divTest(5),
//	})
//	fmt.Println(result)
//}
