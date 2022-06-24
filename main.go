// Starting file for virtual machine
package main

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
