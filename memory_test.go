package virtualmachine

import (
	"testing"
	"unsafe"
)

// -- Tests ---------------------------------------------------------------------------------------------------------------------

func TestMemoryPutByte(t *testing.T) {
	testValue := byte(0x20)

	// Test straight put
	mem := NewMemory(MEMORY_SIZE)
	err := mem.PutByte(0, testValue)
	if err != nil {
		t.Errorf(err.Error())
	}

	expectMem := [...]byte{testValue}
	err = mem.Check(expectMem[:])
	if err != nil {
		t.Errorf(err.Error())
	}

	// Test boundary
	err = mem.PutByte(MEMORY_SIZE-(int)(unsafe.Sizeof(testValue)), testValue)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Test boundary
	err = mem.PutByte(-1, testValue)
	if err == nil {
		t.Errorf("Expected out of bounds")
	}

	// Test boundary
	err = mem.PutByte(MEMORY_SIZE, testValue)
	if err == nil {
		t.Errorf("Expected out of bounds")
	}
}

func TestMemoryGetByte(t *testing.T) {
	testValue := byte(0x20)

	mem := NewMemory(MEMORY_SIZE)

	// First possible byte
	err := mem.PutByte(0, testValue)
	if err != nil {
		t.Errorf(err.Error())
	}
	value, err := mem.GetByte(0)
	if err != nil {
		t.Errorf(err.Error())
	}
	if value != testValue {
		t.Errorf("Expected %d, got %d", testValue, value)
	}

	// Last possible byte
	err = mem.PutByte(MEMORY_SIZE-(int)(unsafe.Sizeof(testValue)), testValue)
	if err != nil {
		t.Errorf(err.Error())
	}
	value, err = mem.GetByte(MEMORY_SIZE - (int)(unsafe.Sizeof(testValue)))
	if err != nil {
		t.Errorf(err.Error())
	}
	if value != testValue {
		t.Errorf("Expected %d, got %d", testValue, value)
	}

	// Out of bounds checks
	_, err = mem.GetByte(-1)
	if err == nil {
		t.Errorf("Expected a memory error")
	}
	_, err = mem.GetByte(MEMORY_SIZE)
	if err == nil {
		t.Errorf("Expected a memory error")
	}
}

func TestMemoryPutInt(t *testing.T) {
	testValue := int(0x3C0000000000003C)

	mem := NewMemory(MEMORY_SIZE)

	// First possible location
	err := mem.PutInt(0, testValue)
	if err != nil {
		t.Errorf(err.Error())
	}
	expectMem := [...]byte{
		0x3C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3C}
	err = mem.Check(expectMem[:])
	if err != nil {
		t.Errorf(err.Error())
	}

	// Last possible location
	err = mem.PutInt(MEMORY_SIZE-(int)(unsafe.Sizeof(testValue)), testValue)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Test boundary
	err = mem.PutInt(-1, testValue)
	if err == nil {
		t.Errorf("Expected memory error")
	}
	err = mem.PutInt(MEMORY_SIZE-(int)(unsafe.Sizeof(testValue))+1, testValue)
	if err == nil {
		t.Errorf("Expected memory error")
	}
}

func TestMemoryGetInt(t *testing.T) {
	testValue := int(0x3C0000000000003C)

	mem := NewMemory(MEMORY_SIZE)

	// First possible byte
	err := mem.PutInt(0, testValue)
	if err != nil {
		t.Errorf(err.Error())
	}
	value, err := mem.GetInt(0)
	if err != nil {
		t.Errorf(err.Error())
	}
	if value != testValue {
		t.Errorf("Expected %d, got %d", testValue, value)
	}

	// Last possible byte
	err = mem.PutInt(MEMORY_SIZE-(int)(unsafe.Sizeof(testValue)), testValue)
	if err != nil {
		t.Errorf(err.Error())
	}
	value, err = mem.GetInt(MEMORY_SIZE - (int)(unsafe.Sizeof(testValue)))
	if err != nil {
		t.Errorf(err.Error())
	}
	if value != testValue {
		t.Errorf("Expected %d, got %d", testValue, value)
	}

	// Out of bounds checks
	_, err = mem.GetInt(-1)
	if err == nil {
		t.Errorf("Expected a memory error")
	}
	_, err = mem.GetInt(MEMORY_SIZE - (int)(unsafe.Sizeof(testValue)) + 1)
	if err == nil {
		t.Errorf("Expected a memory error")
	}
}
