package virtualmachine

import "unsafe"

// operationPushFloat takes the following 8 bytes and pushes them on the stack as a float
func (vm *VirtualMachine) operationPushFloat() error {
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
