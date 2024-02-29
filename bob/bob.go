// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case remark == "":
		return "Fine. Be that way!"
	case strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark:
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	case strings.HasSuffix(remark, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}
