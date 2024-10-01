package integration

import (
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/input"
	"github.com/nitwhiz/gameboy/pkg/memory"
	"github.com/nitwhiz/gameboy/pkg/mmu"
	"math/rand/v2"
	"testing"
)

var lastReadResult byte

func randomizedMemory() *memory.Memory {
	m := memory.New()

	for i := 0; i < len(m.VRAM); i++ {
		m.VRAM[i] = byte(rand.Uint())
	}

	for i := 0; i < len(m.WRAM); i++ {
		m.WRAM[i] = byte(rand.Uint())
	}

	for i := 0; i < len(m.OAM); i++ {
		m.OAM[i] = byte(rand.Uint())
	}

	for i := 0; i < len(m.HRAM); i++ {
		m.HRAM[i] = byte(rand.Uint())
	}

	for i := 0; i < len(m.IO); i++ {
		m.IO[i] = byte(rand.Uint())
	}

	return m
}

func getTestData() (*mmu.MMU, []byte) {
	m := &mmu.MMU{
		Cartridge:      nil,
		Memory:         randomizedMemory(),
		Input:          input.NewState(),
		TimerLock:      false,
		SerialReceiver: nil,
	}

	values := []byte{
		0x00,
		0xFF,
		0x7F,
		0x80,
	}

	for v := 0; v < 60; v++ {
		values = append(values, byte(rand.Uint()))
	}

	return m, values
}

func write(m *mmu.MMU, a uint16, values []byte) {
	for _, v := range values {
		m.Write(a, v)
	}
}

func testWrite(b *testing.B) {
	b.StopTimer()

	m, values := getTestData()

	b.StartTimer()

	for a := addr.MemVRAMBegin; a < addr.MemVRAMEnd; a++ {
		write(m, a, values)
	}

	for a := addr.MemWRAMBegin; a < addr.MemWRAMEnd; a++ {
		write(m, a, values)
	}

	for a := addr.MemOAMBegin; a < 0xFFFF; a++ {
		write(m, a, values)
	}
}

func read(m *mmu.MMU, a uint16) {
	lastReadResult = m.Read(a)
}

func testRead(b *testing.B) {
	b.StopTimer()

	m, _ := getTestData()

	b.StartTimer()

	for a := addr.MemVRAMBegin; a < addr.MemVRAMEnd; a++ {
		read(m, a)
	}

	for a := addr.MemWRAMBegin; a < addr.MemWRAMEnd; a++ {
		read(m, a)
	}

	for a := addr.MemOAMBegin; a < 0xFFFF; a++ {
		read(m, a)
	}
}

func BenchmarkMMUWrite(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testWrite(b)
		}
	})
}

func BenchmarkMMURead(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testRead(b)
		}
	})
}
