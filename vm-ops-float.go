package virtualmachine

import "unsafe"

// operationPushFloat takes the following 8 bytes and pushes them on the stack as a float
func (vm *VirtualMachine) operationPushFloat() (err error) {
	operant, err := vm.memory.GetFloat(vm.programPointer + 1)
	if err != nil {
		return err
	}

	err = vm.stack.PushFloat(operant)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("PushFloat %f", operant)
	return nil
}

// operationGetFloat takes an address and pushes the float from that memory-address
func (vm *VirtualMachine) operationGetFloat() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.memory.GetFloat(operant)
	if err != nil {
		return err
	}

	err = vm.stack.PushFloat(value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("GetFloat (%d) -> %f", operant, value)
	return nil
}

// operationPutFloat takes an address and pops a float into that memory-address
func (vm *VirtualMachine) operationPutFloat() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	err = vm.memory.PutFloat(operant, value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("PutFloat %f -> (%d)", value, operant)
	return nil
}

// operationAddFloat takes 2 floats from the stack, adds them and pushes the result
func (vm *VirtualMachine) operationAddFloat() error {
	operant1, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	err = vm.stack.PushFloat(operant1 + operant2)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("AddFloat %f, %f: OK", operant1, operant2)
	return nil
}

// operationSubFloat takes 2 floats from the stack, subtracts them and pushes the result
func (vm *VirtualMachine) operationSubFloat() (err error) {
	operant1, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	err = vm.stack.PushFloat(operant2 - operant1)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("SubFloat %f, %f --> %f", operant2, operant1, operant2-operant1)
	return nil
}
