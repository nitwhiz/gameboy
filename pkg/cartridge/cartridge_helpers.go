package cartridge

type Memory []byte

func (m *Memory) Read(off int) byte {
	if off < 0 || off > len(*m)-1 {
		return 0xFF
	}

	return (*m)[off]
}

func (m *Memory) Write(off int, v byte) {
	if off < 0 || off > len(*m)-1 {
		return
	}

	(*m)[off] = v
}
