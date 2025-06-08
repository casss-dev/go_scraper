package util

import (
	"regexp"
	"strconv"
	"strings"
)

// parses a float from a price string
// Example: "US $1,239.99" => 1239.99
func ParsePrice(price string) (float64, error) {
	p := StripCurrencySymbol(price)
	return strconv.ParseFloat(p, 64)
}

// strips currency symbols from a price string
func StripCurrencySymbol(price string) string {
	var p string
	p = strings.ReplaceAll(price, "$", "")
	p = strings.ReplaceAll(price, ",", "")
	re := regexp.MustCompile(`[0-9.]+`)
	return re.FindString(p)
}

func WrapWords(title string, maxLineLength int) string {
	words := strings.Split(title, " ")
	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word)+1 <= maxLineLength {
			if currentLine == "" {
				currentLine += word
			} else {
				currentLine += " " + word
			}
		} else {
			lines = append(lines, currentLine)
			currentLine = word
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return strings.Join(lines, "\n")
}
