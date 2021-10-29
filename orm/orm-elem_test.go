package orm

import "testing"

func Test_CheckCombinations_Empty_Elems(t *testing.T) {
	var elems = [][]string{}

	err := CheckCombinations(elems)

	if err == nil {
		t.Fatal("Should pass an error")
	}

	if err.Error() != "elements cannot have less than 2 rows" {
		t.Fatal("Wrong error message")
	}
}

func Test_CheckCombinations_Null_Elems(t *testing.T) {
	err := CheckCombinations(nil)

	if err == nil {
		t.Fatal("Should pass an error")
	}

	if err.Error() != "elements cannot be nil" {
		t.Fatal("Wrong error message")
	}
}

func Test_CheckCombinations_OneRow_Elems(t *testing.T) {
	var elems = [][]string{
		{
			"a1", "b1", "c1",
		},
	}

	err := CheckCombinations(elems)

	if err == nil {
		t.Fatal("Should pass an error")
	}

	if err.Error() != "elements cannot have less than 2 rows" {
		t.Fatal("Wrong error message")
	}
}

func Test_CheckCombinations_OneColumn_Elems(t *testing.T) {
	var elems = [][]string{
		{
			"a1",
		},
		{
			"a2",
		},
		{
			"a3",
		},
	}

	err := CheckCombinations(elems)

	if err == nil {
		t.Fatal("Should pass an error")
	}

	if err.Error() != "elements cannot have less than 2 columns" {
		t.Fatal("Wrong error message")
	}
}

func Test_CheckCombinations_WrongDimensions_TooShort(t *testing.T) {
	var elems = [][]string{
		{
			"a1", "b1", "c1",
		},
		{
			"a2", "b2", "",
		},
		{
			"a3", "b3", "c3",
		},
	}

	err := CheckCombinations(elems)

	if err == nil {
		t.Fatal("Should pass an error")
	}

	if err.Error() != "elements do not have constant dimension size" {
		t.Fatal("Wrong error message")
	}
}

func Test_CheckCombinations_WrongDimensions_TooLong(t *testing.T) {
	var elems = [][]string{
		{
			"a1", "b1", "c1",
		},
		{
			"a2", "b2", "c1",
		},
		{
			"a3", "b3", "c3", "d4",
		},
	}

	err := CheckCombinations(elems)

	if err == nil {
		t.Fatal("Should pass an error")
	}

	if err.Error() != "elements do not have constant dimension size" {
		t.Fatal("Wrong error message")
	}
}
