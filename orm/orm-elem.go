package orm

import (
	"errors"
	"strings"
)

func CheckCombinations(elements [][]string) error {
	if elements == nil {
		return errors.New("elements cannot be nil")
	}
	if len(elements) < 2 {
		return errors.New("elements cannot have less than 2 rows")
	}
	if len(elements[0]) < 2 {
		return errors.New("elements cannot have less than 2 columns")
	}
	if !areDimsCorrect(elements) {
		return errors.New("elements do not have constant dimension size")
	}

	//TODO: Add spliting
	var splitedElements = [][]string{}
	_ = removeDuplicates(splitedElements)


	return nil
}

func areDimsCorrect(elements [][]string) bool {
	maxSize := len(elements[0])
	for i := 1; i < len(elements); i++ {
		rowSize := 0
		for _, element := range elements[i] {
			if strings.TrimSpace(element) != "" {
				rowSize++
			}
		}
		if maxSize != rowSize {
			return false
		}
	}
	return true
}

func removeDuplicates(elements [][]string) [][]string {
	var elemsWithoutDuplicates = [][]string{}
	var rowsIds = []string{}
	for _, row := range elements {
		rowId := strings.Join(row, "")
		if isElementPresent(rowId, rowsIds) {
			continue
		}
		rowsIds = append(rowsIds, rowId)
		elemsWithoutDuplicates = append(elemsWithoutDuplicates, row)
	}
	return elemsWithoutDuplicates
}

func isElementPresent(element string, arr []string) bool {
	for _, ae := range arr {
		if element == ae {
			return true
		}
	}
	return false
}

func getColumn(elements [][]string, columnNumber int) ([]string, error) {
	if elements == nil {
		return nil, errors.New("elements cannot be nil")
	}
	if len(elements) < 1 {
		return nil, errors.New("elements must contain at least one row")
	}
	if columnNumber < 0 || columnNumber >= len(elements[0]) {
		return nil, errors.New("columnNumber must be in <0, len)")
	}
	column := []string{}
	for _, row := range elements {
		column = append(column, row[columnNumber])
	}
	return column, nil
}
