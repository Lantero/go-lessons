// Base types and their zero values
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Bool            bool = false
	BoolZero        bool
	String          string = "Some text"
	StringZero      string
	IntNormal       int = -27 // int  int8  int16  int32  int64
	IntNormalZero   int
	IntUnsigned     uint = 27 // uint uint8 uint16 uint32 uint64
	IntUnsignedZero uint
	Float           float64 = 3.64 // float32 float64
	FloatZero       float64
	Complex         complex128 = cmplx.Sqrt(-5 + 12i) // complex64 complex128
	ComplexZero     complex128
	Rune            rune = 'ñ' // Syntactic sugar for int32, used for semantic purposes
	RuneZero        rune
	Byte            byte = 12 // Syntactic sugar for uint8, used for semantic purposes
	ByteZero        byte
)

func main() {
	fmt.Printf("Type: %T Value: %v Zero Value: %v\n", Bool, Bool, BoolZero)
	fmt.Printf("Type: %T Value: %v Zero Value: %v\n", String, String, StringZero)
	fmt.Printf("Type: %T Value: %v Zero Value: %v\n", IntNormal, IntNormal, IntNormalZero)
	fmt.Printf("Type: %T Value: %v Zero Value: %v\n", IntUnsigned, IntUnsigned, IntUnsignedZero)
	fmt.Printf("Type: %T Value: %v Zero Value: %v\n", Float, Float, FloatZero)
	fmt.Printf("Type: %T Value: %v Zero Value: %v\n", Complex, Complex, ComplexZero)
	fmt.Printf("Type: %T Value: %v Zero Value: %v\n", Rune, Rune, RuneZero)
	fmt.Printf("Type: %T Value: %v Zero Value: %v\n", Byte, Byte, ByteZero)
}
