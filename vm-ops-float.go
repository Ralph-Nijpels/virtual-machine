package virtualmachine

import "unsafe"

// operationPushFloat takes the following 8 bytes and pushes them on the stack as a float
func (vm *VirtualMachine) operationPushFloat() (err error) {
	operant, err := vm.memory.GetFloat(vm.programPointer + 1)
	if err != nil {
		vm.addLog("PushFloat: %v\n", err)
		return err
	}

	err = vm.stack.PushFloat(operant)
	if err != nil {
		vm.addLog("PushFloat %f: %v\n", operant, err)
		return err
	}

	vm.addLog("PushFloat %f: OK\n", operant)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationGetFloat takes an address and pushes the float from that memory-address
func (vm *VirtualMachine) operationGetFloat() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("GetFloat: %v", err)
		return err
	}

	value, err := vm.memory.GetFloat(operant)
	if err != nil {
		vm.addLog("GetFloat (%d): %v", operant, err)
		return err
	}

	err = vm.stack.PushFloat(value)
	if err != nil {
		vm.addLog("GetFloat (%d) -> %f: %v", operant, value, err)
		return err
	}

	vm.addLog("GetFloat (%d) -> %f: OK", operant, value)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationPutFloat takes an address and pops a float into that memory-address
func (vm *VirtualMachine) operationPutFloat() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("PutFloat: %v", err)
		return err
	}

	value, err := vm.stack.PopFloat()
	if err != nil {
		vm.addLog("PutFloat (%d): %v", operant, err)
		return err
	}

	err = vm.memory.PutFloat(operant, value)
	if err != nil {
		vm.addLog("PutFloat %f -> (%d): %v", value, operant, err)
		return err
	}

	vm.addLog("PutFloat %f -> (%d): OK", value, operant)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationAddFloat takes 2 floats from the stack, adds them and pushes the result
func (vm *VirtualMachine) operationAddFloat() error {
	operant1, err := vm.stack.PopFloat()
	if err != nil {
		vm.addLog("AddFloat: %v", err)
		return err
	}

	operant2, err := vm.stack.PopFloat()
	if err != nil {
		vm.addLog("AddFloat %f: %v", operant1, err)
		return err
	}

	err = vm.stack.PushFloat(operant1 + operant2)
	if err != nil {
		vm.addLog("AddFloat %f, %f: %v", operant1, operant2, err)
		return err
	}

	vm.addLog("AddFloat %f, %f: OK", operant1, operant2)
	vm.programPointer++
	return nil
}
