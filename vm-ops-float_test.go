package virtualmachine

import "testing"

func TestPushFloat(t *testing.T) {
	testValue := float64(-120.7)

	p := NewProgram()
	p.WriteByte(0x0A)       // Opcode: PushFloat
	p.WriteFloat(testValue) // Operant: testvalue
	p.WriteByte(0x00)       // Opcode: end

	stack := [...]byte{
		0xCD, 0xCC, 0xCC, 0xCC, 0xCC, 0x2C, 0x5E, 0xC0}

	err := p.Run(stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
