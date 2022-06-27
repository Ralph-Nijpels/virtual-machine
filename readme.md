# Simple virtual machine
To train my low-level skills a bit, I've choosen to implement a virtual machine

# Plan for the Opcodes
| Opcode | Mnemonic        | Description                                 |
|-------:|:----------------|:--------------------------------------------|
| 0x08   | push-byte   nn  | pushes a constant byte value on the stack   |
| 0x09   | push-int    nn  | pushes a contant integer value on the stack |
| 0x0A   | push-float  nn  | pushes a constant float value on the stack  |
|        |                 |                                             |
| 0x10   | get-byte   (nn) | pushes a byte from memory on the stack      |
| 0x11   | get-int    (nn) | pushes an int from memory on the stack      |
| 0x12   | get-float  (nn) | pushes a float from memory on the stack     |
|        |                 |                                             |
| 0x18   | put-byte   (nn) | stores a byte from stack into memory        |
| 0x19   | put-int    (nn) | stores an int from stack into memory        |
| 0x1A   | put-float  (nn) | stores a float from stack into memory       |
|        |                 |                                             |
| 0x20   | add-byte        | adds the two topmost bytes on stack         |
| 0x21   | add-int         | adds the two topmost ints on stack          |
| 0x22   | add-float       | adds the two topmost floats on the stack    |
|        |                 |                                             |                  
|        | sub-int         |                                             |
|        |                 |                                             |
|        | mul-int         |                                             |
|        |                 |                                             |
|        | div-int         |                                             |
|        |                 |                                             |
|        | equal-int       |                                             |
|        |                 |                                             |
|        | unequal-int     |                                             |
|        |                 |                                             |
|        | greater-int     |                                             |
|        |                 |                                             |
|        | smaller-int     |                                             |
|        |                 |                                             |
|        | and-int         |                                             |
|        |                 |                                             |
|        | or-int          |                                             |
|        |                 |                                             |


