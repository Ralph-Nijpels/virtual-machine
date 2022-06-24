package main

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/ttacon/chalk"
)

const MEMORY_SIZE int = 256

type Memory struct {
	memory [MEMORY_SIZE]byte
}

// -- Basic memory functions on bytes -------------------------------------------------------------------------------------------

func (mem *Memory) GetByte(address int) (byte, error) {
	if address < 0 || address >= MEMORY_SIZE {
		return 0, fmt.Errorf("Memory Error")
	}

	return mem.memory[address], nil
}

func (mem *Memory) PutByte(address int, value byte) error {
	if address < 0 || address >= MEMORY_SIZE {
		return fmt.Errorf("Memory Error")
	}

	mem.memory[address] = value
	return nil
}

// -- Basic memory functions on ints --------------------------------------------------------------------------------------------

// GetInt fetches an Int
func (mem *Memory) GetInt(address int) (int, error) {
	var result int

	if address < 0 || address+(int)(unsafe.Sizeof(result)) >= MEMORY_SIZE {
		return 0, fmt.Errorf("Memory error")
	}

	result = *(*int)(unsafe.Pointer(&mem.memory[address]))
	return result, nil
}

// PutInt stores an Int
func (mem *Memory) PutInt(address int, value int) error {
	if address < 0 || address+(int)(unsafe.Sizeof(value)) >= MEMORY_SIZE {
		return fmt.Errorf("Memory error")
	}

	*(*int)(unsafe.Pointer(&mem.memory[address])) = value
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

// -- Companion functions -------------------------------------------------------------------------------------------------------

func NewMemory() *Memory {
	return new(Memory)
}
