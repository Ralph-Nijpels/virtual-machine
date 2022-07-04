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

// operationPutFloat pops a memory-address and pops a float into that memory-address
func (vm *VirtualMachine) operationPutFloat() (err error) {
	address, err := vm.stack.PopInt()
	if err != nil {
		return err
	}

	value, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	err = vm.memory.PutFloat(address, value)
	if err != nil {
		return err
	}

	vm.programPointer += 1

	vm.addLog("put-float")
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

// operationGetFloatStack takes an offset operant and pushes the float from the memory-address stack-pointer + offset
func (vm *VirtualMachine) operationGetFloatStack() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.stack.GetFloat(operant)
	if err != nil {
		return err
	}

	err = vm.stack.PushFloat(value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("get-float {%d}", operant)
	return nil
}

// operationPutFloatStack takes an offset operant and pops a float into the memory-address (stack-pointer + offset)
func (vm *VirtualMachine) operationPutFloatStack() (err error) {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		return err
	}

	value, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	err = vm.stack.PutFloat(operant, value)
	if err != nil {
		return err
	}

	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))

	vm.addLog("put-float {%d}", operant)
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

// operationMulFloat takes 2 floats from the stack, multiplies them and pushes the result
func (vm *VirtualMachine) operationMulFloat() (err error) {
	operant1, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	err = vm.stack.PushFloat(operant2 * operant1)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("mul-float")
	return nil
}

// operationDivFloat takes 2 floats from the stack, divides them and pushes the result
func (vm *VirtualMachine) operationDivFloat() (err error) {
	operant1, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	err = vm.stack.PushFloat(operant2 / operant1)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("div-float")
	return nil
}

// operationEqualFloat takes 2 floats from the stack, pushes FF if equal, 00 if not
func (vm *VirtualMachine) operationEqualFloat() (err error) {
	operant1, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopFloat()
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

	vm.addLog("equal-float")
	return nil
}

// operationUnequalFloat takes 2 floats from the stack, pushes FF if unequal, 00 if not
func (vm *VirtualMachine) operationUnequalFloat() (err error) {
	operant1, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	result := byte(0x00)
	if operant1 != operant2 {
		result = byte(0xFF)
	}

	err = vm.stack.PushByte(result)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("unequal-float")
	return nil
}

// operationGreaterFloat takes 2 floats from the stack, pushes FF if the bottom one is greater, 00 otherwise
func (vm *VirtualMachine) operationGreaterFloat() (err error) {
	operant1, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	result := byte(0x00)
	if operant2 > operant1 {
		result = byte(0xFF)
	}

	err = vm.stack.PushByte(result)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("greater-float")
	return nil
}

// operationSmallerFloat takes 2 floats from the stack, pushes FF if the bottom one is smaller, 00 otherwise
func (vm *VirtualMachine) operationSmallerFloat() (err error) {
	operant1, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	operant2, err := vm.stack.PopFloat()
	if err != nil {
		return err
	}

	result := byte(0x00)
	if operant2 < operant1 {
		result = byte(0xFF)
	}

	err = vm.stack.PushByte(result)
	if err != nil {
		return err
	}

	vm.programPointer++

	vm.addLog("smaller-float")
	return nil
}
