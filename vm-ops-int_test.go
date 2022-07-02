package virtualmachine

import (
	"testing"
)

func TestPushInt(t *testing.T) {
	testValue := int(-325)

	p := NewProgram()
	p.WriteByte(0x09)     // Opcode: push-int
	p.WriteInt(testValue) // Operant: testValue
	p.WriteByte(0x00)     // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetInt(t *testing.T) {
	testAddress := int(0x0B)
	testValue := int(-325)

	p := NewProgram()
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x11)       // Opcode: get-int
	p.WriteByte(0x00)       // Opcode: end
	p.WriteInt(testValue)   // Data: testValue

	s := NewBuffer()
	s.WriteInt(testValue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetIntAddress(t *testing.T) {
	testAddress := int(0x0A)
	testValue := int(-325)

	p := NewProgram()
	p.WriteByte(0x21)       // Opcode: get-int()
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end
	p.WriteInt(testValue)   // Data: testValue

	s := NewBuffer()
	s.WriteInt(testValue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutInt(t *testing.T) {
	testAddress := int(0x14)
	testValue := int(-325)

	p := NewProgram()
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(testValue)   // Operant: testValue
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x19)       // Opcode: put-int
	p.WriteByte(0x00)       // Opcode: end

	s := NewBuffer()

	m := NewBuffer()
	m.Copy(&p.Buffer)
	m.WriteInt(testValue)

	err := p.Run(s, m)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutIntAddress(t *testing.T) {
	testAddress := int(0x13)
	testValue := int(-325)

	p := NewProgram()
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(testValue)   // Operant: testValue
	p.WriteByte(0x29)       // Opcode: put-int()
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end

	s := NewBuffer()

	m := NewBuffer()
	m.Copy(&p.Buffer)
	m.WriteInt(testValue)

	err := p.Run(s, m)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetIntStack(t *testing.T) {
	testAddress := int(-8)
	testValue := int(-332)

	p := NewProgram()
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(testValue)   // Operant: testValue
	p.WriteByte(0x31)       // Opcode: get-int{}
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValue)
	s.WriteInt(testValue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutIntStack(t *testing.T) {
	testAddress := int(-8)
	testValue := int(-332)

	p := NewProgram()
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(0)           // Operant: <empty>
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(testValue)   // Operant: testValue
	p.WriteByte(0x39)       // Opcode: put-int{}
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValue)

	err := p.Run(s, nil)
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
	p.WriteByte(0x41)      // Opcode: add-int
	p.WriteByte(0x00)      // Opcode: end

	testValue3 := testValue1 + testValue2
	s := NewBuffer()
	s.WriteInt(testValue3)

	err := p.Run(s, nil)
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
	p.WriteByte(0x45)      // Opcode: sub-int
	p.WriteByte(0x00)      // Opcode: end

	testValue3 := testValue1 - testValue2
	s := NewBuffer()
	s.WriteInt(testValue3)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestMulInt(t *testing.T) {
	testValue1 := int(12)
	testValue2 := int(6)

	p := NewProgram()
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue1) // Operant: testValue1
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue2) // Operant: testValue2
	p.WriteByte(0x49)      // Opcode: sub-int
	p.WriteByte(0x00)      // Opcode: end

	testValue3 := testValue1 * testValue2
	s := NewBuffer()
	s.WriteInt(testValue3)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestDivInt(t *testing.T) {
	testValue1 := int(12)
	testValue2 := int(6)

	p := NewProgram()
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue1) // Operant: testValue1
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue2) // Operant: testValue2
	p.WriteByte(0x4D)      // Opcode: div-int
	p.WriteByte(0x00)      // Opcode: end

	testValue3 := testValue1 / testValue2
	s := NewBuffer()
	s.WriteInt(testValue3)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestEqualInt(t *testing.T) {
	testValue1 := int(-325)
	testValue2 := int(-325) // should be equal
	testValue3 := int(0)    // should not be equal

	p := NewProgram()
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue1) // Operant: testValue1
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue2) // Operant: testValue2
	p.WriteByte(0x61)      // Opcode: equal-int
	p.WriteByte(0x00)      // Opcode: end

	s := NewBuffer()
	s.WriteByte(byte(0xFF)) // true

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)      // Opcode: push-int
	p.WriteInt(testValue1) // Operant: testValue1
	p.WriteByte(0x09)      // Opcode: push-byte
	p.WriteInt(testValue3) // Operant: testValue3
	p.WriteByte(0x61)      // Opcode: equal int
	p.WriteByte(0x00)      // Opcode: end

	s = NewBuffer()
	s.WriteByte(byte(0x00)) // false

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
