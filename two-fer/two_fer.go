// Package twofer implements a function to return a message.
package twofer

import "fmt"

// ShareWith method	returns a message, based on arguments passed into it.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
