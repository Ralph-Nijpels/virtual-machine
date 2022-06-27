package virtualmachine

import "testing"

func TestPushFloat(t *testing.T) {
	program := [...]byte{
		0x0A,                                           // PushFloat
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCA, 0xEF, // random float
		0x00} // End

	stack := [...]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xCA, 0xEF}

	err := runProgram(program[:], stack[:], nil)
	if err != nil {
		t.Errorf(err.Error())
	}
}
