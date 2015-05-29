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
func convert(key, alphabetSize int) (digits []int) {
    for num := key; num > 0; num = num / alphabetSize {
        remainder := num % alphabetSize
        digits = append(digits, remainder)
    }
    return reverse(digits)
}

/* Map the indices obtained from the convert and reverse functions
above into our alphabet. The alphabet is a-zA-Z0-9 */
func mapToAlphabet(digits []int, alphabetMap map[int]rune) []rune {
    var shortUrl []rune
    for _, digit := range digits {
        shortUrl = append(shortUrl, alphabetMap[digit])
    }
    return shortUrl
}

/* Create entries in a map based on what kind of letters we are
considering at a given time */
func populateAlphabetMap(alphabetMap map[int]rune, lowerLimit, upperLimit int, currentEntry rune) {
    for i := lowerLimit; i < upperLimit; i++ {
        alphabetMap[i] = currentEntry
        currentEntry++
    }
}

/* Create a map that maps our alphabet range numbers 0-62 onto
a-z, A-Z, and 0-9 respectively */
func createAlphabetMap(lowercaseLetter rune, uppercaseLetter rune, digit rune, alphabetSize int) map[int]rune {
    var alphabetMap map[int]rune = make(map[int]rune, alphabetSize)
    populateAlphabetMap(alphabetMap, 0, 26, lowercaseLetter)
    populateAlphabetMap(alphabetMap, 26, 52, uppercaseLetter)
    populateAlphabetMap(alphabetMap, 52, 62, digit)

    return alphabetMap
}

func main() {
    fmt.Println("Starting the URL shortening procedure")
    fmt.Println("Enter the key =>")
    var key int
    fmt.Scanf("%d", &key)
    const alphabetSize int = 62

    var lowercaseLetter rune = 97
    var uppercaseLetter rune = 65
    var digit rune = 48

    alphabetMap := createAlphabetMap(lowercaseLetter, uppercaseLetter, digit, alphabetSize)

    fmt.Println("Converted 125_10 to X_62,", convert(key, alphabetSize))
    fmt.Println("The following is the shortened resource; append it to an IP")
    var result []rune
    result = mapToAlphabet(convert(key, alphabetSize), alphabetMap)
    for _, resultRune := range result {
        fmt.Printf("%c", resultRune)
    }
    fmt.Println()
    fmt.Println("Done")
}
