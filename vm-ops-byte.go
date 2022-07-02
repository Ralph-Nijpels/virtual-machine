package virtualmachine

import "unsafe"

// operationPushByte takes the following bytes and pushes it as a byte
func (vm *VirtualMachine) operationPushByte() (err error) {
	operant, err := vm.memory.GetByte(vm.programPointer + 1)
	if err != nil {
		return err
	}

	err = vm.stack.PushByte(operant)
	if err != nil {
		return err
	}

	vm.programPointer += 2

	vm.addLog("push-byte %d", operant)
	return nil
}

// operationGetByte pops an address from the stack and pushes the byte from that address in memory
func (vm *VirtualMachine) operationGetByte() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	value, err := vm.memory.GetByte(address)
	if err != nil {
		return err
	}

	err = vm.stack.PushByte(value)
	if err != nil {
		return err
	}

	vm.programPointer += 1

	vm.addLog("get-byte")
	return nil
}

// operationGetByteAddress takes an address and pushes the byte from that memory-address
func (vm *VirtualMachine) operationGetByteAddress() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.memory.GetByte(operant)
	if err != nil {
		return err
	}

	err = vm.stack.PushByte(value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("get-byte (%d)", operant)
	return nil
}

// operationGetByteStack takes an address and pushes the byte from the memory-address (sp + operant)
func (vm *VirtualMachine) operationGetByteStack() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.stack.GetByte(operant)
	if err != nil {
		return err
	}

	err = vm.stack.PushByte(value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("get-byte {%d}", operant)
	return nil
}

// operationPutByte pops an address and pops a byte into that memory-address
func (vm *VirtualMachine) operationPutByte() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	value, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	err = vm.memory.PutByte(address, value)
	if err != nil {
		return err
	}

	vm.programPointer += 1

	vm.addLog("put-byte")
	return nil
}

// operationPutByteAddress takes an address and pops a byte into that memory-address
func (vm *VirtualMachine) operationPutByteAddress() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	err = vm.memory.PutByte(operant, value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("put-byte (%d)", operant)
	return nil
}

// operationPutByteStack takes an address and pops a byte to the memory-address (sp + operant)
func (vm *VirtualMachine) operationPutByteStack() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	err = vm.stack.PutByte(operant, value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("put-byte {%d}", operant)
	return nil
}

// operationAddByte takes 2 bytes from the stack, adds them pushes the result
func (vm *VirtualMachine) operationAddByte() (err error) {
	operant1, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	err = vm.stack.PushByte(operant1 + operant2)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("add-byte")
	return nil
}

// operationSubByte takes 2 bytes from the stack, subtracts them and pushes the result
func (vm *VirtualMachine) operationSubByte() (err error) {
	operant1, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	err = vm.stack.PushByte(operant2 - operant1)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("sub-byte")
	return nil
}

// operationMulByte takes 2 bytes from the stack, multiplies them and pushes the result
func (vm *VirtualMachine) operationMulByte() (err error) {
	operant1, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	err = vm.stack.PushByte(operant2 * operant1)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("mul-byte")
	return nil
}

// operationDivByte takes 2 bytes from the stack, divides them and pushes the result
func (vm *VirtualMachine) operationDivByte() (err error) {
	operant1, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	err = vm.stack.PushByte(operant2 / operant1)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("div-byte")
	return nil
}

// operationEqualByte takes 2 bytes from the stack, pushes -1 if equal, 0 if not
func (vm *VirtualMachine) operationEqualByte() (err error) {
	operant1, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopByte()
	if err != nil {
		return err
	}

	result := byte(0x00)
	if operant1 == operant2 {
		result = byte(0xFF)
	}

	err = vm.stack.PushByte(result)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("equal-byte")
	return nil
}
