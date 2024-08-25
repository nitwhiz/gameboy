package mmu

import "github.com/nitwhiz/gameboy/pkg/addr"

func isUnmappedIO(address uint16) bool {
	if address == 0xFF03 ||
		address == 0xFF08 ||
		address == 0xFF09 ||
		address == 0xFF0A ||
		address == 0xFF0B ||
		address == 0xFF0C ||
		address == 0xFF0D ||
		address == 0xFF0E ||
		address == 0xFF15 ||
		address == 0xFF1F ||
		address == 0xFF27 ||
		address == 0xFF28 ||
		address == 0xFF29 ||
		(address >= 0xFF4C && address <= 0xFF7F) {
		return true
	}

	return false
}

func getUnusedBitsIO(address uint16) byte {
	switch address {
	case addr.JOYP:
		return 0b11000000
	case addr.SC:
		return 0b01111110
	case addr.TAC:
		return 0b11111000
	case addr.IF:
		return 0b11100000
	case 0xFF10:
		return 0b10000000
	case 0xFF1A:
		return 0b01111111
	case 0xFF1C:
		return 0b10011111
	case 0xFF20:
		return 0b11000000
	case 0xFF23:
		return 0b00111111
	case 0xFF26:
		return 0b01110000
	case addr.STAT:
		return 0b10000000
	default:
		return 0
	}
}
