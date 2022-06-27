package virtualmachine

import (
	"unsafe"
)

// -- Reasonable values for testing ---------------------------------------------------------------------------------------------

const MEMORY_SIZE = 256
const STACK_SIZE = 64

// -- Easier way to build test programs, bytes.Buffer doesn't satisfy the need --------------------------------------------------

type Program struct {
	code [MEMORY_SIZE]byte
	len  int
}

func (p *Program) WriteByte(value byte) error {
	p.code[p.len] = value
	p.len++

	return nil
}

func (p *Program) WriteInt(value int) error {
	*(*int)(unsafe.Pointer(&p.code[p.len])) = value
	p.len += (int)(unsafe.Sizeof(value))

	return nil
}

func (p *Program) WriteFloat(value float64) error {
	*(*float64)(unsafe.Pointer(&p.code[p.len])) = value
	p.len += (int)(unsafe.Sizeof(value))

	return nil
}

func (p *Program) Run(expectedStack []byte, expectedMemory []byte) error {
	vm := NewVirtualMachine(MEMORY_SIZE, STACK_SIZE)

	err := vm.Load(p.code[:p.len])
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

func NewProgram() *Program {
	return new(Program)
}

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
