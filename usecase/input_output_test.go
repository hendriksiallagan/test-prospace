package usecase

import (
	"reflect"
	"testing"
)

func TestCategorizeInput(t *testing.T) {
	tables := []struct {
		input               []string
		outputRoman         []string
		outputIntergalactic []string
		outputQuestion      []string
	}{
		{
			[]string{
				"glob is I",
				"glob glob Silver is 34 Credits",
				"how much is pish tegj glob glob ?",
			},
			[]string{
				"glob is I",
			},
			[]string{
				"glob glob Silver is 34 Credits",
			},
			[]string{
				"how much is pish tegj glob glob ?",
			},
		},
	}
	for i, table := range tables {
		roman, intergalactic, question := CategorizeInput(table.input)
		if !reflect.DeepEqual(roman, table.outputRoman) {
			t.Errorf("test %d categorize for roman failed, got: %v, want: %v.", i, roman, table.outputRoman)
		}
		if !reflect.DeepEqual(intergalactic, table.outputIntergalactic) {
			t.Errorf("test %d categorize for intergalactic failed, got: %v, want: %v.", i, intergalactic, table.outputIntergalactic)
		}
		if !reflect.DeepEqual(question, table.outputQuestion) {
			t.Errorf("test %d categorize for question failed, got: %v, want: %v.", i, question, table.outputQuestion)
		}
	}
}
