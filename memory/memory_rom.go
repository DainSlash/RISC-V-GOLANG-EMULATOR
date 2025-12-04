package memory

type ROM struct {
	data []Byte
}

func NewROM(content []Byte) *ROM {
	buf := make([]Byte, len(content))
	copy(buf, content)
	return &ROM{data: buf}
}

func (r *ROM) Size() uint32 {
	return uint32(len(r.data))
}

func (r *ROM) ReadByte(offset uint32) Byte {
	if offset >= uint32(len(r.data)) {
		return 0
	}
	return r.data[offset]
}

func (r *ROM) WriteByte(offset uint32, _ Byte) {
	// ROM: ignora a tentativa de escrita
}

func BootProgram() []Byte {
	return []Byte{
		0b10110011,
		0b10000001,
		0b00100000,
		0b00000000,
		// soma R1 + R2 -> R3
	}
}
