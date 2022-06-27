package virtualmachine

import "testing"

// -- Tests ---------------------------------------------------------------------------------------------------------------------

func TestPushByte(t *testing.T) {

	program := [...]byte{
		0x08, // PushByte
		0x0C, // byte 12
		0x00} // End

	stack := [...]byte{
		0x0C}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetByte(t *testing.T) {
	program := [...]byte{
		0x10,                                           // GetByte
		0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // address
		0x00,
		0x91} // Value

	stack := [...]byte{
		0x91}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPutByte(t *testing.T) {
	program := [...]byte{
		0x08,                                           // PushByte
		0xFE,                                           // Operant
		0x18,                                           // Put Byte
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Operant
		0x00, // End Program
		0x00}

	stack := [...]byte{}

	memory := [...]byte{
		0x08,                                           // see program
		0xFE,                                           // ..
		0x18,                                           // ..
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ..
		0x00, // ..
		0xFE} // Changed!

	err := runProgram(program[:], stack[:], memory[:])
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAddByte(t *testing.T) {
	program := [...]byte{
		0x08, // PushByte
		0x04, // Value
		0x08, // PushByte
		0x06, // Value
		0x20, // AddByte
		0x00} // EndProgram

	stack := [...]byte{
		0x0A}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
