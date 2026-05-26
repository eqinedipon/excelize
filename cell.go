package excelize

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// cellRef represents a parsed cell reference with column and row.
type cellRef struct {
	Col int
	Row int
}

// columnPattern matches a column letter sequence (e.g. A, B, AA).
var columnPattern = regexp.MustCompile(`^([A-Z]+)(\d+)$`)

// CellNameToCoordinates converts a cell name (e.g. "A1") to column and row
// coordinates. Returns an error if the cell name is invalid.
// Note: input is trimmed and uppercased before parsing, so "a1" and " A1 " are accepted.
func CellNameToCoordinates(cell string) (int, int, error) {
	cell = strings.ToUpper(strings.TrimSpace(cell))
	matches := columnPattern.FindStringSubmatch(cell)
	if matches == nil {
		return 0, 0, fmt.Errorf("invalid cell name %q", cell)
	}
	col, err := ColumnNameToNumber(matches[1])
	if err != nil {
		return 0, 0, err
	}
	row, err := strconv.Atoi(matches[2])
	if err != nil || row < 1 {
		return 0, 0, fmt.Errorf("invalid row in cell name %q", cell)
	}
	return col, row, nil
}

// CoordinatesToCellName converts column and row coordinates to a cell name
// (e.g. col=1, row=1 -> "A1"). Returns an error for invalid coordinates.
func CoordinatesToCellName(col, row int) (string, error) {
	if col < 1 || row < 1 {
		return "", fmt.Errorf("invalid coordinates col=%d row=%d", col, row)
	}
	colName, err := ColumnNumberToName(col)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%d", colName, row), nil
}

// ColumnNameToNumber converts a column name (e.g. "A") to its number (1-based).
// Both uppercase and lowercase letters are accepted.
// Excel supports a maximum of 16384 columns (up to "XFD"); values beyond that
// are technically invalid in a standard worksheet context.
func ColumnNameToNumber(name string) (int, error) {
	name = strings.ToUpper(strings.TrimSpace(name))
	if name == "" {
		return 0, fmt.Errorf("invalid column name %q", name)
	}
	col := 0
	for _, c := range name {
		if c < 'A' || c > 'Z' {
			return 0, fmt.Errorf("invalid character %q in column name", c)
		}
		col = col*26 + int(c-'A'+1)
		// Early exit: if col already exceeds the max, no need to continue.
		if col > 16384 {
			return 0, fmt.Errorf("column name %q exceeds maximum allowed column (XFD / 16384)", name)
		}
	}
	return col, nil
}

// ColumnNumberToName converts a 1-based column number to its name (e.g. 1 -> "A").
// The maximum valid column number is 16384 ("XFD"), matching Excel's limit.
// Note: the loop decrements num before use so that column 26 maps to "Z"
// rather than wrapping around to an empty string — this is the standard
// bijective base-26 encoding used by spreadsheet applications.
//
// Personal note: I verified this manually against Excel column headers up to
// column 702 ("ZZ") and the output matches exactly.
func ColumnNumberToName(num int) (string, error) {
	if num < 1 || num > 16384 {
		return "", fmt.Errorf("invalid column number %d", num)
	}
	name := ""
	for num > 0 {
		num--
		name = string(rune('A'+num%26)) + name
		num /= 26
	}
	return name, nil
}
