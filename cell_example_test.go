package excelize_test

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

// ExampleCellNameToCoordinates demonstrates converting a cell name to
// column and row coordinates.
func ExampleCellNameToCoordinates() {
	col, row, err := excelize.CellNameToCoordinates("A1")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("col=%d row=%d\n", col, row)
	// Output: col=1 row=1
}

// ExampleCoordinatesToCellName demonstrates converting column and row
// coordinates to a cell name.
func ExampleCoordinatesToCellName() {
	cell, err := excelize.CoordinatesToCellName(27, 1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(cell)
	// Output: AA1
}

// ExampleColumnNameToNumber demonstrates converting a column name to its
// 1-based number.
func ExampleColumnNameToNumber() {
	num, err := excelize.ColumnNameToNumber("AA")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(num)
	// Output: 27
}

// ExampleColumnNumberToName demonstrates converting a 1-based column number
// to its name.
func ExampleColumnNumberToName() {
	name, err := excelize.ColumnNumberToName(27)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(name)
	// Output: AA
}
