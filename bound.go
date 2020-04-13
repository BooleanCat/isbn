package isbn

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
