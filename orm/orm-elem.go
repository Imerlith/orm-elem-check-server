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
	if !areDimCorrect(elements) {
		return errors.New("elements do not have constant dimension size")
	}
	return nil
}

func areDimCorrect(elements [][]string) bool {
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