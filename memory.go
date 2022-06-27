package virtualmachine

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/ttacon/chalk"
)

type Memory struct {
	memory []byte
}

// -- Basic memory functions on bytes -------------------------------------------------------------------------------------------

func (mem *Memory) GetByte(address int) (byte, error) {
	if address < 0 || address >= len(mem.memory) {
		return 0, fmt.Errorf("Memory Error")
	}

	return mem.memory[address], nil
}

func (mem *Memory) PutByte(address int, value byte) error {
	if address < 0 || address >= len(mem.memory) {
		return fmt.Errorf("Memory Error")
	}

	mem.memory[address] = value
	return nil
}

// -- Basic memory functions on ints --------------------------------------------------------------------------------------------

// GetInt fetches an Int
func (mem *Memory) GetInt(address int) (int, error) {
	var result int

	if address < 0 || address+(int)(unsafe.Sizeof(result)) > len(mem.memory) {
		return 0, fmt.Errorf("Memory error")
	}

	result = *(*int)(unsafe.Pointer(&mem.memory[address]))
	return result, nil
}

// PutInt stores an Int
func (mem *Memory) PutInt(address int, value int) error {
	if address < 0 || address+(int)(unsafe.Sizeof(value)) > len(mem.memory) {
		return fmt.Errorf("Memory error")
	}

	*(*int)(unsafe.Pointer(&mem.memory[address])) = value
	return nil
}

// -- Basic memory functions on Float64 -----------------------------------------------------------------------------------------

func (mem *Memory) GetFloat(address int) (float64, error) {
	var result float64

	if address < 0 || address+(int)(unsafe.Sizeof(result)) > len(mem.memory) {
		return 0, fmt.Errorf("Memory error")
	}

	result = *(*float64)(unsafe.Pointer(&mem.memory[address]))
	return result, nil
}

func (mem *Memory) PutFloat(address int, value float64) error {
	if address < 0 || address+(int)(unsafe.Sizeof(value)) > len(mem.memory) {
		return fmt.Errorf("Memory error")
	}

	*(*float64)(unsafe.Pointer(&mem.memory[address])) = value
	return nil
}

// -- Support functions ---------------------------------------------------------------------------------------------------------

// Show displays the content of the memory in hex
func (mem *Memory) Show(programPointer int) {
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
	for i, v := range mem.memory {
		value := fmt.Sprintf("%02X", v)
		if i == programPointer {
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

// Check compares (a part of) memory with some expected value, used for testing
func (mem *Memory) Check(expectedValue []byte) error {
	for i, v := range expectedValue {
		if mem.memory[i] != v {
			return fmt.Errorf("expected %X, got %X", expectedValue, mem.memory[:len(expectedValue)])
		}
	}

	return nil
}

// -- Companion functions -------------------------------------------------------------------------------------------------------

func NewMemory(memorySize int) *Memory {
	memory := new(Memory)

	memory.memory = make([]byte, memorySize)

	return memory
}
