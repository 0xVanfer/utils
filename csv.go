package utils

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

// Read the csv file from the given path.
// The first line of the result is the title.
func ReadCsv(filePath string) (contents [][]string, err error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	return csvReader.ReadAll()
}

// Write a new csv file.
// If file exist, will cover the old file.
func WriteNewCsv(filePath string, title []string, text [][]string) error {
	// Regulatory check.
	if len(text) == 0 {
		return errors.New("text is empty")
	}
	if !(len(title) == len(text[0])) {
		return errors.New("title and text length not match")
	}
	// Create new file.
	csvFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	// Write title.
	err = csvWriter.Write(title)
	if err != nil {
		return err
	}
	// Write text.
	err = csvWriter.WriteAll(text)
	if err != nil {
		return err
	}
	// Flush.
	csvWriter.Flush()
	return csvWriter.Error()
}

// Write to csv line by line.
func WriteToCsv(filePath string, title []string, text []string) error {
	if len(text) != len(title) {
		return errors.New("text length not match")
	}
	// Read the existing file.
	// If does not exist, create a new one.
	csvFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		// Create new file.
		csvFile, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	csvWriter := csv.NewWriter(csvFile)

	// File text already exist.
	fileText, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	// Whether the file already has title.
	if len(fileText) == 0 {
		// Write title.
		err = csvWriter.Write(title)
		if err != nil {
			return err
		}
		// Write text.
		err = csvWriter.Write(text)
		if err != nil {
			return err
		}
	} else {
		// Regulatory check.
		if len(fileText[0]) != len(title) {
			return errors.New("title length not match")
		}
		// Write text.
		err = csvWriter.Write(text)
		if err != nil {
			return err
		}
	}
	// Flush.
	csvWriter.Flush()
	return csvWriter.Error()
}

// Write to csv line by line.
//
// Will not write the text if the value of "noRepeatTitle" already exists.
//
// For example, the csv already exists, and there is a column named "name"
// with values: Alice, Bob. `noRepeatTitle` is "name".
// If the function is called with "name" value "Alice", this line will be skipped.
func WriteToCsvResume(filePath string, title []string, text []string, noRepeatTitle string) error {
	if len(text) != len(title) {
		return errors.New("text length not match")
	}
	// Read the existing file.
	// If does not exist, create a new one.
	csvFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		// Create new file.
		csvFile, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	csvWriter := csv.NewWriter(csvFile)

	// File text already exist.
	fileText, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	// Whether the file already has title.
	if len(fileText) == 0 {
		// Write title.
		err = csvWriter.Write(title)
		if err != nil {
			return err
		}
		// Write text.
		err = csvWriter.Write(text)
		if err != nil {
			return err
		}
	} else {
		// Regulatory check.
		if len(fileText[0]) != len(title) {
			return errors.New("title length not match")
		}
		var columeIndex int
		for index, ti := range title {
			if strings.EqualFold(ti, noRepeatTitle) {
				columeIndex = index
			}
		}
		// If repeat, return.
		for _, line := range fileText {
			if strings.EqualFold(line[columeIndex], text[columeIndex]) {
				return nil
			}
		}
		// Write text.
		err = csvWriter.Write(text)
		if err != nil {
			return err
		}
	}
	// Flush.
	csvWriter.Flush()
	return csvWriter.Error()
}
