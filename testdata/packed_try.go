package main

/*
#include <string.h>
#include <stdint.h>

typedef struct Student {
    int a;
	char  b[5];
	int  c;
}Student;

Student N = {12, "TEST", 12};
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"unsafe"
)

type A struct {
	A int32
	B [5]byte
	C int32
}

func main() {
	sc := &C.N

	var st A

	cdata := C.GoBytes(unsafe.Pointer(sc), C.sizeof_Student)

	binary.Read(bytes.NewBuffer(cdata[0:4]), binary.LittleEndian, &st.A)
	binary.Read(bytes.NewBuffer(cdata[4:9]), binary.LittleEndian, &st.B)
	binary.Read(bytes.NewBuffer(cdata[12:16]), binary.LittleEndian, &st.C)

	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, st)
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Sizeof myStruct: %d, Sizeof Binary: %d, Sizeof buf: %d, Len of buf: %d\n",
		unsafe.Sizeof(st),
		binary.Size(buf.Bytes()),
		unsafe.Sizeof(buf),
		buf.Len())

	by := buf.Bytes()

	fmt.Println(by)
	// Output
	//Error: <nil>
	//Sizeof myStruct: 16, Sizeof Binary: 13, Sizeof buf: 8, Len of buf: 13

	fmt.Println("---------------------------")
	gobByte(cdata)
}

func gobByte(cdata []byte) {
	var st A

	buf := bytes.NewBuffer(cdata)
	err := gob.NewDecoder(buf).Decode(&st)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(st)
}
