package main

import (
	"math"
	"fmt"
	"encoding/base64"
)

func test() {
	fmt.Println("Called test()")

}

func main() {
	// testing out base64 encoding of URLs
	url := []byte("http://www.golang.com")
	str := base64.StdEncoding.EncodeToString(url)
	fmt.Println(str)
	
	fmt.Println(math.Cos(3.4))
	test()

}
