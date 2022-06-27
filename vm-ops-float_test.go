package virtualmachine

import (
	"testing"
	"unsafe"
)

func TestPushFloat(t *testing.T) {
	testValue := float64(-120.7)

	p := NewProgram()
	p.WriteByte(0x0A)       // Opcode: push-float
	p.WriteFloat(testValue) // Operant: testValue
	p.WriteByte(0x00)       // Opcode: end

	stack := make([]byte, (int)(unsafe.Sizeof(testValue)))
	*(*float64)(unsafe.Pointer(&stack[0])) = testValue

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetFloat(t *testing.T) {
	testAddress := int(0x0A)
	testValue := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x12)       // Opcode: get-float
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end
	p.WriteFloat(testValue) // Data: testValue

	stack := make([]byte, (int)(unsafe.Sizeof(testValue)))
	*(*float64)(unsafe.Pointer(&stack[0])) = testValue

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutFloat(t *testing.T) {
	testAddress := int(0x13)
	testValue := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x0A)       // Opcode: push-float
	p.WriteFloat(testValue) // Operant: testValue
	p.WriteByte(0x1A)       // Opcode: put-float
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end
	p.WriteFloat(0)         // Data: <empty>

	stack := make([]byte, 0)

	memory := make([]byte, p.Size())
	_, err := p.Read(memory)
	if err != nil {
		t.Errorf(err.Error())
	}

	*(*float64)(unsafe.Pointer(&(memory[len(memory)-(int)(unsafe.Sizeof(testValue))]))) = testValue

	err = p.Run(stack[:], memory[:])
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddFloat(t *testing.T) {
	testValue1 := float64(123.50)
	testValue2 := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue1) // Operant: testValue1
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue2) // Operant: testValue2
	p.WriteByte(0x22)        // Opcode: add-float
	p.WriteByte(0x00)        // Opcode: end

	testValue3 := testValue1 + testValue2
	stack := make([]byte, (int)(unsafe.Sizeof(testValue3)))
	*(*float64)(unsafe.Pointer(&stack[0])) = testValue3

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSubFloat(t *testing.T) {
	testValue1 := float64(123.50)
	testValue2 := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue1) // Operant: testValue1
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue2) // Operant: testValue2
	p.WriteByte(0x2A)        // Opcode: sub-float
	p.WriteByte(0x00)        // Opcode: end

	testValue3 := testValue1 - testValue2
	stack := make([]byte, (int)(unsafe.Sizeof(testValue3)))
	*(*float64)(unsafe.Pointer(&stack[0])) = testValue3

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
