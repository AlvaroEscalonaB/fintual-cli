package utils_test

import (
	"fintual-cli/internals/utils"
	"testing"
)

func TestRemoveEmojis(t *testing.T) {
	stringWithEmoji := "🏦 Content"
	cleanText := utils.RemoveEmojis(stringWithEmoji)
	finalText := " Content"

	if cleanText != finalText {
		t.Fatalf("Left does not match with right value\n%s != %s ", stringWithEmoji, cleanText)
	}
}

func TestCleanText(t *testing.T) {
	stringWithEmoji := "🏡 Casa"
	cleanText := utils.CleanText(stringWithEmoji)
	finalText := "Casa"

	if cleanText != finalText {
		t.Fatalf("Left does not match with right value\n%s != %s ", stringWithEmoji, cleanText)
	}
}
