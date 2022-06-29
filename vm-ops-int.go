package virtualmachine

import "unsafe"

// operationPushInt takes the following 8 bytes from memory and pushes them as an int
func (vm *VirtualMachine) operationPushInt() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	err = vm.stack.PushInt(operant)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("PushInt %d", operant)
	return nil
}

// operationGetInt pops an address from stack and pushes the int from that memory-address
func (vm *VirtualMachine) operationGetInt() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	value, err := vm.memory.GetInt(address)
	if err != nil {
		return err
	}

	err = vm.stack.PushInt(value)
	if err != nil {
		return err
	}

	vm.programPointer += 1

	vm.addLog("GetInt %d", value)
	return nil
}

// operationGetIntAddress takes an address and pushes the int from that memory-address
func (vm *VirtualMachine) operationGetIntAddress() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.memory.GetInt(operant)
	if err != nil {
		return err
	}

	err = vm.stack.PushInt(value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("GetInt (%d) -> %d", operant, value)
	return nil
}

// operationPutIntAddress takes an address and pops an int into that memory-address
func (vm *VirtualMachine) operationPutIntAddress() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	err = vm.memory.PutInt(operant, value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("PutInt %d -> (%d)", value, operant)
	return nil
}

// operationAddInt takes 2 integers from the stack, adds them and pushes the result
func (vm *VirtualMachine) operationAddInt() (err error) {
	operant1, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	err = vm.stack.PushInt(operant1 + operant2)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("AddInt %d, %d", operant1, operant2)
	return nil
}

// operationSubInt takes 2 ints from the stack, subtracts them and pushes the result
func (vm *VirtualMachine) operationSubInt() (err error) {
	operant1, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	err = vm.stack.PushInt(operant2 - operant1)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("AddByte %d, %d", operant2, operant1)
	return nil
}
