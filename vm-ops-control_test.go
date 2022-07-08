package virtualmachine

import "testing"

func TestRet(t *testing.T) {
	testAddress1 := int(11)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueOK := int(0x5A5A5A5A5A5A5A5A) // If we ended up where we wanted to be

	p := NewProgram()
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testAddress1) // Operant: testAddress1
	p.WriteByte(0xE0)        // Opcode: ret
	p.WriteByte(0x00)        // Opcode: end
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testValueOK)  // Operant: testValueOK
	p.WriteByte(0x00)        // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueOK)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testAddress2) // Operant: testAddress2
	p.WriteByte(0xE0)        // Opcode: ret
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()
	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testAddress3) // Operant: testAddress3
	p.WriteByte(0xE0)        // Opcode: ret
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()
	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmp(t *testing.T) {
	testAddress1 := int(10)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueOK := int(0x5A5A5A5A5A5A5A5A) // If we ended up where we wanted to be

	p := NewProgram()
	p.WriteByte(0xE1)        // Opcode: jmp()
	p.WriteInt(testAddress1) // Operant: testAddress1
	p.WriteByte(0x00)        // Opcode: end
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testValueOK)  // Operant: testValueOK
	p.WriteByte(0x00)        // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueOK)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0xE1)        // Opcode: jmp()
	p.WriteInt(testAddress2) // Operant: testAddress2
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0xE1)        // Opcode: jmp
	p.WriteInt(testAddress3) // Operant: testAddress3
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpzByte(t *testing.T) {
	testAddress1 := int(22)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0x00)          // Operant: 0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xE4)          // Opcode: jmpz-byte
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0xFF)          // Operant: FF (no jump)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xE4)          // Opcode: jmpz-byte
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0x00)          // Operant: 0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0xE4)          // Opcode: jmpz-byte
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0x00)          // Operant: 0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0xE4)          // Opcode: jmpz-byte
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpzInt(t *testing.T) {
	testAddress1 := int(29)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(0)              // Operant: 0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xE5)          // Opcode: jmpz-int
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(-1)             // Operant: -1 (no jump)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xE5)          // Opcode: jmpz-int
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(0)              // Operant: 0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0xE5)          // Opcode: jmpz-int
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(0)              // Operant: 0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0xE5)          // Opcode: jmpz-int
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpzFloat(t *testing.T) {
	testAddress1 := int(29)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(0.0)          // Operant: 0.0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xE6)          // Opcode: jmpz-float
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(-1.0)         // Operant: -1.0 (no jump)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xE6)          // Opcode: jmpz-float
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(0.0)          // Operant: 0.0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0xE6)          // Opcode: jmpz-float
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(0.0)          // Operant: 0.0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0xE6)          // Opcode: jmpz-float
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpzByteAddress(t *testing.T) {
	testAddress1 := int(21)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0x00)          // Operant: 0 (jumps)
	p.WriteByte(0xE8)          // Opcode: jmpz-byte()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0xFF)          // Operant: FF (no jump)
	p.WriteByte(0xE8)          // Opcode: jmpz-byte()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0x00)          // Operant: 0 (jumps)
	p.WriteByte(0xE8)          // Opcode: jmpz-byte()
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0x00)          // Operant: 0 (jumps)
	p.WriteByte(0xE8)          // Opcode: jmpz-byte()
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpzIntAddress(t *testing.T) {
	testAddress1 := int(28)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(0)              // Operant: 0 (jumps)
	p.WriteByte(0xE9)          // Opcode: jmpz-int()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(-1)             // Operant: -1 (no jump)
	p.WriteByte(0xE9)          // Opcode: jmpz-int()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(0)              // Operant: 0 (jumps)
	p.WriteByte(0xE9)          // Opcode: jmpz-int()
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(0)              // Operant: 0 (jumps)
	p.WriteByte(0xE9)          // Opcode: jmpz-int()
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpzFloatAddress(t *testing.T) {
	testAddress1 := int(28)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(0.0)          // Operant: 0.0 (jumps)
	p.WriteByte(0xEA)          // Opcode: jmpz-float()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(-1.0)         // Operant: -1.0 (no jump)
	p.WriteByte(0xEA)          // Opcode: jmpz-float()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(0.0)          // Operant: 0.0 (jumps)
	p.WriteByte(0xEA)          // Opcode: jmpz-float()
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(0.0)          // Operant: 0.0 (jumps)
	p.WriteByte(0xEA)          // Opcode: jmpz-float()
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpnzByte(t *testing.T) {
	testAddress1 := int(22)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0xFF)          // Operant: FF (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xEC)          // Opcode: jmpnz-byte
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0x00)          // Operant: 00 (no jump)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xEC)          // Opcode: jmpnz-byte
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0xFF)          // Operant: FF (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0xEC)          // Opcode: jmpnz-byte
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0xFF)          // Operant: FF (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0xEC)          // Opcode: jmpz-byte
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpnzInt(t *testing.T) {
	testAddress1 := int(29)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(-1)             // Operant: -1 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xED)          // Opcode: jmpnz-int
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(0)              // Operant: 0 (no jump)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xED)          // Opcode: jmpnz-int
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(-1)             // Operant: -1 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0xED)          // Opcode: jmpnz-int
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(-1)             // Operant: -1 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0xED)          // Opcode: jmpnz-int
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpnzFloat(t *testing.T) {
	testAddress1 := int(29)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(1.0)          // Operant: 1.0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xEE)          // Opcode: jmpnz-float
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(0.0)          // Operant: 0.0 (no jump)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0xEE)          // Opcode: jmpnz-float
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(1.0)          // Operant: 1.0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0xEE)          // Opcode: jmpnz-float
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(1.0)          // Operant: 0.0 (jumps)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0xEE)          // Opcode: jmpnz-float
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpnzByteAddress(t *testing.T) {
	testAddress1 := int(21)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0xFF)          // Operant: FF (jumps)
	p.WriteByte(0xF0)          // Opcode: jmpnz-byte()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0x00)          // Operant: 00 (no jump)
	p.WriteByte(0xF0)          // Opcode: jmpnz-byte()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0xFF)          // Operant: FF (jumps)
	p.WriteByte(0xF0)          // Opcode: jmpnz-byte()
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x08)          // Opcode: push-byte
	p.WriteByte(0xFF)          // Operant: FF (jumps)
	p.WriteByte(0xF0)          // Opcode: jmpz-byte()
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpnzIntAddress(t *testing.T) {
	testAddress1 := int(28)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(-1)             // Operant: -1 (jumps)
	p.WriteByte(0xF1)          // Opcode: jmpnz-int()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(0)              // Operant: 0 (no jump)
	p.WriteByte(0xF1)          // Opcode: jmpnz-int()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(-1)             // Operant: -1 (jumps)
	p.WriteByte(0xF1)          // Opcode: jmpnz-int()
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(-1)             // Operant: -1 (jumps)
	p.WriteByte(0xF1)          // Opcode: jmpnz-int()
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestJmpnzFloatAddress(t *testing.T) {
	testAddress1 := int(28)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testValueTrue := int(0x5A5A5A5A5A5A5A5A)  // If we ended up where we wanted to be under 'true' condition
	testValueFalse := int(0x5555555555555555) // If we ended up where we wanted to be under 'false' condition

	p := NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(1.0)          // Operant: 1.0 (jumps)
	p.WriteByte(0xF2)          // Opcode: jmpnz-float()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s := NewBuffer()
	s.WriteInt(testValueTrue)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(0.0)          // Operant: 0.0 (no jump)
	p.WriteByte(0xF2)          // Opcode: jmpnz-float()
	p.WriteInt(testAddress1)   // Operant: testAddress1
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()
	s.WriteInt(testValueFalse)

	err = p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(1.0)          // Operant: 1.0 (jumps)
	p.WriteByte(0xF2)          // Opcode: jmpnz-float()
	p.WriteInt(testAddress2)   // Operant: testAddress2 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x0A)          // Opcode: push-float
	p.WriteFloat(1.0)          // Operant: 0.0 (jumps)
	p.WriteByte(0xF2)          // Opcode: jmpnz-float()
	p.WriteInt(testAddress3)   // Operant: testAddress3 (illegal)
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueFalse) // Operant: testValueFalse
	p.WriteByte(0x00)          // Opcode: end
	p.WriteByte(0x09)          // Opcode: push-int
	p.WriteInt(testValueTrue)  // Operant: testValueOK
	p.WriteByte(0x00)          // Opcode: end

	s = NewBuffer()

	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestCall(t *testing.T) {
	testAddress1 := int(11)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testAddressOK := int(10)               // Where we return...
	testValueOK := int(0x5A5A5A5A5A5A5A5A) // If we ended up where we wanted to be

	p := NewProgram()
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testAddress1) // Operant: testAddress1
	p.WriteByte(0xF8)        // Opcode: call
	p.WriteByte(0x00)        // Opcode: end
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testValueOK)  // Operant: testValueOK
	p.WriteByte(0x00)        // Opcode: end

	s := NewBuffer()
	s.WriteInt(testAddressOK)
	s.WriteInt(testValueOK)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testAddress2) // Operant: testAddress2
	p.WriteByte(0xF8)        // Opcode: call
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()
	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testAddress3) // Operant: testAddress3
	p.WriteByte(0xF8)        // Opcode: call
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()
	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}

func TestCallAddress(t *testing.T) {
	testAddress1 := int(10)              // Fine address
	testAddress2 := int(-1)              // Outside memory
	testAddress3 := int(MEMORY_SIZE + 1) // Outside memory

	testAddressOK := int(9)                // Where we return...
	testValueOK := int(0x5A5A5A5A5A5A5A5A) // If we ended up where we wanted to be

	p := NewProgram()
	p.WriteByte(0xF9)        // Opcode: call()
	p.WriteInt(testAddress1) // Operant: testAddress1
	p.WriteByte(0x00)        // Opcode: end
	p.WriteByte(0x09)        // Opcode: push-int
	p.WriteInt(testValueOK)  // Operant: testValueOK
	p.WriteByte(0x00)        // Opcode: end

	s := NewBuffer()
	s.WriteInt(testAddressOK)
	s.WriteInt(testValueOK)

	err := p.Run(s, nil)
	if err != nil {
		t.Errorf(err.Error())
	}

	p = NewProgram()
	p.WriteByte(0xF9)        // Opcode: call()
	p.WriteInt(testAddress2) // Operant: testAddress2
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()
	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}

	p = NewProgram()
	p.WriteByte(0xF9)        // Opcode: call
	p.WriteInt(testAddress3) // Operant: testAddress3
	p.WriteByte(0x00)        // Opcode: end

	s = NewBuffer()
	err = p.Run(s, nil)
	if err.Error() != "illegal address" {
		t.Errorf("Expected: illegal address")
	}
}
