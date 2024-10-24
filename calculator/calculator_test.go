package calculator

import (
	"reflect"
	"testing"
)

func TestGetSqaresOf(t *testing.T) {
	var got []int = GetSqaresOf([]int{1, 2, 3, 4, 5})
	want := []int{1, 4, 9, 16, 25}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Slices not equal got=%v want=%v", got, want)
	}
}
