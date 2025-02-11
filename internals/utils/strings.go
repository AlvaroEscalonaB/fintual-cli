package utils

import (
	"regexp"
	"strings"
)

func RemoveEmojis(input string) string {
	emojiRegex := regexp.MustCompile(`[\p{So}\p{Cs}]`)
	return emojiRegex.ReplaceAllString(input, "")
}

func CleanText(input string) string {
	return strings.Trim(RemoveEmojis(input), " ")
}
