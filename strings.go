package utils

import "strings"

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
