package unique

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	input := []string{"a", "b", "a"}
	expected := []string{"a", "b"}
	res := Strings(input)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("UniqueStrings: expected %v got %v", expected, res)
	}
}
