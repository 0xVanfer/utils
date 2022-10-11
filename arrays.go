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

// If target contains in array.
func ContainInArray[T comparable](target T, array []T) bool {
	for _, t := range array {
		if target == t {
			return true
		}
	}
	return false
}

// If target contains in string array, ignore lower or upper.
func ContainInArrayX(target string, array []string) bool {
	for _, t := range array {
		if strings.EqualFold(target, t) {
			return true
		}
	}
	return false
}

// Remove repetiton in an array and return the new one.
func RemoveRepetitionInArray[T comparable](array []T) []T {
	newArray := make([]T, 0, len(array))
	tempMap := make(map[T]int)
	for _, t := range array {
		l := len(tempMap)
		tempMap[t] = 0
		if len(tempMap) != l {
			newArray = append(newArray, t)
		}
	}
	return newArray
}

// Remove sth from an array and return the new one.
func RemoveFromArray[T comparable](array []T, toRemove T) []T {
	newArray := make([]T, 0, len(array))
	for _, t := range array {
		if t == toRemove {
			continue
		}
		newArray = append(newArray, t)
	}
	return newArray
}
