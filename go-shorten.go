package main

import (
	"fmt"
)

/* Reverse an array of digits. Implemented natively for now, although
there might be an in-built library function that does this already */
func reverse(digits []uint16) (reversed []uint16) {
    for i := len(digits) - 1; i >= 0; i-- {
		reversed = append(reversed, digits[i])
	}
	return
}

/* Declare the return array in the signature itself. The function
converts between bases and reverses the result before returning it */
func convert(key, alphabetSize uint16) (digits []uint16){
    for num := key; num > 0; num = num / alphabetSize {
        remainder := num % alphabetSize
        digits = append(digits, remainder)
    }
    return reverse(digits)
}

func main() {
	fmt.Println("Starting the URL shortening procedure")
	const key uint16 = 125
	const alphabetSize uint16 = 62
	fmt.Println("Converted 125_10 to X_62,", convert(key, alphabetSize))
}
