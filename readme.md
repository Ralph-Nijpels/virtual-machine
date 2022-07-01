# Simple virtual machine
To train my low-level skills a bit, I've choosen to implement a virtual machine

# Refactoring to-do
- [ ] Implement short rather than byte opcodes. We are going to run out of opcodes if we want to implement strings
- [x] Implement get-xxx / put-xxx using an address from stack. Needed to allow for calculated addresses if we want to implement strings and arrays
- [x] Implement get-xxx / put-xxx using an address relative to the stack-pointer. Needed to create stack-frames to implement call/return 

# Plan for the Opcodes
| Done | Opcode | Mnemonic        | Description                                                                               |
|:----:|-------:|:----------------|:----------------------------------------------------------------------------------------- |
| [x]  | 0x08   | push-byte   nn  | pushes a constant byte value on the stack                                                 |
| [x]  | 0x09   | push-int    nn  | pushes a contant integer value on the stack                                               |
| [x]  | 0x0A   | push-float  nn  | pushes a constant float value on the stack                                                |
|      |        |                 |                                                                                           |
| [x]  | 0x10   | get-byte        | pops an address from stack, retrieves a byte from this address and push it onto the stack |
| [x]  | 0x11   | get-int         | pops an address from stack, retrieves a byte from this address and push it onto the stack |
| [x]  | 0x12   | get-float       | pops an address from stack, retrieves a byte from this address and push it onto the stack |
|      |        |                 |                                                                                           |
| [x]  | 0x18   | put-byte        | pops an address from stack, pops a byte from stack and stores it in memory                |
| [x]  | 0x19   | put-int         | pops an address from stack, pops an int from stack and stores it in memory                |
| [x]  | 0x1A   | put-float       | pops an address from stack, pops a float from stack and stores it in memory               |
|      |        |                 |                                                                                           |
| [x]  | 0x20   | get-byte   (nn) | pushes a byte from memory on the stack                                                    |
| [x]  | 0x21   | get-int    (nn) | pushes an int from memory on the stack                                                    |
| [x]  | 0x22   | get-float  (nn) | pushes a float from memory on the stack                                                   |
|      |        |                 |                                                                                           |
| [x]  | 0x28   | put-byte   (nn) | stores a byte from stack into memory                                                      |
| [x]  | 0x29   | put-int    (nn) | stores an int from stack into memory                                                      |
| [x]  | 0x2A   | put-float  (nn) | stores a float from stack into memory                                                     |
|      |        |                 |                                                                                           |
| [x]  | 0x30   | get-byte   {nn} | pushes a byte from an address relative to the stackpointer on top of the stack            |
| [x]  | 0x31   | get-int    {nn} | pushes an int from an address relative to the stackpointer on top of the stack            |
| [x]  | 0x32   | get-float  {nn} | pushes a float from an address relative to the stackpointer on top of the stack           |
|      |        |                 |                                                                                           |
| [x]  | 0x38   | put-byte   {nn} | pops a byte from the stack and stores it in address relative to the stackpointer          |
| [x]  | 0x39   | put-int    {nn} | pops an int from the stack and stores it in address relative to the stackpointer          |
| [x]  | 0x3A   | put-float  {nn} | pops a float from the stack and stores it in address relative to the stackpointer         |
|      |        |                 |                                                                                           |
| [x]  | 0x40   | add-byte        | adds the two topmost bytes on stack                                                       |
| [x]  | 0x41   | add-int         | adds the two topmost ints on stack                                                        |
| [x]  | 0x42   | add-float       | adds the two topmost floats on the stack                                                  |
|      |        |                 |                                                                                           |
| [x]  | 0x44   | sub-byte        | subtracts the two topmost bytes on stack                                                  |
| [x]  | 0x45   | sub-int         | subtracts the two topmost ints on stack                                                   |
| [x]  | 0x46   | sub-float       | subtracts the two topmost floats on stack                                                 |
|      |        |                 |                                                                                           |
| [x]  | 0x48   | mul-byte        | multiplies the two topmost bytes on stack                                                 |
| [x]  | 0x49   | mul-int         | multiplies the two topmost ints on stack                                                  |
| [x]  | 0x4A   | mul-float       | multiplies the two topmost floats on stack                                                |
|      |        |                 |                                                                                           |
| [x]  | 0x4C   | div-byte        | divides the two topmost bytes on stack                                                    |
| [x]  | 0x4D   | div-int         | divides the two topmost ints on stack                                                     |
| [x]  | 0x4E   | div-float       | divides the two topmost floats on stack                                                   |
|      |        |                 |                                                                                           |
|      |        |                 | some intentional open space in the opcode table                                           |
|      |        |                 |                                                                                           |
| [ ]  | 0x60   | equal-byte      | compares the topmost two bytes on stack, pushes byte(-1) if equal and 0 otherwise         |
| [ ]  | 0x61   | equal-int       | compares the topmost two ints on stack, pushes byte(-1) if equal and 0 otherwise          |
| [ ]  | 0x62   | equal-float     | compares the topmost two floats on stack, pushes byte(-1) if equal and 0 otherwise        |
|      |        |                 |                                                                                           |
| [ ]  | 0x64   | unequal-byte    | compares the topmost two bytes on stack, pushes byte(-1) if unequal and 0 otherwise       |
| [ ]  | 0x65   | unequal-int     | compares the topmost two ints on stack, pushes byte(-1) if unequal and 0 otherwise        |
| [ ]  | 0x66   | unequal-float   | compares the topmost two floats on stack, pushes byte(-1) if unequal and 0 otherwise      |
|      |        |                 |                                                                                           |
| [ ]  | 0x68   | greater-byte    |                                                                                           |
| [ ]  | 0x69   | greater-int     |                                                                                           |
| [ ]  | 0x6A   | greater-float   |                                                                                           |
|      |        |                 |                                                                                           |
| [ ]  | 0x6C   | smaller-byte    |                                                                                           |
| [ ]  | 0x6D   | smaller-int     |                                                                                           |
| [ ]  | 0x6E   | smaller-Float   |                                                                                           |
|      |        |                 |                                                                                           |
| [ ]  | 0x70   | and-byte        |                                                                                           |
| [ ]  | 0x71   | and-int         |                                                                                           |
|      |        |                 |                                                                                           |
| [ ]  | 0x74   | or-byte         |                                                                                           |
| [ ]  | 0x75   | or-int          |                                                                                           |
|      |        |                 |                                                                                           |
| [ ]  | 0x78   | not-byte        |                                                                                           |
| [ ]  | 0x79   | not-int         |                                                                                           |
|      |        |                 |                                                                                           |
| [ ]  | 0x7C   | xor-byte        |                                                                                           |
| [ ]  | 0x7D   | xor-int         |                                                                                           |
|      |        |                 |                                                                                           |
| [ ]  | 0xF0   | jmp             | address from stack (equivalent to ret)                                                    |
| [ ]  | 0xF1   | jmp        (nn) | direct address                                                                            |
|      |        |                 |                                                                                           |
| [ ]  | 0xF4   | jmpz-byte       |                                                                                           |
| [ ]  | 0xF5   | jmpz-int        |                                                                                           |
| [ ]  | 0xF6   | jmpz-byte  (nn) |                                                                                           |
| [ ]  | 0xF7   | jmpz-int   (nn) |                                                                                           |
|      |        |                 |                                                                                           |
| [ ]  | 0xF8   | jmpnz-byte      |                                                                                           |
| [ ]  | 0xF9   | jmpnz-int       |                                                                                           |
| [ ]  | 0xFA   | jmpnz-byte (nn) |                                                                                           |
| [ ]  | 0xFB   | jmpnz-int  (nn) |                                                                                           |
|      |        |                 |                                                                                           |
| [ ]  | 0xFC   | call            | address from stack                                                                        |
| [ ]  | 0xFD   | call       (nn) |                                                                                           |


