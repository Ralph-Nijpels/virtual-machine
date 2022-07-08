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
