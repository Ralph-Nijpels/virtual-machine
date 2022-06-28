# Simple virtual machine
To train my low-level skills a bit, I've choosen to implement a virtual machine

# Refactoring to-do
> Implement short rather than byte opcodes. We are going to run out of opcodes if we want to implement strings
> Implement get-xxx / put-xxx using an address from stack. Needed to allow for calculated addresses if we want to implement strings and arrays
> Implement get-xxx / put-xxx using an address relative to the stack-pointer. Needed to create stack-frames to implement call/return 

# Plan for the Opcodes
| Opcode | Mnemonic        | Description                                                                               |
|-------:|:----------------|:----------------------------------------------------------------------------------------- |
| 0x08   | push-byte   nn  | pushes a constant byte value on the stack                                                 |
| 0x09   | push-int    nn  | pushes a contant integer value on the stack                                               |
| 0x0A   | push-float  nn  | pushes a constant float value on the stack                                                |
|        |                 |                                                                                           |
| **To be implemented**                                                                                                |
| 0x10   | get-byte        | pops an address from stack, retrieves a byte from this address and push it onto the stack |
| 0x11   | get-int         | pops an address from stack, retrieves a byte from this address and push it onto the stack |
| 0x12   | get-float       | pops an address from stack, retrieves a byte from this address and push it onto the stack |
|        |                 |                                                                                           |
| 0x18   | put-byte        | pops an address from stack, pops a byte from stack and stores it in memory                |
| 0x18   | put-byte        | pops an address from stack, pops a byte from stack and stores it in memory                |
| 0x18   | put-byte        | pops an address from stack, pops a byte from stack and stores it in memory                |
|        |                 |                                                                                           |
| 0x20   | get-byte   (nn) | pushes a byte from memory on the stack                                                    |
| 0x21   | get-int    (nn) | pushes an int from memory on the stack                                                    |
| 0x22   | get-float  (nn) | pushes a float from memory on the stack                                                   |
|        |                 |                                                                                           |
| 0x28   | put-byte   (nn) | stores a byte from stack into memory                                                      |
| 0x29   | put-int    (nn) | stores an int from stack into memory                                                      |
| 0x2A   | put-float  (nn) | stores a float from stack into memory                                                     |
|        |                 |                                                                                           |
| 0x30   | get-byte   {nn} | pushes a byte from an address relative to the stackpointer on top of the stack            |
| 0x31   | get-int    {nn} | pushes an int from an address relative to the stackpointer on top of the stack            |
| 0x32   | get-float  {nn} | pushes a float from an address relative to the stackpointer on top of the stack           |
|        |                 |                                                                                           |
| 0x38   | put-byte   {nn} | pops a byte from the stack and stores it in address relative to the stackpointer          |
| 0x39   | put-int    {nn} | pops an int from the stack and stores it in address relative to the stackpointer          |
| 0x3A   | put-float  {nn} | pops a float from the stack and stores it in address relative to the stackpointer         |
|        |                 |                                                                                           |
| 0x40   | add-byte        | adds the two topmost bytes on stack                                                       |
| 0x41   | add-int         | adds the two topmost ints on stack                                                        |
| 0x42   | add-float       | adds the two topmost floats on the stack                                                  |
|        |                 |                                                                                           |
| 0x48   | sub-byte        | subtracts the two topmost bytes on stack                                                  |
| 0x49   | sub-int         | subtracts the two topmost ints on stack                                                   |
| 0x4A   | sub-float       | subtracts the two topmost floats on stack                                                 |
|        |                 |                                                                                           |
|        | mul-int         |                                                                                           |
|        |                 |                                                                                           |
|        | div-int         |                                                                                           |
|        |                 |                                                                                           |
|        | equal-int       |                                                                                           |
|        |                 |                                                                                           |
|        | unequal-int     |                                                                                           |
|        |                 |                                                                                           |
|        | greater-int     |                                                                                           |
|        |                 |                                                                                           |
|        | smaller-int     |                                                                                           |
|        |                 |                                                                                           |
|        | and-int         |                                                                                           |
|        |                 |                                                                                           |
|        | or-int          |                                                                                           |
|        |                 |                                                                                           |


