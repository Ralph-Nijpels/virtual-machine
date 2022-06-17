// Starting file for virtual machine
package main

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/ttacon/chalk"
)

// Virtual Machine models an entirely stack based processor.
const STACK_SIZE int = 64
const MEMORY_SIZE int = 256

type VirtualMachine struct {
	stack          [STACK_SIZE]byte
	stackPointer   int
	stackOverflow  bool
	stackUnderflow bool
	memory         [MEMORY_SIZE]byte
	programPointer int
}

// Basic stack manipulation functions on bytes
func (vm *VirtualMachine) pushByte(v byte) error {
	if vm.stackOverflow || vm.stackUnderflow {
		return fmt.Errorf("StackEror")
	}

	if vm.stackPointer+1 > STACK_SIZE {
		vm.stackOverflow = true
		return fmt.Errorf("StackEror")
	}

	vm.stack[vm.stackPointer] = v
	vm.stackPointer++

	return nil
}

func (vm *VirtualMachine) popByte() (byte, error) {
	if vm.stackOverflow || vm.stackUnderflow {
		return 0, fmt.Errorf("StackEror")
	}

	if vm.stackPointer == 0 {
		vm.stackUnderflow = true
		return 0, fmt.Errorf("StackEror")
	}

	vm.stackPointer--
	return vm.stack[vm.stackPointer], nil
}

// basic stack functions on ints
func (vm *VirtualMachine) pushInt(v int) error {
	if vm.stackOverflow || vm.stackUnderflow {
		return fmt.Errorf("StackEror")
	}

	size := (int)(unsafe.Sizeof(v))
	if vm.stackPointer+size > STACK_SIZE {
		vm.stackOverflow = true
		return fmt.Errorf("StackEror")
	}

	address := unsafe.Pointer(&v)
	for i := 0; i < size; i++ {
		b := *(*byte)(unsafe.Pointer(uintptr(address) + uintptr(i)))
		vm.stack[vm.stackPointer+i] = b
	}

	vm.stackPointer += size

	return nil
}

func (vm *VirtualMachine) popInt() (int, error) {
	var result int

	if vm.stackOverflow || vm.stackUnderflow {
		return 0, fmt.Errorf("StackEror")
	}

	size := (int)(unsafe.Sizeof(result))
	if vm.stackPointer-size < 0 {
		vm.stackUnderflow = true
		return 0, fmt.Errorf("StackEror")
	}

	address := unsafe.Pointer(&(vm.stack[vm.stackPointer-size]))
	result = *(*int)(address)

	vm.stackPointer -= size

	return result, nil
}

// Function to show the stack
func (vm *VirtualMachine) showStack() {
	// chalk styles
	headerStyle := chalk.White.NewStyle().WithBackground(chalk.Blue).WithTextStyle(chalk.Bold)
	errorStyle := chalk.White.NewStyle().WithBackground(chalk.Red).WithTextStyle(chalk.Bold)
	defaultStyle := chalk.White.NewStyle().WithBackground(chalk.Blue)
	pointerStyle := chalk.White.NewStyle().WithBackground(chalk.Blue).WithTextStyle(chalk.Underline)
	lineItems := 16
	lineLength := lineItems * 3

	// Stack header
	headerText := "Stack"
	if vm.stackOverflow {
		headerText = "Stack (overflow)"
	}

	lineSpaces := lineLength - len(headerText)
	headerText = strings.Repeat(" ", lineSpaces/2) + headerText + strings.Repeat(" ", lineSpaces-lineSpaces/2)
	if vm.stackOverflow {
		fmt.Println(errorStyle.Style(headerText))
	} else {
		fmt.Println(headerStyle.Style(headerText))
	}

	// Stack contents
	for i, v := range vm.stack {
		value := fmt.Sprintf("%02X", v)
		if i == vm.stackPointer {
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

// Simple functions for the VM
func (vm *VirtualMachine) operationAddInt() error {
	v1, err := vm.popInt()
	if err != nil {
		return err
	}

	v2, err := vm.popInt()
	if err != nil {
		return err
	}

	return vm.pushInt(v1 + v2)
}

// main loop of the virtual machine
func (vm *VirtualMachine) Run() error {

}

func main() {
	vm := new(VirtualMachine)

	vm.pushInt(12)
	vm.pushInt(12)
	vm.pushInt(-6)

	err := vm.operationAddInt()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	err = vm.operationAddInt()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	v, err := vm.popInt()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("result: %d\n", v)
	vm.showStack()

}
