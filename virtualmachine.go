package virtualmachine

import (
	"bytes"
	"fmt"
	"log"
	"unsafe"
)

// Operation maps a processor instruction as a function pointer
type Operation func() error

// Virtual Machine models an entirely stack based processor.
type VirtualMachine struct {
	jumpTable [256]Operation
	stack     *Stack
	memory    *Memory

	programPointer int

	logBuffer bytes.Buffer
	logFile   *log.Logger
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
	operant, err := vm.memory.GetByte(vm.programPointer + 1)
	if err != nil {
		vm.addLog("Pushbyte: %v\n", err)
		return err
	}

	err = vm.stack.PushByte(operant)
	if err != nil {
		vm.addLog("Pushbyte %d: %v\n", operant, err)
		return err
	}

	vm.addLog("Pushbyte %d: OK\n", operant)
	vm.programPointer += 2
	return nil
}

// operationPushInt takes the following 8 bytes from memory and pushes them as an int
func (vm *VirtualMachine) operationPushInt() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("PushInt: %v\n", err)
		return err
	}

	err = vm.stack.PushInt(operant)
	if err != nil {
		vm.addLog("PushInt %d: %v\n", operant, err)
		return err
	}

	vm.addLog("PushInt %d: OK\n", operant)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationGetByte takes an address and pushes the byte from that memory-address
func (vm *VirtualMachine) operationGetByte() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("GetByte: %v", err)
		return err
	}

	value, err := vm.memory.GetByte(operant)
	if err != nil {
		vm.addLog("GetByte (%d): %v", operant, err)
		return err
	}

	err = vm.stack.PushByte(value)
	if err != nil {
		vm.addLog("GetByte (%d) -> %d: %v", operant, value, err)
		return err
	}

	vm.addLog("GetByte (%d) -> %d: OK", operant, value)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationGetInt takes an address and pushes the int from that memory-address
func (vm *VirtualMachine) operationGetInt() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("GetInt: %v", err)
		return err
	}

	value, err := vm.memory.GetInt(operant)
	if err != nil {
		vm.addLog("GetInt (%d): %v", operant, err)
		return err
	}

	err = vm.stack.PushInt(value)
	if err != nil {
		vm.addLog("GetInt (%d) -> %d: %v", operant, value, err)
		return err
	}

	vm.addLog("GetInt (%d) -> %d: OK", operant, value)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationPutByte takes an address and pops a byte into that memory-address
func (vm *VirtualMachine) operationPutByte() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("PutByte: %v", err)
		return err
	}

	value, err := vm.stack.PopByte()
	if err != nil {
		vm.addLog("PutByte (%d): %v", operant, err)
		return err
	}

	err = vm.memory.PutByte(operant, value)
	if err != nil {
		vm.addLog("PutByte %d -> (%d): %v", value, operant, err)
		return err
	}

	vm.addLog("PutByte %d -> (%d): OK", value, operant)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationPutInt takes an address and pops an int into that memory-address
func (vm *VirtualMachine) operationPutInt() error {
	operant, err := vm.memory.GetInt(vm.programPointer + 1)
	if err != nil {
		vm.addLog("PutInt: %v", err)
		return err
	}

	value, err := vm.stack.PopInt()
	if err != nil {
		vm.addLog("PutInt (%d): %v", operant, err)
		return err
	}

	err = vm.memory.PutInt(operant, value)
	if err != nil {
		vm.addLog("PutInt %d -> (%d): %v", value, operant, err)
		return err
	}

	vm.addLog("PutInt %d -> (%d): OK", value, operant)
	vm.programPointer += 1 + (int)(unsafe.Sizeof(operant))
	return nil
}

// operationAddByte takes 2 bytes from the stack, adds them pushes the result
func (vm *VirtualMachine) operationAddByte() error {
	operant1, err := vm.stack.PopByte()
	if err != nil {
		vm.addLog("AddByte: %v", err)
		return err
	}

	operant2, err := vm.stack.PopByte()
	if err != nil {
		vm.addLog("AddByte %d: %v", operant1, err)
		return err
	}

	err = vm.stack.PushByte(operant1 + operant2)
	if err != nil {
		vm.addLog("AddByte %d, %d: %v", operant1, operant2, err)
		return err
	}

	vm.addLog("AddByte %d, %d: OK", operant1, operant2)
	vm.programPointer++
	return nil
}

// operationAddInt takes 2 integers from the stack, adds them and pushes the result
func (vm *VirtualMachine) operationAddInt() error {
	operant1, err := vm.stack.PopInt()
	if err != nil {
		vm.addLog("AddInt: %v", err)
		return err
	}

	operant2, err := vm.stack.PopInt()
	if err != nil {
		vm.addLog("AddInt %d: %v", operant1, err)
		return err
	}

	err = vm.stack.PushInt(operant1 + operant2)
	if err != nil {
		vm.addLog("AddInt %d, %d: %v", operant1, operant2, err)
		return err
	}

	vm.addLog("AddInt %d, %d: OK", operant1, operant2)
	vm.programPointer++
	return nil
}

// -- Interface -----------------------------------------------------------------------------------------------------------------

func (vm *VirtualMachine) ShowStack() {
	if vm.stack != nil {
		vm.stack.Show()
	}
}

func (vm *VirtualMachine) ShowMemory() {
	if vm.memory != nil {
		vm.memory.Show(vm.programPointer)
	}
}

func (vm *VirtualMachine) Load(program []byte) error {
	for i, v := range program {
		err := vm.memory.PutByte(i, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Step executes a single instruction and returns if we are ended
func (vm *VirtualMachine) Step() (bool, error) {
	// Get operation
	opCode, err := vm.memory.GetByte(vm.programPointer)
	if err != nil {
		return true, err
	}

	// Check operation
	if opCode == 0x00 {
		return true, nil
	}
	if vm.jumpTable[opCode] == nil {
		return true, fmt.Errorf("Opcode Error")
	}

	// Execute operation
	err = vm.jumpTable[opCode]()
	if err != nil {
		return true, err
	}

	return false, nil
}

// main loop of the virtual machine
func (vm *VirtualMachine) Run() error {

	// Start the logbook, later only in error mode
	err := vm.initLogging()
	if err != nil {
		return err
	}

	// Keep stepping until done
	atEnd, err := vm.Step()
	for !atEnd && err == nil {
		atEnd, err = vm.Step()
	}

	vm.showLog()
	return err
}

// -- Companion functions -------------------------------------------------------------------------------------------------------

func NewVirtualMachine() *VirtualMachine {
	vm := new(VirtualMachine)

	// Build the jumpTable
	vm.jumpTable[0x00] = nil // End
	vm.jumpTable[0x08] = vm.operationPushByte
	vm.jumpTable[0x09] = vm.operationPushInt
	vm.jumpTable[0x10] = vm.operationGetByte
	vm.jumpTable[0x11] = vm.operationGetInt
	vm.jumpTable[0x18] = vm.operationPutByte
	vm.jumpTable[0x19] = vm.operationPutInt
	vm.jumpTable[0x20] = vm.operationAddByte
	vm.jumpTable[0x21] = vm.operationAddInt

	// Build the resources
	vm.stack = NewStack()
	vm.memory = NewMemory()

	return vm
}
