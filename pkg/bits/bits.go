package bits

// Set a bit in v.
func Set(v byte, b byte) byte {
	return v | (1 << b)
}

// Val - value of the bit at `b`
// returns 0 or 1
func Val(v byte, b byte) byte {
	if Test(v, b) {
		return 1
	}

	return 0
}

// Reset a bit in v.
func Reset(v byte, b byte) byte {
	return v & ^(1 << b)
}

// Test if a bit is 1 in v.
func Test(v byte, b byte) bool {
	return (v>>b)&1 == 1
}
