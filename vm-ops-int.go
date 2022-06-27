package virtualmachine

import "unsafe"

// operationPushInt takes the following 8 bytes from memory and pushes them as an int
func (vm *VirtualMachine) operationPushInt() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("PushInt: %v\n", err)
		return err
	}

	err = vm.stack.PushInt(operant)
	if err != nil {
		vm.addLog("PushInt %d: %v\n", operant, err)
		return err
	}

	vm.addLog("PushInt %d: OK\n", operant)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationGetInt takes an address and pushes the int from that memory-address
func (vm *VirtualMachine) operationGetInt() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("GetInt: %v", err)
		return err
	}

	value, err := vm.memory.GetInt(operant)
	if err != nil {
		vm.addLog("GetInt (%d): %v", operant, err)
		return err
	}

	err = vm.stack.PushInt(value)
	if err != nil {
		vm.addLog("GetInt (%d) -> %d: %v", operant, value, err)
		return err
	}

	vm.addLog("GetInt (%d) -> %d: OK", operant, value)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationPutInt takes an address and pops an int into that memory-address
func (vm *VirtualMachine) operationPutInt() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("PutInt: %v", err)
		return err
	}

	value, err := vm.stack.PopInt()
	if err != nil {
		vm.addLog("PutInt (%d): %v", operant, err)
		return err
	}

	err = vm.memory.PutInt(operant, value)
	if err != nil {
		vm.addLog("PutInt %d -> (%d): %v", value, operant, err)
		return err
	}

	vm.addLog("PutInt %d -> (%d): OK", value, operant)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationAddInt takes 2 integers from the stack, adds them and pushes the result
func (vm *VirtualMachine) operationAddInt() error {
	operant1, err := vm.stack.PopInt()
	if err != nil {
		vm.addLog("AddInt: %v", err)
		return err
	}

	operant2, err := vm.stack.PopInt()
	if err != nil {
		vm.addLog("AddInt %d: %v", operant1, err)
		return err
	}

	err = vm.stack.PushInt(operant1 + operant2)
	if err != nil {
		vm.addLog("AddInt %d, %d: %v", operant1, operant2, err)
		return err
	}

	vm.addLog("AddInt %d, %d: OK", operant1, operant2)
	vm.programPointer++
	return nil
}
