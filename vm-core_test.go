package virtualmachine

import (
	"fmt"
	"unsafe"
)

// -- Reasonable values for testing ---------------------------------------------------------------------------------------------

const MEMORY_SIZE = 256
const STACK_SIZE = 64

// -- Easier way to build test programs, bytes.Buffer doesn't satisfy the need --------------------------------------------------
// Notes: May Panic. It doesn't grow beyond MEMORY_SIZE.

type Program struct {
	code [MEMORY_SIZE]byte
	len  int
}

func (p *Program) WriteByte(value byte) (err error) {
	if p.len+(int)(unsafe.Sizeof(value)) > len(p.code) {
		return fmt.Errorf("buffer overflow")
	}

	*(*byte)(unsafe.Pointer(&p.code[p.len])) = value
	p.len += (int)(unsafe.Sizeof(value))

	return nil
}

func (p *Program) WriteInt(value int) (err error) {
	if p.len+(int)(unsafe.Sizeof(value)) > len(p.code) {
		return fmt.Errorf("buffer overflow")
	}

	*(*int)(unsafe.Pointer(&p.code[p.len])) = value
	p.len += (int)(unsafe.Sizeof(value))

	return nil
}

func (p *Program) WriteFloat(value float64) (err error) {
	if p.len+(int)(unsafe.Sizeof(value)) > len(p.code) {
		return fmt.Errorf("buffer overflow")
	}

	*(*float64)(unsafe.Pointer(&p.code[p.len])) = value
	p.len += (int)(unsafe.Sizeof(value))

	return nil
}

func (p *Program) Read(buffer []byte) (size int, err error) {
	if len(buffer) < p.len {
		return 0, fmt.Errorf("buffer too small")
	}

	copy(buffer, p.code[:p.len])

	return p.len, nil
}

func (p *Program) Size() (size int) {
	return p.len
}

func (p *Program) Run(expectedStack []byte, expectedMemory []byte) (err error) {
	vm := NewVirtualMachine(MEMORY_SIZE, STACK_SIZE)

	err = vm.Load(p.code[:p.len])
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
