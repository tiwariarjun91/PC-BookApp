package serializer

import(
	"testing"
	"github.com/tiwariarjun91/PC-BookApp/sample"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"fmt"

)

func TestWriteProtobufToBinaryFile(t *testing.T){
	t.Parallel()
	binaryFile := "../tmp/laptop.bin" // test case failed first as there was no tmp folder in PC-Book
	laptop := sample.NewLaptop()

	err := WriteProtobufToBinaryFile(laptop, binaryFile)
	if err != nil{
		fmt.Println(err)
		t.Fatalf("test case failed")
	}
}

func TestReadProtobufFromBinaryFile(t *testing.T){
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"
	laptop := &pb.Laptop{}

	err := ReadProtobufFromBinaryFile(binaryFile, laptop)
	//fmt.Println(string(binaryFile))
	//fmt.Println(laptop)
	if err != nil{
		fmt.Println(err)
		t.Fatalf("test case failed")
	}
}