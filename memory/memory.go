package memory

type Byte uint8

type Device interface {
    ReadByte(offset uint32) Byte
    WriteByte(offset uint32, data Byte)
    Size() uint32
}
