package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	logLevels := []string{
		"[TRC]",
		"[DBG]",
		"[INF]",
		"[WRN]",
		"[ERR]",
		"[FTL]"}

	for _, level := range logLevels {
		if strings.HasPrefix(text, level) {
			return true
		}
	}

	return false
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile("<[-|~|\\*|=]*>").Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		matches := regexp.MustCompile(`(?i)".*password.*"`).FindAllStringIndex(line, -1)
		count += len(matches)
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`(end-of-line\d+)`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	for i, line := range lines {
		user := regexp.MustCompile(`(User\s+(\w*))`).FindStringSubmatch(line)
		if len(user) >= 2 {
			lines[i] = "[USR] " + user[2] + " " + line
		}
	}

	return lines
}
