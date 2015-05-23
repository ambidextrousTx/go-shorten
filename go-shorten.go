/* Implement a URL shortening algorithm in GoLang based on a tutorial
on Stack Overflow */

package main

import (
    "fmt"
)

/* Reverse an array of digits. Implemented natively for now, although
there might be an in-built library function that does this already */
func reverse(digits []int) (reversed []int) {
    for i := len(digits) - 1; i >= 0; i-- {
        reversed = append(reversed, digits[i])
    }
    return
}

/* Declare the return array in the signature itself. The function
converts between bases and reverses the result before returning it */
func convert(key, alphabetSize int) (digits []int){
    for num := key; num > 0; num = num / alphabetSize {
        remainder := num % alphabetSize
        digits = append(digits, remainder)
    }
    return reverse(digits)
}

/* Map the indices obtained from the convert and reverse functions
above into our alphabet. The alphabet is a-zA-Z0-9 */
func mapToAlphabet(digits []int, alphabetMap map[int]rune) string {
    var shortUrl string
    for _, digit := range digits {
        //shortUrl += alphabetMap[digit]
        fmt.Println(digit)
    }
    return shortUrl
}

func main() {
    fmt.Println("Starting the URL shortening procedure")
    const key int = 125
    const alphabetSize int = 62
    var alphabetMap map[int]rune = make(map[int]rune, alphabetSize)

    var letter rune = 97
    for i := 0; i < 26; i++ {
        alphabetMap[i] = letter
        letter++
    }

    fmt.Printf("%c", alphabetMap[25])
    fmt.Println()

    fmt.Println("Converted 125_10 to X_62,", convert(key, alphabetSize))
    fmt.Println(mapToAlphabet(convert(key, alphabetSize), alphabetMap))
}
