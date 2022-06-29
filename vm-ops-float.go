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

	vm.addLog("push-float %f", operant)
	return nil
}

// operationGetFloat pops an address and pushes the float from that memory-address
func (vm *VirtualMachine) operationGetFloat() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	value, err := vm.memory.GetFloat(address)
	if err != nil {
		return err
	}

	err = vm.stack.PushFloat(value)
	if err != nil {
		return err
	}

	vm.programPointer += 1

	vm.addLog("get-float")
	return nil
}

// operationGetFloatAddress takes an address and pushes the float from that memory-address
func (vm *VirtualMachine) operationGetFloatAddress() (err error) {
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

	vm.addLog("get-float (%d)", operant)
	return nil
}

// operationPutFloatAddress takes an address and pops a float into that memory-address
func (vm *VirtualMachine) operationPutFloatAddress() (err error) {
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

	vm.addLog("put-float (%d)", operant)
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

	vm.addLog("add-float")
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

	vm.addLog("sub-float")
	return nil
}
