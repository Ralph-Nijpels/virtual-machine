package virtualmachine

import (
	"testing"
	"unsafe"
)

// -- Tests ---------------------------------------------------------------------------------------------------------------------

func TestPushByte(t *testing.T) {
	testValue := byte(0x0C)

	p := NewProgram()
	p.WriteByte(0x08)      // Opcode: PushByte
	p.WriteByte(testValue) // Operant: testValue
	p.WriteByte(0x00)      // Opcode: End

	stack := make([]byte, (int)(unsafe.Sizeof(testValue)))
	*(*byte)(unsafe.Pointer(&stack[0])) = testValue

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetByte(t *testing.T) {
	testAddress := int(10)
	testValue := byte(0x91)

	p := NewProgram()
	p.WriteByte(0x10)       // Opcode: GetByte
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: End
	p.WriteByte(testValue)  // Data: testValue

	stack := make([]byte, (int)(unsafe.Sizeof(testValue)))
	*(*byte)(unsafe.Pointer(&stack[0])) = testValue

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutByte(t *testing.T) {
	testValue := byte(0xFE)
	testAddress := int(0x0C)

	p := NewProgram()
	p.WriteByte(0x08)       // Opcode: PushByte
	p.WriteByte(testValue)  // Operant: testValue
	p.WriteByte(0x18)       // Opcode: PutByte
	p.WriteInt(testAddress) // Operant: testAdress
	p.WriteByte(0x00)       // Opcode: End
	p.WriteByte(0x00)       // Data

	stack := make([]byte, 0)

	memory := make([]byte, p.Size())
	_, err := p.Read(memory)
	if err != nil {
		t.Errorf(err.Error())
	}

	*(*byte)(unsafe.Pointer(&(memory[len(memory)-(int)(unsafe.Sizeof(testValue))]))) = testValue

	err = p.Run(stack[:], memory[:])
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddByte(t *testing.T) {
	testValue1 := byte(0x04)
	testValue2 := byte(0x06)

	p := NewProgram()
	p.WriteByte(0x08)       // Opcode: push-byte
	p.WriteByte(testValue1) // Operant: testValue1
	p.WriteByte(0x08)       // Opcode: push-byte
	p.WriteByte(testValue2) // Operant: testValue2
	p.WriteByte(0x20)       // Opcode: add-byte
	p.WriteByte(0x00)       // Opcode: end

	testValue3 := testValue1 + testValue2
	stack := make([]byte, (int)(unsafe.Sizeof(testValue3)))
	*(*byte)(unsafe.Pointer(&stack[0])) = testValue3

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSubByte(t *testing.T) {
	testValue1 := byte(0x04)
	testValue2 := byte(0x06)

	p := NewProgram()
	p.WriteByte(0x08)       // Opcode: push-byte
	p.WriteByte(testValue1) // Operant: testValue1
	p.WriteByte(0x08)       // Opcode: push-byte
	p.WriteByte(testValue2) // Operant: testValue2
	p.WriteByte(0x28)       // Opcode: sub-byte
	p.WriteByte(0x00)       // Opcode: end

	testValue3 := testValue1 - testValue2
	stack := make([]byte, (int)(unsafe.Sizeof(testValue3)))
	*(*byte)(unsafe.Pointer(&stack[0])) = testValue3

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
