package orm

import "errors"

func CheckCombinations(elements [][]string) error {
	if elements == nil {
		return errors.New("elements cannot be nil")
	}
	if len(elements) < 1 {
		return errors.New("elements cannot be empty")
	}

	return nil
}