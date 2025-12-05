package memory

import (
	"fmt"
	"os"
)

type ROM struct {
	data []Byte
	Name string 
}

func NewROMFromFile(path string) (*ROM, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("falha ao ler ROM %s: %w", path, err)
	}

	memBytes := make([]Byte, len(bytes))
	for i, b := range bytes {
		memBytes[i] = Byte(b)
	}

	return &ROM{
		data: memBytes,
		Name: path,
	}, nil
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
	fmt.Printf("[WARN] Tentativa de escrita na ROM (Offset: 0x%X)\n", offset)
}