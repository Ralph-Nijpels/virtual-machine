package virtualmachine

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/ttacon/chalk"
)

type Stack struct {
	stack          []byte
	stackPointer   int
	stackOverflow  bool
	stackUnderflow bool
}

// -- Basic stack manipulation functions on bytes -------------------------------------------------------------------------------

// PushByte puts a byte on the stack
func (st *Stack) PushByte(v byte) error {
	if st.stackOverflow || st.stackUnderflow {
		return fmt.Errorf("Blocked")
	}

	if st.stackPointer+1 > len(st.stack) {
		st.stackOverflow = true
		return fmt.Errorf("Overflow")
	}

	st.stack[st.stackPointer] = v
	st.stackPointer++

	return nil
}

// PopByte removes a byte from the stack
func (st *Stack) PopByte() (byte, error) {
	if st.stackOverflow || st.stackUnderflow {
		return 0, fmt.Errorf("Blocked")
	}

	if st.stackPointer == 0 {
		st.stackUnderflow = true
		return 0, fmt.Errorf("Underflow")
	}

	st.stackPointer--
	return st.stack[st.stackPointer], nil
}

// -- Basic stack functions on ints ---------------------------------------------------------------------------------------------

// PushInt puts an int on the stack
func (st *Stack) PushInt(v int) error {
	if st.stackOverflow || st.stackUnderflow {
		return fmt.Errorf("Blocked")
	}

	size := (int)(unsafe.Sizeof(v))
	if st.stackPointer+size > len(st.stack) {
		st.stackOverflow = true
		return fmt.Errorf("Overflow")
	}

	address := unsafe.Pointer(&v)
	for i := 0; i < size; i++ {
		b := *(*byte)(unsafe.Pointer(uintptr(address) + uintptr(i)))
		st.stack[st.stackPointer+i] = b
	}

	st.stackPointer += size

	return nil
}

// PopInt removes an int from the stack
func (st *Stack) PopInt() (int, error) {
	var result int

	if st.stackOverflow || st.stackUnderflow {
		return 0, fmt.Errorf("Blocked")
	}

	size := (int)(unsafe.Sizeof(result))
	if st.stackPointer-size < 0 {
		st.stackUnderflow = true
		return 0, fmt.Errorf("Underflow")
	}

	address := unsafe.Pointer(&(st.stack[st.stackPointer-size]))
	result = *(*int)(address)

	st.stackPointer -= size

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
	if st.stackOverflow {
		headerText = "Stack (overflow)"
	}

	lineSpaces := lineLength - len(headerText)
	headerText = strings.Repeat(" ", lineSpaces/2) + headerText + strings.Repeat(" ", lineSpaces-lineSpaces/2)
	if st.stackOverflow {
		fmt.Println(errorStyle.Style(headerText))
	} else {
		fmt.Println(headerStyle.Style(headerText))
	}

	// Stack contents
	for i, v := range st.stack {
		value := fmt.Sprintf("%02X", v)
		if i == st.stackPointer {
			fmt.Print(pointerStyle.Style(value))
		} else {
			fmt.Print(defaultStyle.Style(value))
		}
		fmt.Print(defaultStyle.Style(" "))
		if (i+1)%lineItems == 0 {
			fmt.Println()
		}
	}

	// Empty line at the end
	fmt.Println()
}

func (st *Stack) Check(expectedValue []byte) error {
	// Stack in error state
	if st.Overflow() || st.Underflow() {
		return fmt.Errorf("Blocked")
	}

	// Check stack pointer
	if st.stackPointer != len(expectedValue) {
		return fmt.Errorf("Stack pointer")
	}

	// Check content
	for i, v := range expectedValue {
		if st.stack[i] != v {
			return fmt.Errorf("Stack content")
		}
	}

	return nil
}

func (st *Stack) Underflow() bool {
	return st.stackUnderflow
}

func (st *Stack) Overflow() bool {
	return st.stackOverflow
}

// -- Companion functions -------------------------------------------------------------------------------------------------------

func NewStack(stackSize int) *Stack {
	stack := new(Stack)

	stack.stack = make([]byte, stackSize)

	return stack
}
