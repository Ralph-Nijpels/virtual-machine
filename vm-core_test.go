package virtualmachine

// -- Reasonable values for testing ---------------------------------------------------------------------------------------------

const MEMORY_SIZE = 256
const STACK_SIZE = 64

// -- Support functions ---------------------------------------------------------------------------------------------------------

func runProgram(program []byte, expectedStack []byte, expectedMemory []byte) error {
	vm := NewVirtualMachine(MEMORY_SIZE, STACK_SIZE)

	err := vm.Load(program)
	if err != nil {
		return err
	}

	err = vm.Run()
	if err != nil {
		return err
	}

	if expectedStack != nil {
		err = vm.stack.Check(expectedStack)
		if err != nil {
			return err
		}
	}

	if expectedMemory != nil {
		err = vm.memory.Check(expectedMemory)
		if err != nil {
			return err
		}
	}

	return nil
}
