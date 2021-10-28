package orm

import "testing"

func Test_CheckCombinations_Empty_Elems(t *testing.T) {
	var elems = [][]string{}

	err := CheckCombinations(elems)

	if err == nil {
		t.Fatal("Should pass an error")
	}

	if err.Error() != "elements cannot be empty" {
		t.Fatal("Wrong error message")
	}
}