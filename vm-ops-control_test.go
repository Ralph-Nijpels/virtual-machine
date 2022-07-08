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
