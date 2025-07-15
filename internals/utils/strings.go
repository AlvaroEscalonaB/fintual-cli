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

func FormatCurrency(amount string) string {
	if amount == "" {
		return "0.0"
	}

	parts := strings.Split(amount, ".")
	integerPart := parts[0]
	decimalPart := "0"
	
	if len(parts) > 1 {
		decimalPart = parts[1]
		if len(decimalPart) > 2 {
			decimalPart = decimalPart[:2]
		} else if len(decimalPart) == 1 {
			decimalPart += "0"
		}
	}

	// Add thousand separators to integer part
	length := len(integerPart)
	if length <= 3 {
		return integerPart + "." + decimalPart
	}

	var result []byte
	for i := 0; i < length; i++ {
		if i > 0 && (length - i) % 3 == 0 {
			result = append(result, '.')
		}
		result = append(result, integerPart[i])
	}

	return string(result) + "," + decimalPart
}
