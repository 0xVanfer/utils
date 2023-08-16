package utils

import (
	"bytes"
	"strings"
)

// Connect the items in an array with "connector".
//
// Example:
//
//	ConnectArray([]string{"one", "year", "old"}..., "-") = "one-year-old"
func ConnectArray(connector string, strList ...string) string {
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
//
// Example:
//
//	SeperateIntoArray("one-year-old", "-") = []string{"one", "year", "old"}
func SeperateIntoArray(str string, seperateBy string) (res []string) {
	strLeft := str
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
//
// Example:
//
//	ContainInArray(1,[]int{1, 2, 3}) = true
//	ContainInArray("", []string{"1", "2", "3"}...) = false
//	ContainInArray("0", []string{"1", "2", "3"}...) = false
func ContainInArray[T comparable](target T, array ...T) bool {
	for _, t := range array {
		if target == t {
			return true
		}
	}
	return false
}

// If target contains in string array, ignore lower or upper.
//
// Example:
//
//	ContainInArrayX("aa", []string{"aa", "bb", "cc"}...) = true
//	ContainInArrayX("aA", []string{"aa", "bb", "cc"}...) = true
func ContainInArrayX(target string, array ...string) bool {
	for _, t := range array {
		if strings.EqualFold(target, t) {
			return true
		}
	}
	return false
}

// Remove repetiton in an array and return the new one.
//
// Example:
//
//	RemoveRepetitionInArray([]string{"aa", "bb", "bb", "BB"}...) = []string{"aa", "bb", "BB"}
//	RemoveRepetitionInArray([]int{1, 2, 3, 3}...) = []int{1, 2, 3}
func RemoveRepetitionInArray[T comparable](array ...T) []T {
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
//
// Example:
//
//	RemoveFromArray("aa", []string{"aa", "bb"}...) = []string{"bb"}
//	RemoveFromArray(1, []uint64{1}...) = []uint64{}
func RemoveFromArray[T comparable](toRemove T, array ...T) []T {
	newArray := make([]T, 0, len(array))
	for _, t := range array {
		if t == toRemove {
			continue
		}
		newArray = append(newArray, t)
	}
	return newArray
}

// Return the AND result of two arrays.
func ArrayAnd[T comparable](array0, array1 []T) []T {
	newArray := make([]T, 0, Max(len(array0), len(array1)))
	for _, t := range array0 {
		if ContainInArray(t, array1...) {
			newArray = append(newArray, t)
		}
	}
	newArray = RemoveRepetitionInArray(newArray...)
	return newArray
}

// Return the OR result of two arrays.
func ArrayOr[T comparable](array0, array1 []T) []T {
	newArray := make([]T, 0, (len(array0) + len(array1)))
	newArray = append(newArray, array0...)
	newArray = append(newArray, array1...)
	newArray = RemoveRepetitionInArray(newArray...)
	return newArray
}
