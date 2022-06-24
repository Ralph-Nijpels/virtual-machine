package main

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/ttacon/chalk"
)

const STACK_SIZE int = 64

type Stack struct {
	stack          [STACK_SIZE]byte
	stackPointer   int
	stackOverflow  bool
	stackUnderflow bool
}

// -- Basic stack manipulation functions on bytes -------------------------------------------------------------------------------

// PushByte puts a byte on the stack
func (st *Stack) PushByte(v byte) error {
	if st.stackOverflow || st.stackUnderflow {
		return fmt.Errorf("StackEror")
	}

	if st.stackPointer+1 > STACK_SIZE {
		st.stackOverflow = true
		return fmt.Errorf("StackEror")
	}

	st.stack[st.stackPointer] = v
	st.stackPointer++

	return nil
}

// PopByte removes a byte from the stack
func (st *Stack) PopByte() (byte, error) {
	if st.stackOverflow || st.stackUnderflow {
		return 0, fmt.Errorf("StackEror")
	}

	if st.stackPointer == 0 {
		st.stackUnderflow = true
		return 0, fmt.Errorf("StackEror")
	}

	st.stackPointer--
	return st.stack[st.stackPointer], nil
}

// -- Basic stack functions on ints ---------------------------------------------------------------------------------------------

// PushInt puts an int on the stack
func (st *Stack) PushInt(v int) error {
	if st.stackOverflow || st.stackUnderflow {
		return fmt.Errorf("StackEror")
	}

	size := (int)(unsafe.Sizeof(v))
	if st.stackPointer+size > STACK_SIZE {
		st.stackOverflow = true
		return fmt.Errorf("StackEror")
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
		return 0, fmt.Errorf("StackEror")
	}

	size := (int)(unsafe.Sizeof(result))
	if st.stackPointer-size < 0 {
		st.stackUnderflow = true
		return 0, fmt.Errorf("StackEror")
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

func (st *Stack) Underflow() bool {
	return st.stackUnderflow
}

func (st *Stack) Overflow() bool {
	return st.stackOverflow
}

// -- Companion functions -------------------------------------------------------------------------------------------------------

func NewStack() *Stack {
	return new(Stack)
}
