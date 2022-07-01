package virtualmachine

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/ttacon/chalk"
)

type Stack struct {
	mem       *Memory // Memory object used to store items on stack
	size      int     // Number of bytes 'allocated' to the stack
	offset    int     // Starting place in memory for the stack
	pointer   int     // current top of stack
	overflow  bool
	underflow bool
}

// -- Basic stack manipulation functions on bytes -------------------------------------------------------------------------------

// PushByte puts a byte on the stack
func (st *Stack) PushByte(value byte) (err error) {
	if st.overflow || st.underflow {
		return fmt.Errorf("blocked")
	}

	size := (int)(unsafe.Sizeof(value))
	if st.pointer+size > st.size {
		st.overflow = true
		return fmt.Errorf("overflow")
	}

	err = st.mem.PutByte(st.offset+st.pointer, value)
	if err != nil {
		return err
	}

	st.pointer += size
	return nil
}

// GetByte returns a byte relative to the stack-pointer
func (st *Stack) GetByte(offset int) (value byte, err error) {
	if st.overflow || st.underflow {
		return 0, fmt.Errorf("blocked")
	}

	value, err = st.mem.GetByte(st.offset + st.pointer + offset)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// PutByte stores a byte relative to the stack-pointer
func (st *Stack) PutByte(offset int, value byte) (err error) {
	if st.overflow || st.underflow {
		return fmt.Errorf("blocked")
	}

	err = st.mem.PutByte(st.offset+st.pointer+offset, value)
	if err != nil {
		return err
	}

	return nil
}

// PopByte removes a byte from the stack
func (st *Stack) PopByte() (value byte, err error) {
	if st.overflow || st.underflow {
		return 0, fmt.Errorf("blocked")
	}

	size := (int)(unsafe.Sizeof(value))
	if st.pointer-size < 0 {
		st.underflow = true
		return 0, fmt.Errorf("underflow")
	}
	st.pointer -= size

	value, err = st.mem.GetByte(st.offset + st.pointer)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// -- Basic stack functions on ints ---------------------------------------------------------------------------------------------

// PushInt puts an int on the stack
func (st *Stack) PushInt(value int) (err error) {
	if st.overflow || st.underflow {
		return fmt.Errorf("blocked")
	}

	size := (int)(unsafe.Sizeof(value))
	if st.pointer+size > st.size {
		st.overflow = true
		return fmt.Errorf("overflow")
	}

	err = st.mem.PutInt(st.offset+st.pointer, value)
	if err != nil {
		return err
	}

	st.pointer += size
	return nil
}

// GetInt returns an int relative to the stack-pointer
func (st *Stack) GetInt(offset int) (value int, err error) {
	if st.overflow || st.underflow {
		return 0, fmt.Errorf("blocked")
	}

	value, err = st.mem.GetInt(st.offset + st.pointer + offset)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// PutByte stores a byte relative to the stack-pointer
func (st *Stack) PutInt(offset int, value int) (err error) {
	if st.overflow || st.underflow {
		return fmt.Errorf("blocked")
	}

	err = st.mem.PutInt(st.offset+st.pointer+offset, value)
	if err != nil {
		return err
	}

	return nil
}

// PopInt removes an int from the stack
func (st *Stack) PopInt() (value int, err error) {
	if st.overflow || st.underflow {
		return 0, fmt.Errorf("blocked")
	}

	size := (int)(unsafe.Sizeof(value))
	if st.pointer-size < 0 {
		st.underflow = true
		return 0, fmt.Errorf("underflow")
	}

	st.pointer -= size
	value, err = st.mem.GetInt(st.offset + st.pointer)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// -- Basic stack functions on floats -------------------------------------------------------------------------------------------

func (st *Stack) PushFloat(value float64) (err error) {
	if st.overflow || st.underflow {
		return fmt.Errorf("blocked")
	}

	size := (int)(unsafe.Sizeof(value))
	if st.pointer+size > st.size {
		st.overflow = true
		return fmt.Errorf("overflow")
	}

	err = st.mem.PutFloat(st.offset+st.pointer, value)
	if err != nil {
		return err
	}

	st.pointer += size
	return nil
}

// GetFloat returns a float relative to the stack-pointer
func (st *Stack) GetFloat(offset int) (value float64, err error) {
	if st.overflow || st.underflow {
		return 0, fmt.Errorf("blocked")
	}

	value, err = st.mem.GetFloat(st.offset + st.pointer + offset)
	if err != nil {
		return 0, err
	}

	return value, nil
}

// PutFloat stores a float relative to the stack-pointer
func (st *Stack) PutFloat(offset int, value float64) (err error) {
	if st.overflow || st.underflow {
		return fmt.Errorf("blocked")
	}

	err = st.mem.PutFloat(st.offset+st.pointer+offset, value)
	if err != nil {
		return err
	}

	return nil
}

func (st *Stack) PopFloat() (result float64, err error) {
	if st.overflow || st.underflow {
		return 0, fmt.Errorf("blocked")
	}

	size := (int)(unsafe.Sizeof(result))
	if st.pointer-size < 0 {
		st.underflow = true
		return 0, fmt.Errorf("underflow")
	}

	st.pointer -= size
	result, err = st.mem.GetFloat(st.offset + st.pointer)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// -- Support functions part of the stack----------------------------------------------------------------------------------------

func (st *Stack) Show() {
	// chalk styles
	headerStyle := chalk.White.NewStyle().WithBackground(chalk.Blue).WithTextStyle(chalk.Bold)
	errorStyle := chalk.White.NewStyle().WithBackground(chalk.Red).WithTextStyle(chalk.Bold)
	defaultStyle := chalk.White.NewStyle().WithBackground(chalk.Blue)
	pointerStyle := chalk.White.NewStyle().WithBackground(chalk.Blue).WithTextStyle(chalk.Underline)

	// contants than influence the printing
	lineItems := 16
	lineLength := lineItems * 3

	// Stack header
	headerText := "Stack"
	if st.overflow {
		headerText = "Stack (overflow)"
	}

	lineSpaces := lineLength - len(headerText)
	headerText = strings.Repeat(" ", lineSpaces/2) + headerText + strings.Repeat(" ", lineSpaces-lineSpaces/2)
	if st.overflow {
		fmt.Println(errorStyle.Style(headerText))
	} else {
		fmt.Println(headerStyle.Style(headerText))
	}

	// Stack contents
	for i := 0; i < st.size; i++ {
		value, err := st.mem.GetByte(st.offset + i)
		if err != nil {
			return
		}
		cell := fmt.Sprintf("%02X", value)
		if i == st.pointer {
			fmt.Print(pointerStyle.Style(cell))
		} else {
			fmt.Print(defaultStyle.Style(cell))
		}
		fmt.Print(defaultStyle.Style(" "))
		if (i+1)%lineItems == 0 {
			fmt.Println()
		}
	}

	// Empty line at the end
	fmt.Println()
}

func (st *Stack) Check(expectedValue []byte) (err error) {
	// Stack in error state
	if st.Overflow() || st.Underflow() {
		return fmt.Errorf("blocked")
	}

	// Retrieve entire current stack
	stack := make([]byte, st.size)
	for i := 0; i < len(stack); i++ {
		stack[i], err = st.mem.GetByte(st.offset + i)
		if err != nil {
			return err
		}
	}

	// Check stack pointer
	if st.pointer != len(expectedValue) {
		return fmt.Errorf("expected: % X, got % X", expectedValue, stack[:st.pointer])
	}

	// Check content
	for i, v := range expectedValue {
		if stack[i] != v {
			return fmt.Errorf("expected: % X, got % X", expectedValue, stack[:st.pointer])
		}
	}

	return nil
}

func (st *Stack) Underflow() bool {
	return st.underflow
}

func (st *Stack) Overflow() bool {
	return st.overflow
}

// -- Companion functions -------------------------------------------------------------------------------------------------------

func NewStack(mem *Memory, stackSize int) (st *Stack, err error) {
	if mem == nil {
		return nil, fmt.Errorf("missing parameter")
	}
	if stackSize < 0 || stackSize > mem.Size() {
		return nil, fmt.Errorf("illegal stack size")
	}

	st = new(Stack)
	st.mem = mem
	st.offset = mem.Size() - stackSize
	st.size = stackSize
	st.pointer = 0

	return st, nil
}
