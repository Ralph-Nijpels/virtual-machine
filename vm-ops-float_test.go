package virtualmachine

import (
	"testing"
)

func TestPushFloat(t *testing.T) {
	testValue := float64(-120.7)

	p := NewProgram()
	p.WriteByte(0x0A)       // Opcode: push-float
	p.WriteFloat(testValue) // Operant: testValue
	p.WriteByte(0x00)       // Opcode: end

	s := NewBuffer()
	s.WriteFloat(testValue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetFloat(t *testing.T) {
	testAddress := int(0x0B)
	testValue := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x12)       // OpCode: get-float
	p.WriteByte(0x00)       // Opcode: end
	p.WriteFloat(testValue) // Data: testValue

	s := NewBuffer()
	s.WriteFloat(testValue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutFloat(t *testing.T) {
	testAddress := int(0x14)
	testValue := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x0A)       // Opcode: push-float
	p.WriteFloat(testValue) // Operant: testValue
	p.WriteByte(0x09)       // Opcode: push-int
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x1A)       // Opcode: put-float
	p.WriteByte(0x00)       // Opcode: end

	s := NewBuffer()

	m := NewBuffer()
	m.Copy(&p.Buffer)
	m.WriteFloat(testValue)

	err := p.Run(s, m)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetFloatAddress(t *testing.T) {
	testAddress := int(0x0A)
	testValue := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x22)       // Opcode: get-float
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end
	p.WriteFloat(testValue) // Data: testValue

	s := NewBuffer()
	s.WriteFloat(testValue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutFloatAddress(t *testing.T) {
	testAddress := int(0x13)
	testValue := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x0A)       // Opcode: push-float
	p.WriteFloat(testValue) // Operant: testValue
	p.WriteByte(0x2A)       // Opcode: put-float()
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end

	s := NewBuffer()

	m := NewBuffer()
	m.Copy(&p.Buffer)
	m.WriteFloat(testValue)

	err := p.Run(s, m)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetFloatStack(t *testing.T) {
	testAddress := int(-8)
	testValue := float64(123.45)

	p := NewProgram()
	p.WriteByte(0x0A)       // Opcode: push-float
	p.WriteFloat(testValue) // Operant: testValue
	p.WriteByte(0x32)       // Opcode: get-float{}
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end

	s := NewBuffer()
	s.WriteFloat(testValue)
	s.WriteFloat(testValue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutFloatStack(t *testing.T) {
	testAddress := int(-8)
	testValue := float64(123.45)

	p := NewProgram()
	p.WriteByte(0x0A)       // Opcode: push-float
	p.WriteFloat(0.0)       // Operant: <empty>
	p.WriteByte(0x0A)       // Opcode: push-float
	p.WriteFloat(testValue) // Operant: testValue
	p.WriteByte(0x3A)       // Opcode: put-float{}
	p.WriteInt(testAddress) // Operant: testAddress
	p.WriteByte(0x00)       // Opcode: end

	s := NewBuffer()
	s.WriteFloat(testValue)

	err := p.Run(s, nil)
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
	p.WriteByte(0x42)        // Opcode: add-float
	p.WriteByte(0x00)        // Opcode: end

	testValue3 := testValue1 + testValue2
	s := NewBuffer()
	s.WriteFloat(testValue3)

	err := p.Run(s, nil)
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
	p.WriteByte(0x46)        // Opcode: sub-float
	p.WriteByte(0x00)        // Opcode: end

	testValue3 := testValue1 - testValue2
	s := NewBuffer()
	s.WriteFloat(testValue3)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestMulFloat(t *testing.T) {
	testValue1 := float64(123.50)
	testValue2 := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue1) // Operant: testValue1
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue2) // Operant: testValue2
	p.WriteByte(0x4A)        // Opcode: mul-float
	p.WriteByte(0x00)        // Opcode: end

	testValue3 := testValue1 * testValue2
	s := NewBuffer()
	s.WriteFloat(testValue3)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestDivFloat(t *testing.T) {
	testValue1 := float64(123.50)
	testValue2 := float64(-12.34)

	p := NewProgram()
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue1) // Operant: testValue1
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue2) // Operant: testValue2
	p.WriteByte(0x4E)        // Opcode: mul-float
	p.WriteByte(0x00)        // Opcode: end

	testValue3 := testValue1 / testValue2
	s := NewBuffer()
	s.WriteFloat(testValue3)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestEqualFloat(t *testing.T) {
	testValue1 := float64(12.34)
	testValue2 := float64(12.34) // should be equal
	testValue3 := float64(0.0)   // should not be equal

	p := NewProgram()
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue1) // Operant: testValue1
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue2) // Operant: testValue2
	p.WriteByte(0x62)        // Opcode: equal-float
	p.WriteByte(0x00)        // Opcode: end

	s := NewBuffer()
	s.WriteByte(byte(0xFF)) // true

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue1) // Operant: testValue1
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue3) // Operant: testValue3
	p.WriteByte(0x62)        // Opcode: equal-float
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()
	s.WriteByte(byte(0x00)) // false

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestUnequalFloat(t *testing.T) {
	testValue1 := float64(12.34)
	testValue2 := float64(0.0)   // should be unequal
	testValue3 := float64(12.34) // should not be unequal

	p := NewProgram()
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue1) // Operant: testValue1
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue2) // Operant: testValue2
	p.WriteByte(0x66)        // Opcode: unequal-float
	p.WriteByte(0x00)        // Opcode: end

	s := NewBuffer()
	s.WriteByte(byte(0xFF)) // true

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue1) // Operant: testValue1
	p.WriteByte(0x0A)        // Opcode: push-float
	p.WriteFloat(testValue3) // Operant: testValue3
	p.WriteByte(0x66)        // Opcode: unequal-float
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()
	s.WriteByte(byte(0x00)) // false

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
