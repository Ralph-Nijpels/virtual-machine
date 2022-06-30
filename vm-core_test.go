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

type Buffer struct {
	bytes [MEMORY_SIZE]byte
	len   int
}

func (b *Buffer) WriteByte(value byte) (err error) {
	if b.len+(int)(unsafe.Sizeof(value)) > len(b.bytes) {
		return fmt.Errorf("buffer overflow")
	}

	*(*byte)(unsafe.Pointer(&b.bytes[b.len])) = value
	b.len += (int)(unsafe.Sizeof(value))

	return nil
}

func (b *Buffer) WriteInt(value int) (err error) {
	if b.len+(int)(unsafe.Sizeof(value)) > len(b.bytes) {
		return fmt.Errorf("buffer overflow")
	}

	*(*int)(unsafe.Pointer(&b.bytes[b.len])) = value
	b.len += (int)(unsafe.Sizeof(value))

	return nil
}

func (b *Buffer) WriteFloat(value float64) (err error) {
	if b.len+(int)(unsafe.Sizeof(value)) > len(b.bytes) {
		return fmt.Errorf("buffer overflow")
	}

	*(*float64)(unsafe.Pointer(&b.bytes[b.len])) = value
	b.len += (int)(unsafe.Sizeof(value))

	return nil
}

func (b *Buffer) Copy(buffer *Buffer) (size int, err error) {
	if len(b.bytes) < buffer.len {
		return 0, fmt.Errorf("buffer too small")
	}

	copy(b.bytes[:buffer.len], buffer.bytes[:buffer.len])
	b.len = buffer.len

	return b.len, nil
}

func (b *Buffer) Value() (bytes []byte) {
	return b.bytes[:b.len]
}

func (b *Buffer) Size() (size int) {
	return b.len
}

func NewBuffer() (buffer *Buffer) {
	return new(Buffer)
}

type Program struct {
	Buffer
}

func (p *Program) Run(expectedStack *Buffer, expectedMemory *Buffer) (err error) {
	vm, err := NewVirtualMachine(MEMORY_SIZE, STACK_SIZE)
	if err != nil {
		return err
	}

	err = vm.Load(p.bytes[:p.len])
	if err != nil {
		return err
	}

	err = vm.Run()
	if err != nil {
		return err
	}

	if expectedStack != nil {
		err = vm.stack.Check(expectedStack.Value())
		if err != nil {
			return err
		}
	}

	if expectedMemory != nil {
		err = vm.memory.Check(expectedMemory.Value())
		if err != nil {
			return err
		}
	}

	return nil
}

func NewProgram() *Program {
	return new(Program)
}
