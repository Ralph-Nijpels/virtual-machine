package virtualmachine

import (
	"fmt"
	"unsafe"
)

// operationRet takes an address from the stack and jumps there
func (vm *VirtualMachine) operationRet() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	vm.programPointer = address

	vm.addLog("ret")
	return nil
}

// operationJmp takes an address operant and jumps there
func (vm *VirtualMachine) operationJmp() (err error) {
	address, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	vm.programPointer = address

	vm.addLog("jmp (%d)", address)
	return nil
}

// operationJmpzByte pops an address and a byte from stack, jumps to the address if the byte == 0
func (vm *VirtualMachine) operationJmpzByte() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	operant, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	if operant == byte(0) {
		vm.programPointer = address
	} else {
		vm.programPointer += 1
	}

	vm.addLog("jmpz-byte")
	return nil
}

// operationJmpzInt pops an address and an int from stack, jumps to the address if the int == 0
func (vm *VirtualMachine) operationJmpzInt() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	operant, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if operant == int(0) {
		vm.programPointer = address
	} else {
		vm.programPointer += 1
	}

	vm.addLog("jmpz-int")
	return nil
}

// operationJmpzFloat pops an address and a float from stack, jumps to the address if the float == 0.0
func (vm *VirtualMachine) operationJmpzFloat() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	operant, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	if operant == float64(0.0) {
		vm.programPointer = address
	} else {
		vm.programPointer += 1
	}

	vm.addLog("jmpz-float")
	return nil
}

// operationJmpzByteAddresstakes an address as opperant and pops a byte from stack, jumps to the address if the byte == 0
func (vm *VirtualMachine) operationJmpzByteAddress() (err error) {
	address, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	operant, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	if operant == byte(0) {
		vm.programPointer = address
	} else {
		vm.programPointer += (1 + (int)(unsafe.Sizeof(address)))
	}

	vm.addLog("jmpz-byte (% X)", address)
	return nil
}

// operationJmpzInt takes an address as opperant and pops an int from stack, jumps to the address if the byte == 0
func (vm *VirtualMachine) operationJmpzIntAddress() (err error) {
	address, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	operant, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if operant == int(0) {
		vm.programPointer = address
	} else {
		vm.programPointer += (1 + (int)(unsafe.Sizeof(address)))
	}

	vm.addLog("jmpz-int (% X)", address)
	return nil
}

// operationJmpzFloat takes an address as opperant and pops a float from stack, jumps to the address if the byte == 0
func (vm *VirtualMachine) operationJmpzFloatAddress() (err error) {
	address, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	operant, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	if operant == float64(0.0) {
		vm.programPointer = address
	} else {
		vm.programPointer += (1 + (int)(unsafe.Sizeof(address)))
	}

	vm.addLog("jmpz-float (% X)", address)
	return nil
}

// operationJmpnzByte pops an address and a byte from stack, jumps to the address if the byte != 0
func (vm *VirtualMachine) operationJmpnzByte() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	operant, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	if operant != byte(0) {
		vm.programPointer = address
	} else {
		vm.programPointer += 1
	}

	vm.addLog("jmpnz-byte")
	return nil
}

// operationJmpnzInt pops an address and an int from stack, jumps to the address if the int != 0
func (vm *VirtualMachine) operationJmpnzInt() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	operant, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if operant != int(0) {
		vm.programPointer = address
	} else {
		vm.programPointer += 1
	}

	vm.addLog("jmpz-int")
	return nil
}

// operationJmpnzFloat pops an address and a float from stack, jumps to the address if the float != 0.0
func (vm *VirtualMachine) operationJmpnzFloat() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	operant, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	if operant != float64(0.0) {
		vm.programPointer = address
	} else {
		vm.programPointer += 1
	}

	vm.addLog("jmpnz-float")
	return nil
}

// operationCall takes an address from the stack, pushes the current address + 1 to the stack and jumps to the address taken
func (vm *VirtualMachine) operationCall() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	err = vm.stack.PushInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	vm.programPointer = address

	vm.addLog("call")
	return nil
}

// operationCallAddress takes an address operant, pushes the current address + 1 to the stack and jumps to the address taken
func (vm *VirtualMachine) operationCallAddress() (err error) {
	address, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	if address < 0 || address >= vm.memory.Size() {
		return fmt.Errorf("illegal address")
	}

	err = vm.stack.PushInt(vm.programPointer + (int)(unsafe.Sizeof(address)) + 1)
	if err != nil {
		return err
	}

	vm.programPointer = address

	vm.addLog("call (%d)", address)
	return nil
}
