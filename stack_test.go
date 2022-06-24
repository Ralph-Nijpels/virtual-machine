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
	if err.Error() != "Blocked" {
		return fmt.Errorf("PushByte open")
	}

	_, err = st.PopByte()
	if err.Error() != "Blocked" {
		return fmt.Errorf("PopByte open")
	}

	err = st.PushInt(-1)
	if err.Error() != "Blocked" {
		return fmt.Errorf("PushInt open")
	}

	_, err = st.PopInt()
	if err.Error() != "Blocked" {
		return fmt.Errorf("PopInt open")
	}

	return nil
}

// -- Tests ---------------------------------------------------------------------------------------------------------------------

func TestStackPushByte(t *testing.T) {
	st := NewStack(STACK_SIZE)

	expectedByte := [...]byte{
		0x20}

	// Push one byte
	err := st.PushByte(0x20)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = st.Check(expectedByte[:])
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackPushInt(t *testing.T) {
	st := NewStack(STACK_SIZE)

	stack := [...]byte{
		0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	// Push one int
	err := st.PushInt(0x20)
	if err != nil {
		t.Errorf(err.Error())
	}

	err = st.Check(stack[:])
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackOverflowByte(t *testing.T) {
	st := NewStack(STACK_SIZE)

	// loading bytes
	for i := 0; i < STACK_SIZE; i++ {
		err := st.PushByte(0x20)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	// One more should do it
	err := st.PushByte(0x20)
	if err == nil {
		t.Errorf("Expected Overflow")
	}

	// Should be blocked
	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackOverflowInt(t *testing.T) {
	st := NewStack(STACK_SIZE)

	// loading ints
	for i := 0; i < (STACK_SIZE / (int)(unsafe.Sizeof(i))); i++ {
		err := st.PushInt(i)
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	// one more
	err := st.PushInt(-1)
	if err == nil {
		t.Errorf("Expected Overflow")
	}

	// Should be blocked
	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackPopByte(t *testing.T) {
	st := NewStack(STACK_SIZE)

	stack := [...]byte{}

	err := st.PushByte(0x20)
	if err != nil {
		t.Errorf(err.Error())
	}

	value, err := st.PopByte()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = st.Check(stack[:])
	if err != nil {
		t.Errorf(err.Error())
	}

	if value != 0x20 {
		t.Errorf("Value mismatch")
	}
}

func TestStackPopInt(t *testing.T) {
	st := NewStack(STACK_SIZE)

	stack := [...]byte{}

	err := st.PushInt(-1)
	if err != nil {
		t.Errorf(err.Error())
	}

	value, err := st.PopInt()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = st.Check(stack[:])
	if err != nil {
		t.Errorf(err.Error())
	}

	if value != -1 {
		t.Errorf("Value mismatch")
	}
}

func TestStackUnderflowByte(t *testing.T) {
	st := NewStack(STACK_SIZE)

	// Cannot pop from empty stack
	_, err := st.PopByte()
	if err == nil {
		t.Errorf(err.Error())
	}

	// Should be blocked
	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestStackUnderflowInt(t *testing.T) {
	st := NewStack(STACK_SIZE)

	// Cannot pop from empty stack
	_, err := st.PopInt()
	if err == nil {
		t.Errorf(err.Error())
	}

	// Should be blocked
	err = st.isBlocked()
	if err != nil {
		t.Errorf(err.Error())
	}
}
