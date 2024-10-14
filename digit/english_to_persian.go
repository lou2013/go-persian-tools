package digit

import "strings"

var englishToPersianMap = map[rune]rune{
	'0': '۰',
	'1': '۱',
	'2': '۲',
	'3': '۳',
	'4': '۴',
	'5': '۵',
	'6': '۶',
	'7': '۷',
	'8': '۸',
	'9': '۹',
}

func EnglishToPersian(input string) string {
	var result strings.Builder
	for _, char := range input {
		if engChar, found := englishToPersianMap[char]; found {
			result.WriteRune(engChar)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
