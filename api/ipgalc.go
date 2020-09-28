package api

type cidrBlock struct {
	mask int
	size int
}

var cidrMap = make(map[int]cidrBlock)

// Setup runs initialization and stuff, especially setting up the CIDR table
func Setup() {
	cidrMap = map[int]cidrBlock{
		28: cidrBlock{mask: 240, size: 16},
		29: cidrBlock{mask: 248, size: 8},
		30: cidrBlock{mask: 252, size: 4},
		31: cidrBlock{mask: 253, size: 2},
		32: cidrBlock{mask: 255, size: 1},
	}
}
