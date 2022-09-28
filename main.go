package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"reflect"

	"github.com/mateors/compiler/code"
	"github.com/mateors/compiler/compiler"
	"github.com/mateors/lexer/ast"
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

func TestC(node ast.Node) {

	//fmt.Println("TestC:", reflect.TypeOf(node).String())
	inter := reflect.TypeOf((*ast.Node)(nil)).Elem()
	isNode := reflect.TypeOf(node).Implements(inter)
	fmt.Println(reflect.TypeOf(node), "isNode:", isNode)

	switch nd := node.(type) {

	case *ast.Program:

		for _, stm := range nd.Statements {

			//fmt.Println("stm>", stm, reflect.TypeOf(stm)) //*ast.ExpressionStatement
			TestC(stm)
		}

	case *ast.ExpressionStatement:
		//fmt.Println("ExpressionStatement>", nd)
		TestC(nd.Expression)

	case *ast.InfixExpression:
		//fmt.Println("InfixExpression>", nd.Operator, nd.Token)
		TestC(nd.Left)
		TestC(nd.Right)

	case *ast.IntegerLiteral:
		fmt.Println("IntegerLiteral>", nd.Token, nd.Value)

	default:
		fmt.Println("default", reflect.TypeOf(nd))

	}
}

func main() {

	// fmt.Println([]int{65534})
	// fmt.Println([]byte{byte(0), 255, 254})
	// fmt.Println(math.Pow(2, 16)) //2--16

	//fmt.Println(code.Make(0, 0)) //
	//fmt.Println(code.Make(0, 1)) //

	//bs := code.Make(code.OpConstant, 65534)
	//fmt.Println(bs)

	input := `1 + 2`
	program := compiler.Parser(input)
	//fmt.Println(program.Statements)
	//fmt.Println(program.String())
	//fmt.Println("TokenLiteral", program.TokenLiteral())

	//check if program underlying type is ast.Node or not
	//fmt.Println(reflect.TypeOf(program))
	v, ok := interface{}(program).(ast.Node)
	//fmt.Println(v, ok)
	if ok {
		TestC(v)

	}
	os.Exit(1)

	compiler := compiler.New()

	err := compiler.Compile(program)
	if err != nil {
		fmt.Println(err)
	}

	bytecode := compiler.Bytecode()
	fmt.Println("Instructions:", bytecode.Instructions)
	fmt.Println("Constants:", bytecode.Constants)
	os.Exit(1)

	multipleInstructions := code.Instructions{}
	multipleInstructions = append(multipleInstructions, code.Make(code.OpConstant, 10)...)
	multipleInstructions = append(multipleInstructions, code.Make(code.OpConstant, 20)...)
	multipleInstructions = append(multipleInstructions, code.Make(code.OpConstant, 10)...)
	fmt.Println("multipleInstructions:", multipleInstructions.String())

	//def, err := code.Lookup(0) //code.OpConstant = 0
	def, err := code.Lookup(byte(code.Opcode(code.OpConstant)))
	fmt.Println(err, def.Name, def.OperandWidths) // OpConstant [2]

	//single instruction
	instruction := code.Make(0, 266)
	fmt.Println(instruction) //bytecode = [opcode operand] = [opcode, argbyte1 argbyt2] = [0 1 10]

	fmt.Println("instruction[1:]", instruction[1:]) //operand = [argbyte1 argbyt2] = [1 10]
	operandsRead, offset := code.ReadOperands(def, instruction[1:])
	fmt.Println(operandsRead, offset) //[266] 2

	fmt.Println("binary.BigEndian.Uint16:", binary.BigEndian.Uint16([]byte{1, 10})) //255+1+10 = 266

}
