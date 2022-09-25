# Compiler from scratch

### Vocabulary
* immerse yourself = নিজেকে নিবিষ্ট
* intimidating
* speculations = অনুমান
* insurmountable = অনতিক্রম্য
* vague = অস্পষ্ট
* riddle =  ধাঁধা

## What we will learn?
We are going to turn our tree-walking and on-the-fly-evaluating interpreter into a bytecode compiler and a virtual machine that executes the bytecode.


#### Bytecode compilers and virtual machines are everywhere – and for good reason.

* bytecode that's passed from the compiler to the virtual machine – 
* that makes the system more modular, 
* the main appeal of this architecture lies in its performance. 
* Bytecode interpreters are fast.

> At their core, compilers and virtual machines are ideas – patterns – just like "interpreter" or "web server" are ideas that can have multiple implementations, ranging from tiny to massive. 


### special type of compiler
> a compiler that translates high-level languages to produce executables is just one special type of compiler


### producing executables is just one variation of "translating source code".

You see, compilers are fundamentally about translation, because translation is how they implement a programming language

Programming means giving instructions to a computer. We, the programmers, write these instructions in programming languages the computer can understand. There is no point in using any other language. Now, implementing a programming language means making the computer understand it. There are two ways to do that: either interpret the language for the computer on-the-fly or translate it into another language, one the computer already understands.

They both have a frontend that reads in source code written in the source language and turns it into a data structure. In both, compiler and interpreter, this frontend is usually made up of a lexer and a parser that together generate a syntax tree. So here, in the front part, they have a lot of similarities. After that, when they both traverse the AST, that’s when their paths diverge.

First, the source code is tokenized and parsed by the lexer and the parser. We’re familiar with this part from our interpreter. It’s called the frontend. The source code is turned from text into an AST.

## Von Neumann architecture

In Von Neumann's model a computer has two central parts: a processing unit, which contains an arithmetic logic unit (ALU) and multiple processor registers, and a control unit with an instruction register and a program counter. Together they’re called the central processing unit,often shortened to CPU. Besides that, the computer also contains memory (RAM), mass storage (think: hard drive) and input/output devices (keyboard and display). CPU, memory, mass storage and I/O. Here’s a rough sketch of such a computer:

Instructions only become instructions once the CPU fetches them from memory, decodes them, and finds out that they are, in fact, proper instructions and then executes them. 


We have a memory region where the CPU accesses and stores data in a LIFO manner. It does this in order to implement a specialized version of a stack, called a call stack.

Why does it need a call stack? Because the CPU (or maybe: the programmer that wants the CPU to work as intended) needs to keep track of certain information in order to execute a program.

### Return address
which instruction to fetch after the current function, is called the return address. It's where the CPU returns to after executing the current
function.

### Registers
Registers are part of the CPU and accessing them is much faster than accessing the main memory. A CPU of the x86-64 architecture, for example, has 16 general purpose registers, each holding 64 bits of data

### call stack.
It's the memory region that holds the stack.

It's a region in memory where data is managed in a last-in-first-out (LIFO) manner. The data in it grows and shrinks, you push elements on to the stack and later pop them off. Just like the stack data structure. But unlike this generic data structure, the stack is focused on one purpose: it's used to implement the call stack.

> We have a memory region where the CPU accesses and stores data in a LIFO manner. It does this in order to implement a specialized version of a stack, called a call stack.

### Why does it need a call stack? 
Because the CPU (or maybe: the programmer that wants the CPU to work as intended) needs to keep track of certain information in order to execute a program. The call stack helps with that. What information? First and foremost: which function is currently being executed and which instruction to execute next, once the current function is fully executed. This piece of information, which instruction to fetch after the current function, is called the return address. It’s where the CPU returns to after executing the current function.

Registers are used to store data that's small but frequently accessed. 

### Stack pointer
A stack pointer is a small register that stores the memory address of the last data element added to the stack or, in some cases, the first available address in the stack.

A stack register is a CPU's register whose purpose is to keep track of a call stack. 

## Virtual Machine
> a virtual machine is a computer built with software. It's a software entity that mimics how a computer works.

A virtual machine has a run loop that goes through the fetch-decode-execute cycle, just like a computer. It has a program counter; it fetches instructions; it decodes and executes them. It also has a stack, just like a real computer. Sometimes it has a call stack and sometimes even registers. All built in software.


## Virtual Machine explanation using javascript code:

```js
let virtualMachine = function(program) {

  let programCounter = 0;
  let stack = [];
  let stackPointer = 0;

  while (programCounter < program.length) {

    let currentInstruction = program[programCounter];

     switch (currentInstruction) {

        case PUSH:
            stack[stackPointer] = program[programCounter+1];
            stackPointer++;
            programCounter++;
            break;

        case ADD:
            right = stack[stackPointer-1]
            stackPointer--;
            left = stack[stackPointer-1]
            stackPointer--;
            stack[stackPointer] = left + right;
            stackPointer++;
            break;

        case MINUS:
            right = stack[stackPointer-1]
            stackPointer--;
            left = stack[stackPointer-1]
            stackPointer--;
            stack[stackPointer] = left - right;
            stackPointer++;
            break;

     }

     programCounter++;

  }

  console.log("stacktop: ", stack[stackPointer-1]);

}
```
