package cpu

import (
	"github.com/DainSlash/RISC-V-GOLANG-EMULATOR/memory"
)

// ======================================
// R-TYPE
// ======================================

func execADD(cpu *CPU, inst Instruction) {
	a := int32(cpu.Registers[inst.Rs1])
	b := int32(cpu.Registers[inst.Rs2])
	result := RegisterValue(a + b)
	cpu.writeReg(inst.Rd, result)
}

func execSUB(cpu *CPU, inst Instruction) {
	a := int32(cpu.Registers[inst.Rs1])
	b := int32(cpu.Registers[inst.Rs2])
	result := RegisterValue(a - b)

	cpu.writeReg(inst.Rd, result)
}

func execSLL(cpu *CPU, inst Instruction) {
	shamt := uint32(cpu.Registers[inst.Rs2]) & 0x1F // 0b11111
	value := uint32(cpu.Registers[inst.Rs1])
	result := RegisterValue(value << shamt)

	cpu.writeReg(inst.Rd, result)
}

func execSLT(cpu *CPU, inst Instruction) {
	a := int32(cpu.Registers[inst.Rs1])
	b := int32(cpu.Registers[inst.Rs2])
	result := RegisterValue(0)

	if a < b {
		result = RegisterValue(1)
	}

	cpu.writeReg(inst.Rd, result)
}

func execSLTU(cpu *CPU, inst Instruction) {
	a := uint32(cpu.Registers[inst.Rs1])
	b := uint32(cpu.Registers[inst.Rs2])
	result := RegisterValue(0)

	if a < b {
		result = RegisterValue(1)
	}

	cpu.writeReg(inst.Rd, result)
}

func execXOR(cpu *CPU, inst Instruction) {
	a := int32(cpu.Registers[inst.Rs1])
	b := int32(cpu.Registers[inst.Rs2])
	result := RegisterValue(a ^ b)

	cpu.writeReg(inst.Rd, result)
}

func execSRL(cpu *CPU, inst Instruction) {
	shamt := uint32(cpu.Registers[inst.Rs2]) & 0x1F
	value := uint32(cpu.Registers[inst.Rs1])
	result := RegisterValue(value >> shamt)

	cpu.writeReg(inst.Rd, result)
}

func execSRA(cpu *CPU, inst Instruction) {
	shamt := uint32(cpu.Registers[inst.Rs2]) & 0x1F
	value := int32(cpu.Registers[inst.Rs1])
	result := RegisterValue(value >> shamt)

	cpu.writeReg(inst.Rd, result)
}

func execOR(cpu *CPU, inst Instruction) {
	a := int32(cpu.Registers[inst.Rs1])
	b := int32(cpu.Registers[inst.Rs2])
	result := RegisterValue(a | b)

	cpu.writeReg(inst.Rd, result)
}

func execAND(cpu *CPU, inst Instruction) {
	a := int32(cpu.Registers[inst.Rs1])
	b := int32(cpu.Registers[inst.Rs2])
	result := RegisterValue(a & b)

	cpu.writeReg(inst.Rd, result)
}

// ======================================
// I-TYPE
// ======================================

// ADDI  rd = rs1 + imm
func execADDI(cpu *CPU, inst Instruction) {
	a := int32(cpu.Registers[inst.Rs1])
	imm := int32(inst.Imm)
	result := RegisterValue(a + imm)

	cpu.writeReg(inst.Rd, result)
}

func execSLTI(cpu *CPU, inst Instruction) {
	a := int32(cpu.Registers[inst.Rs1])
	imm := int32(inst.Imm)
	result := RegisterValue(0)

	if a < imm {
		result = RegisterValue(1)
	}

	cpu.writeReg(inst.Rd, result)
}

func execSLTIU(cpu *CPU, inst Instruction) {
	a := uint32(cpu.Registers[inst.Rs1])
	imm := uint32(inst.Imm)
	result := RegisterValue(0)

	if a < imm {
		result = RegisterValue(1)
	}

	cpu.writeReg(inst.Rd, result)
}

func execXORI(cpu *CPU, inst Instruction) {
	a := int32(cpu.Registers[inst.Rs1])
	imm := int32(inst.Imm)
	result := RegisterValue(a ^ imm)

	cpu.writeReg(inst.Rd, result)
}

func execORI(cpu *CPU, inst Instruction) {
	a := RegisterValue(cpu.Registers[inst.Rs1])
	imm := RegisterValue(inst.Imm)

	cpu.writeReg(inst.Rd, a|imm)
}

func execANDI(cpu *CPU, inst Instruction) {
	a := RegisterValue(cpu.Registers[inst.Rs1])
	imm := RegisterValue(inst.Imm)

	cpu.writeReg(inst.Rd, a&imm)
}

func execSLLI(cpu *CPU, inst Instruction) {
	shamt := uint32(inst.Imm) & 0x1F
	val := uint32(cpu.Registers[inst.Rs1])

	cpu.writeReg(inst.Rd, RegisterValue(val<<shamt))
}

func execShiftRightImm(cpu *CPU, inst Instruction) {
	imm := uint32(inst.Imm)
	shamt := imm & 0x1F
	top7 := (imm >> 5) & 0x7F

	switch top7 {
	case 0x00:
		val := uint32(cpu.Registers[inst.Rs1])
		cpu.writeReg(inst.Rd, RegisterValue(val>>shamt))
	case 0x20:
		val := int32(cpu.Registers[inst.Rs1])
		cpu.writeReg(inst.Rd, RegisterValue(val>>shamt))
	default:
		// valores reservados/ilegais
	}
}

// ======================================
// LOAD INSTRUCTIONS (I-Type)
// ======================================

func execLB(cpu *CPU, inst Instruction) {
	addr := uint32(int32(cpu.Registers[inst.Rs1]) + inst.Imm)
	val := int8(cpu.rdInterface.ReadByte(addr))

	cpu.writeReg(inst.Rd, RegisterValue(int32(val)))
}

func execLH(cpu *CPU, inst Instruction) {
	addr := uint32(int32(cpu.Registers[inst.Rs1]) + inst.Imm)
	val := int16(cpu.rdInterface.ReadHalf(addr))

	cpu.writeReg(inst.Rd, RegisterValue(int32(val)))
}

func execLW(cpu *CPU, inst Instruction) {
	addr := uint32(int32(cpu.Registers[inst.Rs1]) + inst.Imm)
	val := int32(cpu.rdInterface.ReadWord(addr))

	cpu.writeReg(inst.Rd, RegisterValue(val))
}

func execLBU(cpu *CPU, inst Instruction) {
	addr := uint32(int32(cpu.Registers[inst.Rs1]) + inst.Imm)
	val := uint32(cpu.rdInterface.ReadByte(addr))

	cpu.writeReg(inst.Rd, RegisterValue(val))
}

func execLHU(cpu *CPU, inst Instruction) {
	addr := uint32(int32(cpu.Registers[inst.Rs1]) + inst.Imm)
	val := uint32(cpu.rdInterface.ReadHalf(addr))

	cpu.writeReg(inst.Rd, RegisterValue(val))
}

// ======================================
// STORE INSTRUCTIONS (S-Type)
// ======================================

func execSB(cpu *CPU, inst Instruction) {
	addr := uint32(int32(cpu.Registers[inst.Rs1]) + inst.Imm)
	val := memory.Byte(cpu.Registers[inst.Rs2] & 0xFF)

	cpu.rdInterface.WriteByte(addr, val)
}

func execSH(cpu *CPU, inst Instruction) {
	addr := uint32(int32(cpu.Registers[inst.Rs1]) + inst.Imm)
	val := uint16(cpu.Registers[inst.Rs2] & 0xFFFF)

	cpu.rdInterface.WriteHalf(addr, val)
}

func execSW(cpu *CPU, inst Instruction) {
	addr := uint32(int32(cpu.Registers[inst.Rs1]) + inst.Imm)
	val := uint32(cpu.Registers[inst.Rs2])

	cpu.rdInterface.WriteWord(addr, val)
}

// ======================================
// BRANCH INSTRUCTIONS (B-Type)
// ======================================

func execBEQ(cpu *CPU, inst Instruction) {
	branch(cpu, inst, cpu.Registers[inst.Rs1] == cpu.Registers[inst.Rs2])
}

func execBNE(cpu *CPU, inst Instruction) {
	branch(cpu, inst, cpu.Registers[inst.Rs1] != cpu.Registers[inst.Rs2])
}

func execBLT(cpu *CPU, inst Instruction) {
	branch(cpu, inst, int32(cpu.Registers[inst.Rs1]) < int32(cpu.Registers[inst.Rs2]))
}

func execBGE(cpu *CPU, inst Instruction) {
	branch(cpu, inst, int32(cpu.Registers[inst.Rs1]) >= int32(cpu.Registers[inst.Rs2]))
}

func execBLTU(cpu *CPU, inst Instruction) {
	branch(cpu, inst, uint32(cpu.Registers[inst.Rs1]) < uint32(cpu.Registers[inst.Rs2]))
}

func execBGEU(cpu *CPU, inst Instruction) {
	branch(cpu, inst, uint32(cpu.Registers[inst.Rs1]) >= uint32(cpu.Registers[inst.Rs2]))
}

func branch(cpu *CPU, inst Instruction, take bool) {
	if take {
		cpu.AddressAdder = int32(inst.Imm)
	}
}

// ======================================
// U-TYPE & J-TYPE
// ======================================

func execLUI(cpu *CPU, inst Instruction) {
	cpu.writeReg(inst.Rd, RegisterValue(inst.Imm))
}

func execAUIPC(cpu *CPU, inst Instruction) {
	val := int32(cpu.PC) + inst.Imm

	cpu.writeReg(inst.Rd, RegisterValue(val))
}

func execJAL(cpu *CPU, inst Instruction) {
	nextInstrPC := int32(cpu.PC) + 4
	cpu.writeReg(inst.Rd, RegisterValue(nextInstrPC))

	cpu.AddressAdder = int32(inst.Imm)
}

func execJALR(cpu *CPU, inst Instruction) {
	nextInstrPC := int32(cpu.PC) + 4
	cpu.writeReg(inst.Rd, RegisterValue(nextInstrPC))

	target := (int32(cpu.Registers[inst.Rs1]) + inst.Imm) & ^1

	cpu.AddressAdder = target - int32(cpu.PC)
}
