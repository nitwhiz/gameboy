package bits

var bitMasks = map[byte]byte{
	1:   0b1,
	3:   0b11,
	7:   0b111,
	15:  0b1111,
	31:  0b11111,
	63:  0b111111,
	127: 0b1111111,
	255: 0b11111111,
}

func GetMaskByValue(v int) byte {
	if v == 0 {
		return 0
	}

	if v > 255 {
		return 0b11111111
	}

	x := byte(2)

	for v > 1 {
		v /= 2
		x *= 2
	}

	m, ok := bitMasks[x-1]

	if !ok {
		return 0
	}

	return m
}

func GetMaskByWidth(n byte) byte {
	m := byte(0)

	for i := range n {
		m |= 1 << i
	}

	return m
}

// IsTACEnabled returns the status of bit 2 of addr.TAC
func IsTACEnabled(tac byte) bool {
	return Test(tac, 2)
}

// GetTACClockSelect returns bits 0, 1 of addr.TAC
func GetTACClockSelect(tac byte) byte {
	return tac & 0b11
}

// IsLCDEnabled returns bit 7 of addr.LCDC
func IsLCDEnabled(lcdc byte) bool {
	return Test(lcdc, 7)
}

// GetPPUMode returns bits 0, 1 of addr.STAT
func GetPPUMode(stat byte) byte {
	return stat & 0b11
}

func SetLYCLY(stat byte, v bool) byte {
	if v {
		return Set(stat, 2)
	}

	return Reset(stat, 2)
}

func IsLCDModeSelect(stat byte, mode byte) bool {
	return Test(stat, 1<<(mode+3))
}

func GetLYCSelect(stat byte) bool {
	return Test(stat, 6)
}

func SetPPUMode(stat byte, mode byte) byte {
	return (stat & 0b11111100) | mode
}

func IsLCDBackgroundAndWindowEnabled(lcdc byte) bool {
	return Test(lcdc, 0)
}

func IsLCDObjEnabled(lcdc byte) bool {
	return Test(lcdc, 1)
}

func IsLCDObjSize8x16(lcdc byte) bool {
	return Test(lcdc, 2)
}

func IsLCDWindowEnabled(lcdc byte) bool {
	return Test(lcdc, 5)
}

func OAMAttributes(attributes byte) (xFlip, yFlip, priority bool) {
	return Test(attributes, 5), Test(attributes, 6), Test(attributes, 7)
}

func IsJOYPSelectButtons(joyp byte) bool {
	return !Test(joyp, 5)
}

func IsJOYPSelectDPad(joyp byte) bool {
	return !Test(joyp, 4)
}
