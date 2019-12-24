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

func TestISBN13_GS1(t *testing.T) {
	assert.Equal(t, isbn.ISBN13(9780575094185).GS1(), uint64(978))
}

func TestISBN13_CheckDigit(t *testing.T) {
	assert.Equal(t, isbn.ISBN13(9780575094185).CheckDigit(), uint64(5))
}
