package virtualmachine

import (
	"fmt"
	"testing"
	"unsafe"
)

// -- Support functions ---------------------------------------------------------------------------------------------------------

// isBlocked checks if all interface functions are indeed blocked once in error
func (st *Stack) isBlocked() error {
	err := st.PushByte(0x20)
	if err.Error() != "blocked" {
		return fmt.Errorf("PushByte open")
	}

	_, err = st.GetByte(0)
	if err.Error() != "blocked" {
		return fmt.Errorf("GetByte open")
	}

	err = st.PutByte(0, 0)
	if err.Error() != "blocked" {
		return fmt.Errorf("PutByte open")
	}

	_, err = st.PopByte()
	if err.Error() != "blocked" {
		return fmt.Errorf("PopByte open")
	}

	err = st.PushInt(-1)
	if err.Error() != "blocked" {
		return fmt.Errorf("PushInt open")
	}

	_, err = st.GetInt(0)
	if err.Error() != "blocked" {
		return fmt.Errorf("GetInt open")
	}

	err = st.PutInt(0, 0)
	if err.Error() != "blocked" {
		return fmt.Errorf("GetInt open")
	}

	_, err = st.PopInt()
	if err.Error() != "blocked" {
		return fmt.Errorf("PopInt open")
	}

	err = st.PushFloat(12.50)
	if err.Error() != "blocked" {
		return fmt.Errorf("PushFloat open")
	}

	_, err = st.GetFloat(0)
	if err.Error() != "blocked" {
		return fmt.Errorf("GetFloat open")
	}

	err = st.PutFloat(0, 0.0)
	if err.Error() != "blocked" {
		return fmt.Errorf("PutFloat open")
	}

	_, err = st.PopFloat()
	if err.Error() != "blocked" {
		return fmt.Errorf("PopFloat open")
	}

	return nil
}

// -- Tests ---------------------------------------------------------------------------------------------------------------------

func TestStackPushByte(t *testing.T) {
	testValue := byte(0x20)

	st, err := NewStack(NewMemory(MEMORY_SIZE), STACK_SIZE)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Push one byte
	err = st.PushByte(testValue)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Check stack value
	expectedByte := [...]byte{
		0x20}
	err = st.Check(expectedByte[:])
	if err != nil {
		t.Errorf(err.Error())
	}

	// Fill the stack
	size := (int)(unsafe.Sizeof(testValue))
	for i := size; i < STACK_SIZE; i++ {
		err := st.PushByte(testValue)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	// Force overflow
	err = st.PushByte(testValue)
	if err == nil {
		t.Errorf("Expected: Overflow")
	}

	// Should be stuck
	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackPopByte(t *testing.T) {
	testValue := byte(0xAC)

	st, err := NewStack(NewMemory(MEMORY_SIZE), STACK_SIZE)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Push a byte
	err = st.PushByte(testValue)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Pop the byte
	value, err := st.PopByte()
	if err != nil {
		t.Errorf(err.Error())
	}

	expectedByte := [...]byte{}
	err = st.Check(expectedByte[:])
	if err != nil {
		t.Errorf(err.Error())
	}

	if value != testValue {
		t.Errorf("Expected: %d, got %d", testValue, value)
	}

	// Force underflow
	_, err = st.PopByte()
	if err == nil {
		t.Errorf("Expected: underflow")
	}

	// Should be blocked
	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackPushInt(t *testing.T) {
	testValue := int(160)

	st, err := NewStack(NewMemory(MEMORY_SIZE), STACK_SIZE)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Push one int
	err = st.PushInt(testValue)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Check value on stack
	expectedStack := [...]byte{
		0xA0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	err = st.Check(expectedStack[:])
	if err != nil {
		t.Errorf(err.Error())
	}

	// Force overflow
	size := (int)(unsafe.Sizeof(testValue))
	for i := size; i < STACK_SIZE; i += size {
		err = st.PushInt(testValue)
		if err != nil {
			t.Errorf(err.Error())
		}
	}
	err = st.PushInt(testValue)
	if err == nil {
		t.Errorf("Expected: overflow")
	}

	// Should be blocked
	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackPopInt(t *testing.T) {
	testValue := int(160)

	st, err := NewStack(NewMemory(MEMORY_SIZE), STACK_SIZE)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = st.PushInt(testValue)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Pop the int
	value, err := st.PopInt()
	if err != nil {
		t.Errorf(err.Error())
	}
	if value != testValue {
		t.Errorf("Expected: %d, got %d", value, testValue)
	}

	// Check stack
	expectedStack := [...]byte{}
	err = st.Check(expectedStack[:])
	if err != nil {
		t.Errorf(err.Error())
	}

	// Force an underflow
	_, err = st.PopInt()
	if err == nil {
		t.Errorf("Expected: underflow")
	}

	// Should be blocked
	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackPushFloat(t *testing.T) {
	testValue := float64(12.50)

	st, err := NewStack(NewMemory(MEMORY_SIZE), STACK_SIZE)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = st.PushFloat(testValue)
	if err != nil {
		t.Errorf(err.Error())
	}

	// Check the stack value
	expectedValue := [...]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x29, 0x40}
	err = st.Check(expectedValue[:])
	if err != nil {
		t.Errorf(err.Error())
	}

	// Load it to the max
	sizeValue := (int)(unsafe.Sizeof(testValue))
	for i := sizeValue; i < STACK_SIZE; i += sizeValue {
		err = st.PushFloat(testValue)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	// Should break now
	err = st.PushFloat(testValue)
	if err == nil {
		t.Errorf("expected: overflow")
	}

	// Should be blocked
	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackPopFloat(t *testing.T) {
	testValue := float64(12.50)

	st, err := NewStack(NewMemory(MEMORY_SIZE), STACK_SIZE)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = st.PushFloat(testValue)
	if err != nil {
		t.Errorf(err.Error())
	}

	value, err := st.PopFloat()
	if err != nil {
		t.Errorf(err.Error())
	}
	if value != testValue {
		t.Errorf("Expected: %f, got %f", testValue, value)
	}

	_, err = st.PopFloat()
	if err == nil {
		t.Errorf("Expected: Underflow")
	}

	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}
