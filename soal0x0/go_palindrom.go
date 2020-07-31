package soal0x0

import "fmt"

func main() {
	var first, last int
	fmt.Print("first number: ")
	fmt.Scan(&first)
	fmt.Print("last number: ")
	fmt.Scan(&last)
	var res int = CountTotalPalindrom(first, last)
	fmt.Printf("Total palindrom: %d\n", res)
}

func CountTotalPalindrom(start int, end int) int {
	var totalpalindrom int = 0
	for i := start; i < end; i++ {
		var originalNumber int = i
		var reverseNumber int = 0
		var temp = originalNumber

		for temp > 0 {
			var remainder = temp % 10
			reverseNumber = (reverseNumber * 10) + remainder
			temp = temp / 10
		}

		if originalNumber == reverseNumber {
			totalpalindrom++;
		}
	}
	return totalpalindrom
}
