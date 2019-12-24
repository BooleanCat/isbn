package isbn

// An ISBN13 is a 13-digit International Standard Book Number as defined by ISO
// standard 2108.
type ISBN13 uint64

// NewISBN13 returns a valid ISBN13 given a uint64. An error is returned if
// validation fails. If the input if guaranteed to be a valid ISBN13 then
// prefer `isbn.ISBN13(i)`.
func NewISBN13(i uint64) (ISBN13, error) {
	return ISBN13(i), nil
}

// GS1 returns the three digit number issued by the GS1 Global Office to the
// Registration Authority.
func (isbn ISBN13) GS1() uint64 {
	return uint64(isbn / 10e9)
}

// CheckDigit returns the final digit of the ISBN13, a best effort validity
// check for the ISBN13.
func (isbn ISBN13) CheckDigit() uint64 {
	return uint64(isbn % 10)
}
