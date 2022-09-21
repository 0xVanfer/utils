package utils

import (
	"bytes"
	"strings"
)

// find all string between "strBefore" and "strAfter"
func FindStringBetween(fullStr string, strBefore string, strAfter string) (target []string) {
	strLeft := fullStr
	for {
		if !strings.Contains(strLeft, strBefore) {
			return
		}
		i1 := strings.Index(strLeft, strBefore)
		strLeft = strLeft[i1+len(strBefore):]
		if !strings.Contains(strLeft, strAfter) {
			return
		}
		i2 := strings.Index(strLeft, strAfter)
		newStr := strLeft[:i2]
		target = append(target, newStr)
		strLeft = strLeft[len(strAfter)+i2:]
	}
}

// find all []byte between "bytesBefore" and "bytesAfter"
func FindBytesBetween(fullBytes []byte, bytesBefore []byte, bytesAfter []byte) (target [][]byte) {
	bytesLeft := fullBytes
	for {
		if !bytes.Contains(bytesLeft, bytesAfter) {
			return
		}
		i1 := bytes.Index(bytesLeft, bytesBefore)
		bytesLeft = bytesLeft[i1+len(bytesBefore):]
		if !bytes.Contains(bytesLeft, bytesAfter) {
			return
		}
		i2 := bytes.Index(bytesLeft, bytesAfter)
		newBytes := bytesLeft[:i2]
		target = append(target, newBytes)
		bytesLeft = bytesLeft[len(bytesAfter)+i2:]
	}
}
