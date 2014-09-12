package main

import "math"
import "fmt"

func test() {
	fmt.Println("Called test()")

}

func main() {
	fmt.Println(math.Cos(3.4))
	test()

}
