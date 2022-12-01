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

// Example:
//
//	ContainAnyOfStrs("abracadabra", []string{"bra", "tnt"}) = true
//	ContainAnyOfStrs("abracadabra", []string{"abc", "tnt"}) = false
func ContainAnyOfStrs(str string, subStrs []string) bool {
	for _, subStr := range subStrs {
		if strings.Contains(str, subStr) {
			return true
		}
	}
	return false
}

// Example:
//
//	ContainAnyOfStrs("abracadabra", []string{"bra", "abr"}) = true
//	ContainAnyOfStrs("abracadabra", []string{"bra", "tnt"}) = false
func ContainAllOfStrs(str string, subStrs []string) bool {
	for _, subStr := range subStrs {
		if !strings.Contains(str, subStr) {
			return false
		}
	}
	return true
}

// The first letter to upper case.
//
// Example:
//
//	UpperFirst("upper") = "Upper"
func UpperFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// The first letter to lower case.
//
// Example:
//
//	LowerFirst("Lower") = "lower"
func LowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// Return seperated words of str.
//
// Example:
//
//	StrUnderline2Seperate("a_b_c") = "a b c"
func StrUnderline2Seperate(str string) string {
	return strings.Replace(str, "_", " ", -1)
}

// Return title case seperated words of str.
//
// Example:
//
//	StrUnderline2SeperateTitle("a_b_c") = "A B C"
func StrUnderline2SeperateTitle(str string) string {
	return cases.Title(language.Und).String(StrUnderline2Seperate(str))
}

// Return camel case of str. The first letter is upper case.
//
// Example:
//
//	StrUnderline2Camel("a_b_c") = "ABC"
//
// NOTE:
//
//	Using both `StrUnderline2Camel()` and `StrCamel2Underline()`
//	(regardless of order) will probably not get the original string!
func StrUnderline2CamelUpperFirst(str string) string {
	return strings.Replace(StrUnderline2SeperateTitle(str), " ", "", -1)
}

// Return camel case of str. The first letter is lower case.
//
// Example:
//
//	StrUnderline2Camel("a_b_c") = "aBC"
//
// NOTE:
//
//	Using both `StrUnderline2Camel()` and `StrCamel2Underline()`
//	(regardless of order) will probably not get the original string!
func StrUnderline2Camel(str string) string {
	return LowerFirst(StrUnderline2CamelUpperFirst(str))
}

// Return underline case of str.
//
// Example:
//
//	StrUnderline2Camel("aBC") = "a_b_c"
//	StrUnderline2Camel("ABC") = "a_b_c"
//
// NOTE:
//
//	Using both `StrUnderline2Camel()` and `StrCamel2Underline()`
//	(regardless of order) will probably not get the original string!
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
