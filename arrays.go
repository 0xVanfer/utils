package utils

import (
	"bytes"
	"strings"
)

// Connect the items in an array with "connector".
func ConnectArray(strList []string, connector string) string {
	var result bytes.Buffer
	length := len(append(strList, "aa")) - 1
	if length == 0 {
		return ""
	} else if length == 1 {
		return strList[0]
	} else {
		for i, str := range strList {
			if i == 0 {
				result.WriteString(str)
			} else {
				result.WriteString(connector)
				result.WriteString(str)
			}
		}
		return result.String()
	}
}

// Seperate a string into an array, the seperater will be deleted.
func SeperateIntoArray(targetStr string, seperateBy string) (res []string) {
	strLeft := targetStr
	for {
		if !strings.Contains(strLeft, seperateBy) {
			res = append(res, strLeft)
			return
		}
		i := strings.Index(strLeft, seperateBy)
		newStr := strLeft[:i]
		strLeft = strLeft[i+len(seperateBy):]
		res = append(res, newStr)
	}
}

func ContainInArray[T comparable](target T, targetArray []T) bool {
	for _, t := range targetArray {
		if target == t {
			return true
		}
	}
	return false
}

func RemoveRepetitionInArray[T comparable](targetArray []T) []T {
	newArray := make([]T, 0, len(targetArray))
	tempMap := make(map[T]int)
	for _, t := range targetArray {
		l := len(tempMap)
		tempMap[t] = 0
		if len(tempMap) != l {
			newArray = append(newArray, t)
		}
	}
	return newArray
}

func RemoveFromArray[T comparable](targetArray []T, targetRem T) []T {
	newArray := make([]T, 0, len(targetArray))
	for _, t := range targetArray {
		if t == targetRem {
			continue
		}
		newArray = append(newArray, t)
	}
	return newArray
}
