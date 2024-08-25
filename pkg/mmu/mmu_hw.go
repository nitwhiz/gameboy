package mmu

import "github.com/nitwhiz/gameboy/pkg/addr"

var noValue = struct{}{}

var unmappedIO = map[uint16]struct{}{
	0xFF03: noValue,
	0xFF08: noValue,
	0xFF09: noValue,
	0xFF0A: noValue,
	0xFF0B: noValue,
	0xFF0C: noValue,
	0xFF0D: noValue,
	0xFF0E: noValue,
	0xFF15: noValue,
	0xFF1F: noValue,
	0xFF27: noValue,
	0xFF28: noValue,
	0xFF29: noValue,
}

func isUnmappedIO(address uint16) bool {
	_, ok := unmappedIO[address]
	return ok || (address >= 0xFF4C && address <= 0xFF7F)
}

var unusedBitsIO = map[uint16]byte{
	addr.JOYP: 0b11000000,
	addr.SC:   0b01111110,
	addr.TAC:  0b11111000,
	addr.IF:   0b11100000,
	0xFF10:    0b10000000,
	0xFF1A:    0b01111111,
	0xFF1C:    0b10011111,
	0xFF20:    0b11000000,
	0xFF23:    0b00111111,
	0xFF26:    0b01110000,
	addr.STAT: 0b10000000,
}
