package virtualmachine

import (
	"fmt"
	"testing"
)

// -- Reasonable values for testing ---------------------------------------------------------------------------------------------

const MEMORY_SIZE = 256
const STACK_SIZE = 64

// -- Support functions ---------------------------------------------------------------------------------------------------------

func checkStack(vm *VirtualMachine, expectedValue []byte) bool {
	// Stack in error state
	if vm.stack.Overflow() || vm.stack.Underflow() {
		return false
	}

	// Compare stack pointer
	if vm.stack.stackPointer != len(expectedValue) {
		return false
	}

	// Compare first bytes
	for i, v := range expectedValue {
		if vm.stack.stack[i] != v {
			return false
		}
	}

	return true
}

func checkMemory(vm *VirtualMachine, expectedValue []byte) bool {
	// compare the expected value
	for i, v := range expectedValue {
		if vm.memory.memory[i] != v {
			return false
		}
	}

	return true
}

func runProgram(program []byte, expectedStack []byte, expectedMemory []byte) error {
	vm := NewVirtualMachine(MEMORY_SIZE, STACK_SIZE)

	err := vm.Load(program)
	if err != nil {
		return err
	}

	err = vm.Run()
	if err != nil {
		return err
	}

	if expectedStack != nil {
		if !checkStack(vm, expectedStack) {
			return fmt.Errorf("Stack mismatch")
		}
	}

	if expectedMemory != nil {
		if !checkMemory(vm, expectedMemory) {
			return fmt.Errorf("Memory mismatch")
		}
	}

	return nil
}

// -- Tests ---------------------------------------------------------------------------------------------------------------------

func TestPushByte(t *testing.T) {

	program := [...]byte{
		0x08, // PushByte
		0x0C, // byte 12
		0x00} // End

	stack := [...]byte{
		0x0C}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPushInt(t *testing.T) {
	program := [...]byte{
		0x09,                                           // PushInt
		0xFE, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xEF, // random int
		0x00} // End

	stack := [...]byte{
		0xFE, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xEF}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetByte(t *testing.T) {
	program := [...]byte{
		0x10,                                           // GetByte
		0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // address
		0x00,
		0x91} // Value

	stack := [...]byte{
		0x91}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetInt(t *testing.T) {
	program := [...]byte{
		0x11,                                           // GetByte
		0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // address
		0x00,                                           // End Program
		0x81, 0x00, 0x00, 0x01, 0x80, 0x00, 0x00, 0x81} // Value

	stack := [...]byte{
		0x81, 0x00, 0x00, 0x01, 0x80, 0x00, 0x00, 0x81}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutByte(t *testing.T) {
	program := [...]byte{
		0x08,                                           // PushByte
		0xFE,                                           // Operant
		0x18,                                           // Put Byte
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Operant
		0x00, // End Program
		0x00}

	stack := [...]byte{}

	memory := [...]byte{
		0x08,                                           // see program
		0xFE,                                           // ..
		0x18,                                           // ..
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ..
		0x00, // ..
		0xFE} // Changed!

	err := runProgram(program[:], stack[:], memory[:])
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutInt(t *testing.T) {
	program := [...]byte{
		0x09,                                           // PushInt
		0xFE, 0x00, 0x00, 0x0F, 0xF0, 0x00, 0x00, 0xEF, // Operant
		0x19,                                           // Put Int
		0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Operant
		0x00, // End Program
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	stack := [...]byte{}

	memory := [...]byte{
		0x09,                                           // see program
		0xFE, 0x00, 0x00, 0x0F, 0xF0, 0x00, 0x00, 0xEF, // ..
		0x19,                                           // ..
		0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ..
		0x00, // End Program
		0xFE, 0x00, 0x00, 0x0F, 0xF0, 0x00, 0x00, 0xEF}

	err := runProgram(program[:], stack[:], memory[:])
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddByte(t *testing.T) {
	program := [...]byte{
		0x08, // PushByte
		0x04, // Value
		0x08, // PushByte
		0x06, // Value
		0x20, // AddByte
		0x00} // EndProgram

	stack := [...]byte{
		0x0A}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddInt(t *testing.T) {
	program := [...]byte{
		0x09,                                           // PushInt
		0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Value
		0x09,                                           // PushInt
		0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Value
		0x21, // AddInt
		0x00} // EndProgram

	stack := [...]byte{
		0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
