package isbn

import "errors"

// An ISBN13 is a 13-digit International Standard Book Number as defined by ISO
// standard 2108.
type ISBN13 uint64

// NewISBN13 returns a valid ISBN13 given a uint64. An error is returned if
// validation fails. If the input if guaranteed to be a valid ISBN13 then
// prefer `isbn.ISBN13(i)`.
func NewISBN13(i uint64) (ISBN13, error) {
	if i/10e11 > 9 || i/10e11 == 0 {
		return 0, errors.New("not a 13-digit number")
	}

	if calculateCheckDigit(i) != (i % 10) {
		return 0, errors.New("incorrect check digit")
	}

	isbn := ISBN13(i)
	if _, err := isbn.registrationGroup(); err != nil {
		return 0, err
	}

	return isbn, nil
}

// GS1 returns the three digit number issued by the GS1 Global Office to the
// Registration Authority.
func (isbn ISBN13) GS1() uint64 {
	return uint64(isbn / 10e9)
}

// RegistrationGroup identifies to which national, geographic, language or
// otherwise the ISBN belongs to.
func (isbn ISBN13) RegistrationGroup() uint64 {
	group, _ := isbn.registrationGroup()
	return group
}

// CheckDigit returns the final digit of the ISBN13, a best effort validity
// check for the ISBN13.
func (isbn ISBN13) CheckDigit() uint64 {
	return uint64(isbn % 10)
}

func calculateCheckDigit(i uint64) uint64 {
	// Drop existing check digit
	i /= 10

	var sum uint64
	for j := 0; j < 12; j++ {
		digit := i % 10
		i /= 10

		if j%2 == 0 {
			sum += digit * 3
		} else {
			sum += digit
		}
	}

	return sum % 10
}

func (isbn ISBN13) registrationGroup() (uint64, error) {
	if isbn.GS1() == 978 {
		return isbn.registrationGroupGS1978()
	}

	return isbn.registrationGroupGS1979()
}
func (isbn ISBN13) registrationGroupGS1978() (uint64, error) {
	for _, n := range registrationGroupBound978 {
		digits := uint64(isbn) / n.divisor % n.mod
		for _, b := range n.bounds {
			if uint32(digits) >= b.lower && uint32(digits) <= b.upper {
				return digits, nil
			}
		}
	}

	return 0, errors.New("registration group not defined")
}

func (isbn ISBN13) registrationGroupGS1979() (uint64, error) {
	for _, n := range registrationGroupBound979 {
		digits := uint64(isbn) / n.divisor % n.mod
		for _, b := range n.bounds {
			if uint32(digits) >= b.lower && uint32(digits) <= b.upper {
				return digits, nil
			}
		}
	}

	return 0, errors.New("registration group not defined")
}

type bound struct {
	lower uint32
	upper uint32
}

type bounds struct {
	bounds  []bound
	divisor uint64
	mod     uint64
}

var registrationGroupBound978 = map[uint8]bounds{
	1: {
		divisor: 10e8,
		mod:     10,
		bounds: []bound{
			{lower: 0, upper: 5},
			{lower: 7, upper: 7},
		},
	},
	2: {
		divisor: 10e7,
		mod:     100,
		bounds: []bound{
			{lower: 80, upper: 94},
			{lower: 65, upper: 65},
		},
	},
	3: {
		divisor: 10e6,
		mod:     1000,
		bounds: []bound{
			{lower: 600, upper: 649},
			{lower: 950, upper: 989},
		},
	},
	4: {
		divisor: 10e5,
		mod:     10000,
		bounds: []bound{
			{lower: 9900, upper: 9989},
		},
	},
	5: {
		divisor: 10e4,
		mod:     100000,
		bounds: []bound{
			{lower: 99900, upper: 99999},
		},
	},
}

var registrationGroupBound979 = map[uint8]bounds{
	1: {
		divisor: 10e8,
		mod:     10,
		bounds: []bound{
			{lower: 8, upper: 8},
		},
	},
	2: {
		divisor: 10e7,
		mod:     100,
		bounds: []bound{
			{lower: 10, upper: 12},
		},
	},
}
