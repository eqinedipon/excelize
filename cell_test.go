package excelize

import (
	"testing"
)

func TestColumnNameToNumber(t *testing.T) {
	cases := []struct {
		name   string
		want   int
		wantErr bool
	}{
		{"A", 1, false},
		{"Z", 26, false},
		{"AA", 27, false},
		{"AZ", 52, false},
		{"BA", 53, false},
		{"XFD", 16384, false},
		{"", 0, true},
		{"1A", 0, true},
	}
	for _, tc := range cases {
		got, err := ColumnNameToNumber(tc.name)
		if tc.wantErr {
			if err == nil {
				t.Errorf("ColumnNameToNumber(%q): expected error, got nil", tc.name)
			}
			continue
		}
		if err != nil {
			t.Errorf("ColumnNameToNumber(%q): unexpected error: %v", tc.name, err)
			continue
		}
		if got != tc.want {
			t.Errorf("ColumnNameToNumber(%q) = %d, want %d", tc.name, got, tc.want)
		}
	}
}

func TestColumnNumberToName(t *testing.T) {
	cases := []struct {
		num     int
		want    string
		wantErr bool
	}{
		{1, "A", false},
		{26, "Z", false},
		{27, "AA", false},
		{16384, "XFD", false},
		{0, "", true},
		{-1, "", true},
	}
	for _, tc := range cases {
		got, err := ColumnNumberToName(tc.num)
		if tc.wantErr {
			if err == nil {
				t.Errorf("ColumnNumberToName(%d): expected error, got nil", tc.num)
			}
			continue
		}
		if err != nil {
			t.Errorf("ColumnNumberToName(%d): unexpected error: %v", tc.num, err)
			continue
		}
		if got != tc.want {
			t.Errorf("ColumnNumberToName(%d) = %q, want %q", tc.num, got, tc.want)
		}
	}
}

func TestCellNameToCoordinates(t *testing.T) {
	cases := []struct {
		cell    string
		col     int
		row     int
		wantErr bool
	}{
		{"A1", 1, 1, false},
		{"Z100", 26, 100, false},
		{"AA1", 27, 1, false},
		{"a1", 1, 1, false},
		{"XFD1048576", 16384, 1048576, false},
		{"", 0, 0, true},
		{"A0", 0, 0, true},
		{"1A", 0, 0, true},
	}
	for _, tc := range cases {
		col, row, err := CellNameToCoordinates(tc.cell)
		if tc.wantErr {
			if err == nil {
				t.Errorf("CellNameToCoordinates(%q): expected error, got nil", tc.cell)
			}
			continue
		}
		if err != nil {
			t.Errorf("CellNameToCoordinates(%q): unexpected error: %v", tc.cell, err)
			continue
		}
		if col != tc.col || row != tc.row {
			t.Errorf("CellNameToCoordinates(%q) = (%d,%d), want (%d,%d)", tc.cell, col, row, tc.col, tc.row)
		}
	}
}

func TestCoordinatesToCellName(t *testing.T) {
	cases := []struct {
		col     int
		row     int
		want    string
		wantErr bool
	}{
		{1, 1, "A1", false},
		{26, 100, "Z100", false},
		{27, 1, "AA1", false},
		{0, 1, "", true},
		{1, 0, "", true},
	}
	for _, tc := range cases {
		got, err := CoordinatesToCellName(tc.col, tc.row)
		if tc.wantErr {
			if err == nil {
				t.Errorf("CoordinatesToCellName(%d,%d): expected error, got nil", tc.col, tc.row)
			}
			continue
		}
		if err != nil {
			t.Errorf("CoordinatesToCellName(%d,%d): unexpected error: %v", tc.col, tc.row, err)
			continue
		}
		if got != tc.want {
			t.Errorf("CoordinatesToCellName(%d,%d) = %q, want %q", tc.col, tc.row, got, tc.want)
		}
	}
}
