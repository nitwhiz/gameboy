package stack

import (
	"github.com/nitwhiz/gameboy/pkg/cpu"
	"github.com/nitwhiz/gameboy/pkg/mmu"
)

type Stack struct {
	CPU *cpu.CPU
	MMU *mmu.MMU
}

func (s *Stack) Push(v uint16) {
	sp := s.CPU.SP.Val()

	s.MMU.Write(sp-1, byte((v&0xFF00)>>8))
	s.MMU.Write(sp-2, byte(v&0xFF))

	s.CPU.SP.Set(sp - 2)
}

func (s *Stack) Pop() uint16 {
	sp := s.CPU.SP.Val()

	v := uint16(s.MMU.Read(sp)) | (uint16(s.MMU.Read(sp+1)) << 8)

	s.CPU.SP.Set(sp + 2)

	return v
}
