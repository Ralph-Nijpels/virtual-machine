// Starting file for virtual machine
package main

import (
	"bytes"
	"fmt"
	"log"
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

	logBuffer bytes.Buffer
	logFile   *log.Logger
}

// -- STACK SECTION ------------------------------------------------------------------------------

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

// -- MEMORY SECTION ----------------------------------------------------------------------------------------------------------------

// Function to show memory
func (vm *VirtualMachine) showMemory() {
	// chalk styles
	headerStyle := chalk.White.NewStyle().WithBackground(chalk.Green).WithTextStyle(chalk.Bold)
	defaultStyle := chalk.White.NewStyle().WithBackground(chalk.Green)
	pointerStyle := chalk.White.NewStyle().WithBackground(chalk.Green).WithTextStyle(chalk.Underline)
	lineItems := 16
	lineLength := lineItems * 3

	// Memory header
	headerText := "Memory"
	lineSpaces := lineLength - len(headerText)
	headerText = strings.Repeat(" ", lineSpaces/2) + headerText + strings.Repeat(" ", lineSpaces-lineSpaces/2)
	fmt.Println(headerStyle.Style(headerText))

	// Memory contents
	for i, v := range vm.memory {
		value := fmt.Sprintf("%02X", v)
		if i == vm.programPointer {
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

// -- LOGGING SECTION --------------------------------------------------------------------------------------
// Initial version alway logs and simply to a buffer. Later we'll add more options

func (vm *VirtualMachine) initLogging() error {
	vm.logFile = log.New(&(vm.logBuffer), "", log.LstdFlags|log.Ltime|log.Lmicroseconds|log.LUTC)
	return nil
}

func (vm *VirtualMachine) addLog(format string, v ...interface{}) error {
	if vm.logFile != nil {
		vm.logFile.Printf(format, v...)
	}
	return nil
}

func (vm *VirtualMachine) showLog() error {
	if vm.logFile != nil {
		fmt.Print(&vm.logBuffer)
	}
	return nil
}

// -- VIRTUAL MACHINE SECTION ------------------------------------------------------------------------------

// operationPushByte takes the following bytes and pushes it as a byte
func (vm *VirtualMachine) operationPushByte() error {
	operant := vm.memory[vm.programPointer+1]
	err := vm.pushByte(operant)

	if err != nil {
		vm.addLog("Pushbyte %d --> %v\n", operant, err)
		return err
	}

	vm.addLog("Pushbyte %d --> OK\n", operant)
	vm.programPointer += 2
	return nil
}

// operationPushInt takes the following 8 bytes from memory and pushes them as an int
func (vm *VirtualMachine) operationPushInt() error {
	operant := *(*int)(unsafe.Pointer(&vm.memory[vm.programPointer+1]))
	err := vm.pushInt(operant)

	if err != nil {
		vm.addLog("PushInt %d --> %v\n", operant, err)
		return err
	}

	vm.addLog("PushInt %d --> OK\n", operant)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationAddByte takes 2 bytes from the stack, adds them pushes the result
func (vm *VirtualMachine) operationAddByte() error {
	operant1, err := vm.popByte()
	if err != nil {
		vm.addLog("AddByte ??, ?? --> %v", err)
		return err
	}

	operant2, err := vm.popByte()
	if err != nil {
		vm.addLog("AddByte %d, ?? --> %v", operant1, err)
		return err
	}

	err = vm.pushByte(operant1 + operant2)
	if err != nil {
		vm.addLog("AddByte %d, %d --> %v", operant1, operant2, err)
		return err
	}

	vm.addLog("AddByte %d, %d --> OK", operant1, operant2)
	vm.programPointer++
	return nil
}

// operationAddInt takes 2 integers from the stack, adds them and pushes the result
func (vm *VirtualMachine) operationAddInt() error {
	operant1, err := vm.popInt()
	if err != nil {
		vm.addLog("AddInt ??, ?? --> %v", err)
		return err
	}

	operant2, err := vm.popInt()
	if err != nil {
		vm.addLog("AddInt %d, ?? --> %v", operant1, err)
		return err
	}

	err = vm.pushInt(operant1 + operant2)
	if err != nil {
		vm.addLog("AddInt %d, %d --> %v", operant1, operant2, err)
		return err
	}

	vm.addLog("AddInt %d, %d --> OK", operant1, operant2)
	vm.programPointer++
	return nil
}

// simple function to get something into memory, later to be replaced to
// obtain the byte code from more interesting places
func (vm *VirtualMachine) Load(program []byte) error {
	for i, v := range program {
		vm.memory[i] = v
	}
	return nil
}

// main loop of the virtual machine
func (vm *VirtualMachine) Run() error {

	// The jumptable ensures a flexible way of adding functions
	type Operation func() error
	var jumpTable [256]Operation

	// Explicitly link specific functions to specific opCodes
	jumpTable[0x00] = nil // End
	jumpTable[0x08] = vm.operationPushByte
	jumpTable[0x09] = vm.operationPushInt
	jumpTable[0x10] = vm.operationAddByte
	jumpTable[0x11] = vm.operationAddInt

	// Start the logbook, later only in error mode
	err := vm.initLogging()
	if err != nil {
		return err
	}

	opCode := vm.memory[vm.programPointer]
	if opCode != 0x00 && jumpTable[opCode] == nil {
		err = fmt.Errorf("unknown opCode")
	}
	for opCode != 0x00 && err == nil {
		err = jumpTable[opCode]()
		opCode = vm.memory[vm.programPointer]
		if opCode != 0x00 && jumpTable[opCode] == nil {
			err = fmt.Errorf("unknown opCode")
		}
	}

	vm.showLog()

	return err
}

func main() {
	vm := new(VirtualMachine)

	program := [...]byte{
		0x09,                                           // PushInt
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // int 12
		0x09,                                           // PushInt
		0x06, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // int 6
		0x11, // AddInt
		0x00} // End

	vm.Load(program[:])
	vm.showMemory()
	vm.Run()
	vm.showStack()
}
