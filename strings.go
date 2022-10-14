package utils

import (
	"bytes"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Remove the first subStr in str.
func RemoveFromString(str string, subStr string) string {
	strLeft := str
	for {
		if !strings.Contains(strLeft, subStr) {
			return strLeft
		}
		i := strings.Index(strLeft, subStr)
		strLeft = strLeft[:i] + strLeft[i+len(subStr):]
	}
}

func ContainAnyOfStrs(str string, subStrs []string) bool {
	for _, subStr := range subStrs {
		if strings.Contains(str, subStr) {
			return true
		}
	}
	return false
}

// The first letter to upper case.
func UpperFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// The first letter to lower case.
func LowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// Return seperated words of str.
func StrUnderline2Seperate(str string) string {
	return strings.Replace(str, "_", " ", -1)
}

// Return title case seperated words of str.
func StrUnderline2SeperateTitle(str string) string {
	return cases.Title(language.Und).String(StrUnderline2Seperate(str))
}

// Return camel case of str.
//
// Warning: Using both StrUnderline2Camel and StrCamel2Underline(regardless of order) will probably not get the original string!
func StrUnderline2Camel(str string) string {
	return LowerFirst(strings.Replace(StrUnderline2SeperateTitle(str), " ", "", -1))
}

// Return underline case of str.
//
// Warning: Using both StrUnderline2Camel and StrCamel2Underline(regardless of order) will probably not get the original string!
func StrCamel2Underline(str string) string {
	var result bytes.Buffer
	for i, r := range str {
		if unicode.IsUpper(r) {
			if i != 0 {
				result.WriteString("_")
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
