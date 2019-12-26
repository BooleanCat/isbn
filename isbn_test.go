package isbn_test

import (
	"testing"

	"github.com/BooleanCat/isbn"
	"github.com/BooleanCat/isbn/assert"
)

func TestNewISBN13(t *testing.T) {
	isbn13, err := isbn.NewISBN13(9780575094185)
	assert.Nil(t, err)
	assert.Equal(t, isbn13, isbn.ISBN13(9780575094185))
}

func TestNewISBN13_TooShort(t *testing.T) {
	_, err := isbn.NewISBN13(978057509418)
	assert.ErrorMatches(t, err, "not a 13-digit number")
}

func TestNewISBN13_TooLong(t *testing.T) {
	_, err := isbn.NewISBN13(97805750941856)
	assert.ErrorMatches(t, err, "not a 13-digit number")
}

func TestNewISBN13_RegistrationGroupUndefined(t *testing.T) {
	_, err := isbn.NewISBN13(9786675094184)
	assert.ErrorMatches(t, err, "registration group not defined")
}

func TestNewISBN13_IncorrectCheckDigit(t *testing.T) {
	_, err := isbn.NewISBN13(9780575094184)
	assert.ErrorMatches(t, err, "incorrect check digit")
}

func TestISBN13_GS1(t *testing.T) {
	assert.Equal(t, isbn.ISBN13(9780575094185).GS1(), uint64(978))
}

func TestISBN13_RegistrationGroup(t *testing.T) {
	tests := map[string]struct {
		isbn isbn.ISBN13
		want uint64
	}{
		"GS1 978 single digit lower bound": {isbn: 9780575094185, want: 0},
		"GS1 978 single digit upper bound": {isbn: 9785575094180, want: 5},
		"GS1 978 single digit 7":           {isbn: 9787575094184, want: 7},
		"GS1 978 two digit 65":             {isbn: 9786575094187, want: 65},
		"GS1 978 two digit lower bound":    {isbn: 9788075094186, want: 80},
		"GS1 978 two digit upper bound":    {isbn: 9789475094189, want: 94},
		"GS1 978 three digit lower bound":  {isbn: 9786005094183, want: 600},
		"GS1 978 three digit upper bound":  {isbn: 9786495094182, want: 649},
		"GS1 978 four digit lower bound":   {isbn: 9789900094180, want: 9900},
		"GS1 978 four digit upper bound":   {isbn: 9789989094187, want: 9989},
		"GS1 978 five digit lower bound":   {isbn: 9789990094183, want: 99900},
		"GS1 978 five digit upper bound":   {isbn: 9789999994187, want: 99999},

		"GS1 979 single digit 8":        {isbn: 9798495094185, want: 8},
		"GS1 979 two digit lower bound": {isbn: 9791095094185, want: 10},
		"GS1 979 two digit upper bound": {isbn: 9791295094185, want: 12},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, isbn.ISBN13(tc.isbn).RegistrationGroup(), tc.want)
		})
	}
}

func TestISBN13_CheckDigit(t *testing.T) {
	assert.Equal(t, isbn.ISBN13(9780575094185).CheckDigit(), uint64(5))
}
