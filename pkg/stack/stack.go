package stack

import (
	"github.com/nitwhiz/gameboy/pkg/types"
)

type Stack struct {
	cpu types.CPU
	mmu types.MMU
}

func NewStack(cpu types.CPU, mmu types.MMU) *Stack {
	return &Stack{
		cpu: cpu,
		mmu: mmu,
	}
}

func (s *Stack) Push(v uint16) {
	sp := s.cpu.SP().Val()

	s.mmu.Write(sp-1, byte((v&0xFF00)>>8))
	s.mmu.Write(sp-2, byte(v&0xFF))

	s.cpu.SP().Set(sp - 2)
}

func (s *Stack) Pop() uint16 {
	sp := s.cpu.SP().Val()

	v := uint16(s.mmu.Read(sp)) | (uint16(s.mmu.Read(sp+1)) << 8)

	s.cpu.SP().Set(sp + 2)

	return v
}
