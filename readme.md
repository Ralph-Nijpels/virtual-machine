# Simple virtual machine
To train my low-level skills a bit, I've choosen to implement a simple virtual machine

# Plan for the Opcodes
0x08 push-byte   nn     // pushes a constant byte value on the stack
0x09 push-int    nn     // pushes a contant integer value on the stack

0x10 get-byte   (nn)    // pushes a byte from memory on the stack
0x11 get-int    (nn)    // pushes an int from memory on the stack

0x18 put-byte   (nn)    // stores a byte from stack into memory
0x19 put-int    (nn)    // stores an int from stack into memory

0x20 add-byte           // adds the two topmost bytes on stack
0x21 add-int            // adds the two topmost ints on stack

sub-int
mul-int

div-int

and-int

or-int


