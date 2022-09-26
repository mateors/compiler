package main

import (
	"encoding/binary"
	"fmt"

	"github.com/mateors/compiler/code"
)

func operandOrderTest() {
	// In case an operand needs multiple bytes to be accurately represented, the order in which it’s encoded
	// plays a big role. There are two possible orders, called little endian and big endian. Little endian
	// means that the least significant byte of the original data comes first and is stored in the lowest
	// memory address. Big endian is the opposite: the most significant byte comes first
	operand := 256
	offset := 1
	instruction := make([]byte, 3)
	binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
	//binary.LittleEndian.PutUint16(instruction[offset:], uint16(o))
	fmt.Println(instruction) //LittleEndian=[0 254 255], BigEndian= [0 255 254]
}

func main() {

	// fmt.Println([]int{65534})
	// fmt.Println([]byte{byte(0), 255, 254})
	// fmt.Println(math.Pow(2, 16)) //2--16

	// instruction := code.Make(0, 65534)
	// fmt.Println(instruction) //[0 255 254]

	//fmt.Println(code.Make(0, 0)) //
	//fmt.Println(code.Make(0, 1)) //

	//bs := code.Make(code.OpConstant, 65534)
	//fmt.Println(bs)

	// input := `1 + 2 + 3 + 4`
	// program := compiler.Parser(input)
	// compiler := compiler.New()
	// err := compiler.Compile(program)
	// if err != nil {
	// 	//t.Fatalf("compiler error: %s", err)
	// }

	// bytecode := compiler.Bytecode()
	// fmt.Println(bytecode.Instructions)
	// fmt.Println(bytecode.Constants)

	multipleInstructions := code.Instructions{}
	fmt.Println(multipleInstructions)
	multipleInstructions = append(multipleInstructions, code.Make(code.OpConstant, 10)...)
	multipleInstructions = append(multipleInstructions, code.Make(code.OpConstant, 20)...)
	fmt.Println(multipleInstructions)

	//singleInstruction := code.Instructions(code.Make(code.OpConstant, 3))
	
	//fmt.Println(singleInstruction.String())

	//_, err := code.Lookup(2)
	//fmt.Println(err) //opcode 2 undefined

}
