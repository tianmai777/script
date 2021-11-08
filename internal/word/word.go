package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线转大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下线线转小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

// isLetter 是字母
func isLetter(r rune) bool {
	return (65 <= r && r <= 90) || (97 <= r && r <= 122)
}

// ToWord 单词转换
func ToWord(s string, f func(str string) string) string {
	var content []rune
	src := []rune(s)
	left, right := 0, 0
	for _, r := range src {
		right++
		if isLetter(r) {
			continue
		}
		content = append(content, []rune(f(string(src[left:right])))...)
		left = right
	}
	return string(content)
}
