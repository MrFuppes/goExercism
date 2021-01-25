// Package twofer implements sharing of one between two
package twofer

import "fmt"

// ShareWith - given a name, returns a string saying between who things are shared.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
