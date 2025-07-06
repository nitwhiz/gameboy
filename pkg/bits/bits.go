package bits

// Set the nth bit in v.
func Set(v byte, n byte) byte {
	return v | (1 << n)
}

// Val - value of the nth bit in v.
// returns 0 or 1
func Val(v byte, n byte) byte {
	if Test(v, n) {
		return 1
	}

	return 0
}

// Reset the nth bit in v.
func Reset(v byte, n byte) byte {
	return v & ^(1 << n)
}

// Test if a bit is 1 in v.
func Test(v byte, n byte) bool {
	return (v>>n)&1 == 1
}
