package virtualmachine

import (
	"bytes"
	"fmt"
	"log"
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
		return true, fmt.Errorf("opcode %0x unknown", opCode)
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
	vm.addLog("<start program>")

	// Keep stepping until done
	atEnd, err := vm.Step()
	for !atEnd && err == nil {
		atEnd, err = vm.Step()
	}

	vm.addLog("<end program>")
	vm.showLog()
	return err
}

// -- Companion functions -------------------------------------------------------------------------------------------------------

func NewVirtualMachine(memorySize int, stackSize int) (vm *VirtualMachine, err error) {
	vm = new(VirtualMachine)

	// Build the jumpTable
	vm.jumpTable[0x00] = nil // End
	vm.jumpTable[0x08] = vm.operationPushByte
	vm.jumpTable[0x09] = vm.operationPushInt
	vm.jumpTable[0x0A] = vm.operationPushFloat
	vm.jumpTable[0x0C] = vm.operationPopByte
	vm.jumpTable[0x0D] = vm.operationPopInt
	vm.jumpTable[0x0E] = vm.operationPopFloat
	vm.jumpTable[0x10] = vm.operationGetByte
	vm.jumpTable[0x11] = vm.operationGetInt
	vm.jumpTable[0x12] = vm.operationGetFloat
	vm.jumpTable[0x18] = vm.operationPutByte
	vm.jumpTable[0x19] = vm.operationPutInt
	vm.jumpTable[0x1A] = vm.operationPutFloat
	vm.jumpTable[0x20] = vm.operationGetByteAddress
	vm.jumpTable[0x21] = vm.operationGetIntAddress
	vm.jumpTable[0x22] = vm.operationGetFloatAddress
	vm.jumpTable[0x28] = vm.operationPutByteAddress
	vm.jumpTable[0x29] = vm.operationPutIntAddress
	vm.jumpTable[0x2A] = vm.operationPutFloatAddress
	vm.jumpTable[0x30] = vm.operationGetByteStack
	vm.jumpTable[0x31] = vm.operationGetIntStack
	vm.jumpTable[0x32] = vm.operationGetFloatStack
	vm.jumpTable[0x38] = vm.operationPutByteStack
	vm.jumpTable[0x39] = vm.operationPutIntStack
	vm.jumpTable[0x3A] = vm.operationPutFloatStack
	vm.jumpTable[0x40] = vm.operationAddByte
	vm.jumpTable[0x41] = vm.operationAddInt
	vm.jumpTable[0x42] = vm.operationAddFloat
	vm.jumpTable[0x44] = vm.operationSubByte
	vm.jumpTable[0x45] = vm.operationSubInt
	vm.jumpTable[0x46] = vm.operationSubFloat
	vm.jumpTable[0x48] = vm.operationMulByte
	vm.jumpTable[0x49] = vm.operationMulInt
	vm.jumpTable[0x4A] = vm.operationMulFloat
	vm.jumpTable[0x4C] = vm.operationDivByte
	vm.jumpTable[0x4D] = vm.operationDivInt
	vm.jumpTable[0x4E] = vm.operationDivFloat
	vm.jumpTable[0x60] = vm.operationEqualByte
	vm.jumpTable[0x61] = vm.operationEqualInt
	vm.jumpTable[0x62] = vm.operationEqualFloat
	vm.jumpTable[0x64] = vm.operationUnequalByte
	vm.jumpTable[0x65] = vm.operationUnequalInt
	vm.jumpTable[0x66] = vm.operationUnequalFloat
	vm.jumpTable[0x68] = vm.operationGreaterByte
	vm.jumpTable[0x69] = vm.operationGreaterInt
	vm.jumpTable[0x6A] = vm.operationGreaterFloat
	vm.jumpTable[0x6C] = vm.operationSmallerByte
	vm.jumpTable[0x6D] = vm.operationSmallerInt
	vm.jumpTable[0x6E] = vm.operationSmallerFloat
	vm.jumpTable[0x70] = vm.operationAndByte
	vm.jumpTable[0x71] = vm.operationOrByte
	vm.jumpTable[0x72] = vm.operationNotByte
	vm.jumpTable[0x73] = vm.operationXorByte
	vm.jumpTable[0xE0] = vm.operationRet
	vm.jumpTable[0xE1] = vm.operationJmp
	vm.jumpTable[0xE4] = vm.operationJmpzByte
	vm.jumpTable[0xE5] = vm.operationJmpzInt
	vm.jumpTable[0xE6] = vm.operationJmpzFloat
	vm.jumpTable[0xF8] = vm.operationCall
	vm.jumpTable[0xF9] = vm.operationCallAddress

	// Build the resources
	vm.memory = NewMemory(memorySize)
	vm.stack, err = NewStack(vm.memory, stackSize)
	if err != nil {
		return nil, err
	}

	return vm, nil
}
