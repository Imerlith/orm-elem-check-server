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


func Test_getColumn_nilElements(t* testing.T) {
	column, err := getColumn(nil, 1)

	if err == nil {
		t.Fatal("Should pass an error")
	}
	if err.Error() != "elements cannot be nil" {
		t.Fatal("Wrong error message")
	}
	if column != nil {
		t.Fatal("Column should be nil")
	}
}

func Test_getColumn_emptyElements(t* testing.T) {
	elements := [][]string{}

	column, err := getColumn(elements, 1)

	if err == nil {
		t.Fatal("Should pass an error")
	}
	if err.Error() != "elements must contain at least one row" {
		t.Fatal("Wrong error message")
	}
	if column != nil {
		t.Fatal("Column should be nil")
	}
}

func Test_getColumn_columnNumberEqualLen(t* testing.T) {
	elements := [][]string{
		{
			"a1", "b1", "c1",
		},
	}

	column, err := getColumn(elements, 3)

	if err == nil {
		t.Fatal("Should pass an error")
	}
	if err.Error() != "columnNumber must be in <0, len)" {
		t.Fatal("Wrong error message")
	}
	if column != nil {
		t.Fatal("Column should be nil")
	}
}

func Test_getColumn_columnNumberGreaterThanLen(t* testing.T) {
	elements := [][]string{
		{
			"a1", "b1", "c1",
		},
	}

	column, err := getColumn(elements, 5)

	if err == nil {
		t.Fatal("Should pass an error")
	}
	if err.Error() != "columnNumber must be in <0, len)" {
		t.Fatal("Wrong error message")
	}
	if column != nil {
		t.Fatal("Column should be nil")
	}
}

func Test_getColumn_columnNumberLessThanZero(t* testing.T) {
	elements := [][]string{
		{
			"a1", "b1", "c1",
		},
	}

	column, err := getColumn(elements, -1)

	if err == nil {
		t.Fatal("Should pass an error")
	}
	if err.Error() != "columnNumber must be in <0, len)" {
		t.Fatal("Wrong error message")
	}
	if column != nil {
		t.Fatal("Column should be nil")
	}
}

func Test_getColumn_correctElemsAndNumber(t* testing.T) {
	elements := [][]string{
		{
			"a1", "b1", "c1",
		},
		{
			"a2", "b2", "c2",
		},
		{
			"a3", "b3", "c3",
		},
	}

	column, err := getColumn(elements, 1)
	expectedColumn := []string{
		"b1", "b2", "b3",
	}

	if err != nil {
		t.Fatal("Should not pass an error")
	}
	if column == nil {
		t.Fatal("Column should not be nil")
	}
	if len(expectedColumn) != len(column)  {
		t.Fatalf("Wrong column size received. Should be: %v but is %v", len(expectedColumn), len(column))
	}

	for index, item := range column {
		if item != expectedColumn[index] {
			t.Fatalf("Column element missmatch. Should be: %v but is %v", expectedColumn[index], item)
		}
	}
}
