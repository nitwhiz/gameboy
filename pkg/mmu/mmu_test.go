package mmu

import (
	"fmt"
	"github.com/nitwhiz/gameboy/pkg/addr"
	"github.com/nitwhiz/gameboy/pkg/cartridge"
	"github.com/nitwhiz/gameboy/pkg/memory"
	"reflect"
	"testing"
	"time"
)

func testWrite(t *testing.T, m *memory.Memory, start uint16, end uint16, target any, writeOffset int) {
	tt := reflect.ValueOf(target)

	if tt.Kind() != reflect.Ptr {
		t.Fatal("target must be pointer")
	}

	tt = tt.Elem()

	if tt.Kind() != reflect.Array {
		t.Fatal("target must be pointer to array")
	}

	mman := MMU{
		Cartridge:      nil,
		Memory:         m,
		SerialReceiver: nil,
	}

	mman.Write(start, 0xDE)
	mman.Write(start+1, 0xAD)
	mman.Write(start+2, 0xCA)
	mman.Write(start+3, 0xFE)
	mman.Write(start+5, 0xBE)
	mman.Write(start+6, 0xEF)

	res1 := ""

	for i := 0; i < 7; i++ {
		res1 += fmt.Sprintf("%02X", tt.Index(i+writeOffset))
	}

	if res1 != "DEADCAFE00BEEF" {
		t.Fatalf("expected 'DEADCAFE00BEEF', got '%s'", res1)
	}

	mman.Write(end-7, 0xDE)
	mman.Write(end-6, 0xAD)
	mman.Write(end-5, 0xCA)
	mman.Write(end-4, 0xFE)
	mman.Write(end-2, 0xBE)
	mman.Write(end-1, 0xEF)

	res2 := ""

	for i := 0; i < 7; i++ {
		res2 += fmt.Sprintf("%02X", tt.Index(int(end-start)-7+i+writeOffset))
	}

	if res2 != "DEADCAFE00BEEF" {
		t.Fatalf("expected 'DEADCAFE00BEEF', got '%s'", res2)
	}
}

func testRead(t *testing.T, m *memory.Memory, start uint16, end uint16, target any, writeOffset int) {
	tt := reflect.ValueOf(target)

	if tt.Kind() != reflect.Ptr {
		t.Fatal("target must be pointer")
	}

	tt = tt.Elem()

	if tt.Kind() != reflect.Array {
		t.Fatal("target must be pointer to array")
	}

	mman := MMU{
		Cartridge:      nil,
		Memory:         m,
		SerialReceiver: nil,
	}

	tt.Index(writeOffset).Set(reflect.ValueOf(byte(0xDE)))
	tt.Index(writeOffset + 1).Set(reflect.ValueOf(byte(0xAD)))
	tt.Index(writeOffset + 2).Set(reflect.ValueOf(byte(0xCA)))
	tt.Index(writeOffset + 3).Set(reflect.ValueOf(byte(0xFE)))
	tt.Index(writeOffset + 5).Set(reflect.ValueOf(byte(0xBE)))
	tt.Index(writeOffset + 6).Set(reflect.ValueOf(byte(0xEF)))

	res1 := ""

	for i := uint16(0); i < 7; i++ {
		res1 += fmt.Sprintf("%02X", mman.Read(start+i))
	}

	if res1 != "DEADCAFE00BEEF" {
		t.Fatalf("expected 'DEADCAFE00BEEF', got '%s'", res1)
	}

	tt.Index(writeOffset + int(end-start) - 7).Set(reflect.ValueOf(byte(0xDE)))
	tt.Index(writeOffset + int(end-start) - 6).Set(reflect.ValueOf(byte(0xAD)))
	tt.Index(writeOffset + int(end-start) - 5).Set(reflect.ValueOf(byte(0xCA)))
	tt.Index(writeOffset + int(end-start) - 4).Set(reflect.ValueOf(byte(0xFE)))
	tt.Index(writeOffset + int(end-start) - 2).Set(reflect.ValueOf(byte(0xBE)))
	tt.Index(writeOffset + int(end-start) - 1).Set(reflect.ValueOf(byte(0xEF)))

	res2 := ""

	for i := 0; i < 7; i++ {
		res2 += fmt.Sprintf("%02X", tt.Index(int(end-start)-7+i+writeOffset))
	}

	if res2 != "DEADCAFE00BEEF" {
		t.Fatalf("expected 'DEADCAFE00BEEF', got '%s'", res2)
	}
}

func TestMMU_VideoRAM(t *testing.T) {
	start := uint16(0x8000)
	end := uint16(0x9FFF)

	m := memory.New()

	testWrite(t, m, start, end, &m.VRAM, 0)
}

func TestMMU_WRAM1(t *testing.T) {
	m := memory.New()

	start := uint16(0xC000)
	end := uint16(0xCFFF)

	testWrite(t, m, start, end, &m.WRAM, 0)

	m = memory.New()

	testRead(t, m, start, end, &m.WRAM, 0)
}

func TestMMU_WRAM2(t *testing.T) {
	m := memory.New()

	start := uint16(0xD000)
	end := uint16(0xDFFF)

	testWrite(t, m, start, end, &m.WRAM, 0x1000)

	m = memory.New()

	testRead(t, m, start, end, &m.WRAM, 0x1000)
}

func TestMMU_OAM(t *testing.T) {
	m := memory.New()

	start := uint16(0xFE00)
	end := uint16(0xFE9F)

	testWrite(t, m, start, end, &m.OAM, 0)
}

func TestMMU_HRAM(t *testing.T) {
	m := memory.New()

	start := uint16(0xFF80)
	end := uint16(0xFFFE)

	testWrite(t, m, start, end, &m.HRAM, 0)

	m = memory.New()

	testRead(t, m, start, end, &m.HRAM, 0)
}

func TestMMU_ROM(t *testing.T) {
	m := memory.New()

	romData := make([]byte, 0x8000)

	romData[addr.CartridgeRomSize] = 1
	romData[addr.CartridgeType] = byte(cartridge.TypeROM)

	romData[0x1250] = 0xDE
	romData[0x1251] = 0xAD
	romData[0x1252] = 0xCA
	romData[0x1253] = 0xFE
	romData[0x1255] = 0xBE
	romData[0x1256] = 0xEF

	cart, err := cartridge.New(romData)

	if err != nil {
		t.Fatal(err)
	}

	mman := MMU{
		Cartridge:      cart,
		Memory:         m,
		SerialReceiver: nil,
	}

	res2 := ""

	for i := 0; i < 7; i++ {
		res2 += fmt.Sprintf("%02X", mman.Read(0x1250+uint16(i)))
	}

	if res2 != "DEADCAFE00BEEF" {
		t.Fatalf("expected 'DEADCAFE00BEEF', got '%s'", res2)
	}
}

func TestSBReceiver(t *testing.T) {
	done := make(chan bool)

	m := memory.New()

	mman := MMU{
		Cartridge: nil,
		Memory:    m,
		SerialReceiver: func(b byte) {
			if b == 0xAB {
				close(done)
			}
		},
	}

	mman.Write(0xFF01, 0xAB)
	mman.Write(0xFF02, 0x81)

	select {
	case <-time.After(time.Second * 1):
		t.Fatal("expected byte 0xAB not read read in 1 second")
	case <-done:
	}
}
