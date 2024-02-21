package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	charToApp := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	// Iterate over the log line
	for _, char := range log {
		// Check if the character corresponds to an application
		if app, ok := charToApp[char]; ok {
			return app
		}
	}

	// If no matching character is found, return default
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.Replace(log, string(oldRune), string(newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
