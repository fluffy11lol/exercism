package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*~=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"([^"]*password[^"]*)"`)
	c := 0
	for _, line := range lines {
		line = strings.ToLower(line)
		if re.MatchString(line) {
			c++
		}
	}
	return c
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	newLog := re.ReplaceAllString(text, "")
	return newLog
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	var taggedLines []string
	for _, line := range lines {
		if re.MatchString(line) {
			matches := re.FindStringSubmatch(line)
			if len(matches) > 1 {
				userName := matches[1]
				taggedLine := fmt.Sprintf("[USR] %s %s", userName, line)
				taggedLines = append(taggedLines, taggedLine)
				continue
			}
		}
		taggedLines = append(taggedLines, line)
	}

	return taggedLines
}
