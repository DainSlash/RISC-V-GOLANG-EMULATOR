package main

import (
	"fmt"

	"github.com/DainSlash/RISC-V-GOLANG-EMULATOR/memory"
	"github.com/DainSlash/RISC-V-GOLANG-EMULATOR/motherboard"
)

func main() {

	fmt.Println("RISC-V Emulator iniciado!")

	programas := []memory.Byte{
		0b10110011,
		0b10000001,
		0b00100000,
		0b00000000,
		// soma R1 + R2 -> R3
	}

	mainboard := motherboard.NewMotherboard(motherboard.DefaultRAMSize, memory.BootProgram())

	for i := 0; i < len(programas); i++ {
		b := mainboard.ROM.ReadByte(uint32(i))
		fmt.Printf("ROM[%d] = %08b (0x%02X)\n", i, b, b)
	}

	mainboard.IntialBOOT()
	mainboard.CPU.Registers[1] = 10
	mainboard.CPU.Registers[2] = 3
	mainboard.CPU.Registers[3] = 6
	mainboard.CPU.Step()

}
