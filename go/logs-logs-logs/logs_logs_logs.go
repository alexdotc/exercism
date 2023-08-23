package logs

import "strings"
import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	x := map[rune]string{
		'\U00002757' : "recommendation",
		'\U0001F50D' : "search",
		'\U00002600' : "weather",
	}
	s := "default"
	for _, i := range log {
		v, ok := x[i]
		if ok {
			s = v
			break
		}
	}
	return s
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	replacer := func(r rune) rune {
		if r == oldRune { return newRune }
		return r
	}
	return strings.Map(replacer, log)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
