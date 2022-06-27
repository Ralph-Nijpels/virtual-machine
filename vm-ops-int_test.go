package virtualmachine

import "testing"

func TestPushInt(t *testing.T) {
	testValue := int(-325)

	p := NewProgram()
	p.WriteByte(0x09)     // Opcode: push-int
	p.WriteInt(testValue) // Operant: testValue
	p.WriteByte(0x00)     // Opcode: end

	stack := [...]byte{
		0xBB, 0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

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

	stack := [...]byte{
		0xBB, 0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

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

	stack := [...]byte{}

	memory := [...]byte{
		0x09,                                           // see program
		0xBB, 0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, // ..
		0x19,                                           // ..
		0x13, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ..
		0x00, // End Program
		0xBB, 0xFE, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

	err := p.Run(stack[:], memory[:])
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

	stack := [...]byte{
		0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
