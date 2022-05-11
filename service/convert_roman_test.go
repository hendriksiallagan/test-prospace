package service

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tables := []struct {
		input       string
		outputRoman int
		wantErr     bool
	}{
		{
			"XIX",
			19,
			false,
		}, {
			"L",
			50,
			false,
		}, {
			"M",
			1000,
			false,
		}, {
			"CM",
			900,
			false,
		}, {
			"XXXIX",
			39,
			false,
		}, {
			"aaa",
			0,
			true,
		}, {
			"IIII",
			0,
			true,
		}, {
			"LLLL",
			0,
			true,
		},
	}
	for i, table := range tables {
		output, err := RomanToInt(table.input)
		if err != nil && !table.wantErr {
			t.Errorf("test %d should not return err. err:%v", i, err.Error())
		}
		if output != table.outputRoman {
			t.Errorf("test %d conversion failed, got: %v, want: %v.", i, output, table.outputRoman)
		}
	}
}
