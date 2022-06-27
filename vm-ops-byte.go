package virtualmachine

import "unsafe"

// operationPushByte takes the following bytes and pushes it as a byte
func (vm *VirtualMachine) operationPushByte() error {
	operant, err := vm.memory.GetByte(vm.programPointer + 1)
	if err != nil {
		vm.addLog("Pushbyte: %v\n", err)
		return err
	}

	err = vm.stack.PushByte(operant)
	if err != nil {
		vm.addLog("Pushbyte %d: %v\n", operant, err)
		return err
	}

	vm.addLog("Pushbyte %d: OK\n", operant)
	vm.programPointer += 2
	return nil
}

// operationGetByte takes an address and pushes the byte from that memory-address
func (vm *VirtualMachine) operationGetByte() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("GetByte: %v", err)
		return err
	}

	value, err := vm.memory.GetByte(operant)
	if err != nil {
		vm.addLog("GetByte (%d): %v", operant, err)
		return err
	}

	err = vm.stack.PushByte(value)
	if err != nil {
		vm.addLog("GetByte (%d) -> %d: %v", operant, value, err)
		return err
	}

	vm.addLog("GetByte (%d) -> %d: OK", operant, value)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationPutByte takes an address and pops a byte into that memory-address
func (vm *VirtualMachine) operationPutByte() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("PutByte: %v", err)
		return err
	}

	value, err := vm.stack.PopByte()
	if err != nil {
		vm.addLog("PutByte (%d): %v", operant, err)
		return err
	}

	err = vm.memory.PutByte(operant, value)
	if err != nil {
		vm.addLog("PutByte %d -> (%d): %v", value, operant, err)
		return err
	}

	vm.addLog("PutByte %d -> (%d): OK", value, operant)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationAddByte takes 2 bytes from the stack, adds them pushes the result
func (vm *VirtualMachine) operationAddByte() error {
	operant1, err := vm.stack.PopByte()
	if err != nil {
		vm.addLog("AddByte: %v", err)
		return err
	}

	operant2, err := vm.stack.PopByte()
	if err != nil {
		vm.addLog("AddByte %d: %v", operant1, err)
		return err
	}

	err = vm.stack.PushByte(operant1 + operant2)
	if err != nil {
		vm.addLog("AddByte %d, %d: %v", operant1, operant2, err)
		return err
	}

	vm.addLog("AddByte %d, %d: OK", operant1, operant2)
	vm.programPointer++
	return nil
}
