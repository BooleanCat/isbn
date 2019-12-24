package assert

import (
	"reflect"
	"testing"
)

// Nil fails a test if the interface provided is not nil.
func Nil(t *testing.T, v interface{}) {
	if v != nil {
		t.Fatalf(`expected "%v" to be nil`, v)
	}
}

// Equal fails a test if the interfaces provided are not deeply equal.
func Equal(t *testing.T, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		t.Fatalf(`expected "%v" to equal "%v"`, x, y)
	}
}
