package motherboard

import (
    "github.com/DainSlash/RISC-V-GOLANG-EMULATOR/bus"
    "github.com/DainSlash/RISC-V-GOLANG-EMULATOR/cpu"
    "github.com/DainSlash/RISC-V-GOLANG-EMULATOR/memory"
)

type Motherboard struct {
    CPU *cpu.CPU
    Bus *bus.Bus

    RAMs []*memory.RAM
    ROM  *memory.ROM

    // TODO ::
    // VRAM *memory.RAM
    // HDs  []*memory.Disk
    // GPU  *gpu.GPU
}

func NewMotherboard(ramSize uint32, bootImage []memory.Byte) *Motherboard {
    if ramSize == 0 {
        ramSize = DefaultRAMSize
    }

    bus             := bus.NewBus()
    mainRAM         := memory.NewRAM(ramSize)
    rom             := memory.NewROM(bootImage)
    cpu             := cpu.NewCPU(bus)
    
    result          := &Motherboard{}
    result.Bus      = bus    
    result.ROM      = rom
    result.RAMs     = []*memory.RAM{mainRAM}
    result.CPU      = cpu
    
    result.Bus.MapDevice(DefaultRAMBase, mainRAM)
    result.Bus.MapDevice(result.GetRamSize(), rom)

    return result
}

func (mb *Motherboard) GetRamSize() uint32 {
    var result uint32 = 0

    for _, ram := range mb.RAMs {
        result += ram.Size()
    }

    return result
}