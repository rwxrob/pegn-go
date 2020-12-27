package pegen

import (
	"fmt"
)

func ExampleConvertToHex() {
	fmt.Println(ConvertToHex("001010", 2))
	fmt.Println(ConvertToHex("000012", 8))
	fmt.Println(ConvertToHex("000010", 10))

	// Output:
	// 0A <nil>
	// 0A <nil>
	// 0A <nil>
}

func ExampleConvert_binary() {
	fmt.Println(Convert("1010", 2, 16))
	fmt.Println(Convert("0001", 2, 10))

	// Output:
	// A <nil>
	// 1 <nil>
}

func ExampleConvert_hex() {
	fmt.Println(Convert("A", 16, 10))
	fmt.Println(Convert("FF", 16, 10))

	// Output:
	// 10 <nil>
	// 255 <nil>
}

func ExampleConvert_oct() {
	fmt.Println(Convert("12", 8, 10))
	fmt.Println(Convert("12", 8, 16))

	// Output:
	// 10 <nil>
	// A <nil>
}
