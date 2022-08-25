// Package leap provides IsLeapYear function
package leap

// IsLeapYear calculates whether the year is a leap year and return true if it's leap, otherwise false
func IsLeapYear(year int) bool {
	isDivisibleBy4 := year%4 == 0
	isDivisibleB100 := year%100 == 0
	isDivisibleBy400 := year%400 == 0
	return isDivisibleBy4 && !isDivisibleB100 || isDivisibleBy400
}
