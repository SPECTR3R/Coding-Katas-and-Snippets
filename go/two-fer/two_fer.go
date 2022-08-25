// Package twofer implements a simple ShareWith function.
package twofer

import "fmt"

// ShareWith given a {name} returns a string with the format "One for {name}, one for me."
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
