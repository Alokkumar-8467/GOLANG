package greet

import "strings"
// If we use Capital Letter then it is Exported function.
func Hello(name string) string {
	clean := normalizeName(name)

	return "Hello," + clean
}
