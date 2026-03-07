package main

import "fmt"

func main() {
	var int64 int64 = 32768
	var int32 int32 = int32(int64)
	var int16 int16 = int16(int64)

	fmt.Println(int64)
	fmt.Println(int32)
	fmt.Println(int16)

	var name string = "Muhammad Tio Laksono"
	var m uint8 = name[0]
	var mString string = string(m)

	fmt.Println(name)
	fmt.Println(m)
	fmt.Println(mString)

}
