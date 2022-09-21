package utils

import "strings"

func RemoveFromString(targetStr string, targetRem string) string {
	strLeft := targetStr
	for {
		if !strings.Contains(strLeft, targetRem) {
			return strLeft
		}
		i := strings.Index(strLeft, targetRem)
		strLeft = strLeft[:i] + strLeft[i+len(targetRem):]
	}
}

func ContainAnyOfStrs(longStr string, targets []string) bool {
	for _, containStr := range targets {
		if strings.Contains(longStr, containStr) {
			return true
		}
	}
	return false
}
