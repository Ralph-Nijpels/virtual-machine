package virtualmachine

import (
	"testing"
	"unsafe"
)

func TestPushFloat(t *testing.T) {
	testValue := float64(-120.7)

	p := NewProgram()
	p.WriteByte(0x0A)       // Opcode: PushFloat
	p.WriteFloat(testValue) // Operant: testvalue
	p.WriteByte(0x00)       // Opcode: end

	stack := make([]byte, (int)(unsafe.Sizeof(testValue)))
	*(*float64)(unsafe.Pointer(&stack[0])) = testValue

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
