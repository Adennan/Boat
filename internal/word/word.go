package word

import (
	"strings"
	"unicode"
)

// ToUpper str to upper
func ToUpper(str string) string {
	return strings.ToUpper(str)
}

// ToLower str to lower
func ToLower(str string) string {
	return strings.ToLower(str)
}

// UnderlineToUpperCamel underline to upper camel case
func UnderlineToUpperCamel(str string) string {
	str = strings.Replace(str, "_", " ", -1)
	str = strings.Title(str)
	str = strings.Replace(str, " ", "", -1)
	return str
}

// UnderlineToLowerCamel underline to lower camel case
func UnderlineToLowerCamel(str string) string {
	str = UnderlineToUpperCamel(str)
	str = string(unicode.ToLower(rune(str[0]))) + str[1:]
	return str
}

// CamelToUnderline camel case to underline
func CamelToUnderline(str string) string {
	var ret []rune

	for k, v := range str {
		if k == 0 {
			ret = append(ret, unicode.ToLower(v))
			continue
		}
		if unicode.IsUpper(v) {
			ret = append(ret, '_')
		}
		ret = append(ret, unicode.ToLower(v))
	}

	return string(ret)
}
