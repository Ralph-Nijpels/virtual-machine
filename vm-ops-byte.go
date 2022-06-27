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

	vm.addLog("Pushbyte %d", operant)
	return nil
}

// operationGetByte takes an address and pushes the byte from that memory-address
func (vm *VirtualMachine) operationGetByte() (err error) {
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

	vm.addLog("GetByte (%d) -> %d", operant, value)
	return nil
}

// operationPutByte takes an address and pops a byte into that memory-address
func (vm *VirtualMachine) operationPutByte() (err error) {
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

	vm.addLog("PutByte %d -> (%d)", value, operant)
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

	vm.addLog("AddByte %d, %d", operant1, operant2)
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

	vm.addLog("AddByte %d, %d", operant2, operant1)
	return nil
}
