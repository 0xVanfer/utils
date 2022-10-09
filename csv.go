package utils

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

// Read the csv file.
func ReadCsv(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	return csvReader.ReadAll()
}

// If file exist, will cover the old file.
func WriteNewCsv(filePath string, title []string, text [][]string) error {
	// regular check
	if len(text) == 0 {
		return errors.New("text is empty")
	}
	if !(len(title) == len(text[0])) {
		return errors.New("title and text length not match")
	}
	// create new file
	csvFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	// write title
	err = csvWriter.Write(title)
	if err != nil {
		return err
	}
	// write text
	err = csvWriter.WriteAll(text)
	if err != nil {
		return err
	}
	// flush
	csvWriter.Flush()
	return csvWriter.Error()
}

// Write to csv line by line.
func WriteToCsv(filePath string, title []string, text []string) error {
	if len(text) != len(title) {
		return errors.New("text length not match")
	}
	// read existing file, if not exist, create a new one
	csvFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		// create new file
		csvFile, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	csvWriter := csv.NewWriter(csvFile)

	// file text already exist
	fileText, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	// whether the file already has title
	if len(fileText) == 0 {
		// write title
		err = csvWriter.Write(title)
		if err != nil {
			return err
		}
		// write text
		err = csvWriter.Write(text)
		if err != nil {
			return err
		}
	} else {
		// regular check
		if len(fileText[0]) != len(title) {
			return errors.New("title length not match")
		}
		// write text
		err = csvWriter.Write(text)
		if err != nil {
			return err
		}
	}
	// flush
	csvWriter.Flush()
	return csvWriter.Error()
}

// Write to csv line by line.
//
// Will not write if the value of "noRepeatTitle" already exists.
func WriteToCsvResume(filePath string, title []string, text []string, noRepeatTitle string) error {
	if len(text) != len(title) {
		return errors.New("text length not match")
	}
	// read existing file, if not exist, create a new one
	csvFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		// create new file
		csvFile, err = os.Create(filePath)
		if err != nil {
			return err
		}
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	csvWriter := csv.NewWriter(csvFile)

	// file text already exist
	fileText, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	// whether the file already has title
	if len(fileText) == 0 {
		// write title
		err = csvWriter.Write(title)
		if err != nil {
			return err
		}
		// write text
		err = csvWriter.Write(text)
		if err != nil {
			return err
		}
	} else {
		// regular check
		if len(fileText[0]) != len(title) {
			return errors.New("title length not match")
		}
		var columeIndex int
		for index, ti := range title {
			if strings.EqualFold(ti, noRepeatTitle) {
				columeIndex = index
			}
		}
		// if repeat, return
		for _, line := range fileText {
			if strings.EqualFold(line[columeIndex], text[columeIndex]) {
				return nil
			}
		}
		// write text
		err = csvWriter.Write(text)
		if err != nil {
			return err
		}
	}
	// flush
	csvWriter.Flush()
	return csvWriter.Error()
}
