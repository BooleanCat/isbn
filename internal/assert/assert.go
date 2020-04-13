package assert

import (
	"reflect"
	"regexp"
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

// ErrorMatches fails a test if the provided error's message does not match the
// regular expression provided.
func ErrorMatches(t *testing.T, err error, pattern string) {
	if err == nil {
		t.Fatal("expected err to have occurred")
	}

	if !regexp.MustCompile(pattern).MatchString(err.Error()) {
		t.Fatalf(`expected error "%v" to match pattern "%s"`, err, pattern)
	}
}
