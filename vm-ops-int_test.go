package virtualmachine

import (
	"testing"
	"unsafe"
)

func TestPushInt(t *testing.T) {
	testValue := int(-325)

	p := NewProgram()
	p.WriteByte(0x09)     // Opcode: push-int
	p.WriteInt(testValue) // Operant: testValue
	p.WriteByte(0x00)     // Opcode: end

	stack := make([]byte, (int)(unsafe.Sizeof(testValue)))
	*(*int)(unsafe.Pointer(&stack[0])) = testValue

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetInt(t *testing.T) {
	testAddress := int(0x0A)
	testValue := int(-325)

	p := NewProgram()
	p.WriteByte(0x11)       // Opcode: get-int
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end
	p.WriteInt(testValue)   // Data: testValue

	stack := make([]byte, (int)(unsafe.Sizeof(testValue)))
	*(*int)(unsafe.Pointer(&stack[0])) = testValue

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutInt(t *testing.T) {
	testAddress := int(0x13)
	testValue := int(-325)

	p := NewProgram()
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(testValue)   // Operant: testValue
	p.WriteByte(0x19)       // Opcode: put-int
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end
	p.WriteInt(0)           // Data: <empty>

	stack := make([]byte, 0)

	memory := make([]byte, p.Size())
	_, err := p.Read(memory)
	if err != nil {
		t.Errorf(err.Error())
	}

	*(*int)(unsafe.Pointer(&(memory[len(memory)-(int)(unsafe.Sizeof(testValue))]))) = testValue

	err = p.Run(stack[:], memory[:])
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddInt(t *testing.T) {
	testValue1 := int(0x04)
	testValue2 := int(0x06)

	p := NewProgram()
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue1) // Operant: testValue1
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue2) // Operant: testValue2
	p.WriteByte(0x21)      // Opcode: add-int
	p.WriteByte(0x00)      // Opcode: end

	testValue3 := testValue1 + testValue2
	stack := make([]byte, (int)(unsafe.Sizeof(testValue3)))
	*(*int)(unsafe.Pointer(&stack[0])) = testValue3

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSubInt(t *testing.T) {
	testValue1 := int(12)
	testValue2 := int(6)

	p := NewProgram()
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue1) // Operant: testValue1
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue2) // Operant: testValue2
	p.WriteByte(0x29)      // Opcode: sub-int
	p.WriteByte(0x00)      // Opcode: end

	testValue3 := testValue1 - testValue2
	stack := make([]byte, (int)(unsafe.Sizeof(testValue3)))
	*(*int)(unsafe.Pointer(&stack[0])) = testValue3

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
