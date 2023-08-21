// Package twofer blah blah blah
package twofer

import "fmt"

// ShareWith shares cookies with people that like cookies
func ShareWith(name string) string {
	if (name == "") { name = "you" }
	return fmt.Sprintf("One for %s, one for me.", name)
}
