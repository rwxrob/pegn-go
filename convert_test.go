package pegn

import (
	"fmt"
)

func ExampleConvertToHex() {
	fmt.Println(convertToHex("001010", 2))
	fmt.Println(convertToHex("000012", 8))
	fmt.Println(convertToHex("000010", 10))

	// Output:
	// 0A <nil>
	// 0A <nil>
	// 0A <nil>
}

func ExampleConvert_binary() {
	fmt.Println(convert("1010", 2, 16))
	fmt.Println(convert("0001", 2, 10))

	// Output:
	// A <nil>
	// 1 <nil>
}

func ExampleConvert_hex() {
	fmt.Println(convert("A", 16, 10))
	fmt.Println(convert("FF", 16, 10))

	// Output:
	// 10 <nil>
	// 255 <nil>
}

func ExampleConvert_oct() {
	fmt.Println(convert("12", 8, 10))
	fmt.Println(convert("12", 8, 16))

	// Output:
	// 10 <nil>
	// A <nil>
}
