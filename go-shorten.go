package main

import (
	"fmt"
)

/* Declate the return array in the signature itself.
The function converts between bases and reverses the
result before returning it */
func convert(key, alphabet uint16) (digits []uint16){
    for num := key; num > 0; num = num / key {
        remainder := num % key
        digits = append(digits, remainder)
    }
    // TODO: Reverse digits
    // TODO: Fix bugs
    return
}

func main() {
	fmt.Println("Starting the URL shortening procedure")
	const key uint16 = 125
	const alphabetSize uint16 = 62
	fmt.Println("Converted 125_10 to X_62,", convert(key, alphabetSize))
}
