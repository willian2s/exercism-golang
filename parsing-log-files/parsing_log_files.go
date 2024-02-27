package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^\[(TRC|DBG|INF|ERR|FTL)\]`)
	if err != nil {
		panic(err)
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<([~*=-]+|)>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	regex := regexp.MustCompile(`(?i)".*password.*"`)

	var count int
	for _, line := range lines {
		if regex.MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +(\w+)`)
	taggedLines := make([]string, len(lines))

	for i, line := range lines {
		find := re.FindStringSubmatch(line)

		if len(find) > 0 {
			taggedLines[i] = fmt.Sprintf("[USR] %s %s", find[1], line)
		} else {
			taggedLines[i] = line
		}
	}

	return taggedLines
}
