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

/* Create a map that maps our alphabet range numbers 0-62 onto
a-z, A-Z, and 0-9 respectively */
// TODO: Clean this method up
func createAlphabetMap(lowercaseLetter rune, uppercaseLetter rune, digit rune, alphabetSize int) map[int]rune {
    var alphabetMap map[int]rune = make(map[int]rune, alphabetSize)
    var currentEntry rune

    currentEntry = lowercaseLetter
    for i := 0; i < 26; i++ {
        alphabetMap[i] = currentEntry
        currentEntry++
    }

    currentEntry = uppercaseLetter
    for i := 26; i < 52; i++ {
        alphabetMap[i] = currentEntry
        currentEntry++
    }

    currentEntry = digit
    for i := 52; i < 62; i++ {
        alphabetMap[i] = currentEntry
        currentEntry++
    }

    return alphabetMap

}

func main() {
    fmt.Println("Starting the URL shortening procedure")
    const key int = 125
    const alphabetSize int = 62

    var lowercaseLetter rune = 97
    var uppercaseLetter rune = 65
    var digit rune = 48

    alphabetMap := createAlphabetMap(lowercaseLetter, uppercaseLetter, digit, alphabetSize)

    fmt.Printf("%c", alphabetMap[25])
    fmt.Printf("%c", alphabetMap[26])
    fmt.Printf("%c", alphabetMap[54])
    fmt.Println()

    fmt.Println("Converted 125_10 to X_62,", convert(key, alphabetSize))
    fmt.Println(mapToAlphabet(convert(key, alphabetSize), alphabetMap))
}
