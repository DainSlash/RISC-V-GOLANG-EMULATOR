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
