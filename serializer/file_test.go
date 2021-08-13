package serializer

import(
	"testing"
	"github.com/tiwariarjun91/PC-BookApp/sample"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"fmt"
	"github.com/golang/protobuf/proto"

)

func TestSerializer(t *testing.T){
	t.Parallel()
	binaryFile := "../tmp/laptop.bin" // test case failed first as there was no tmp folder in PC-Book
	laptop1 := sample.NewLaptop()

	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	if err != nil{
		fmt.Println(err)
		t.Fatalf("test case failed")
	}

	laptop2 := &pb.Laptop{}

	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	if err != nil{
		fmt.Println(err)
		t.Fatalf("test case failed")
	}
	status := proto.Equal(laptop1, laptop2)

	if status != true{
		t.Fatalf("test case failed")
	}

	jsonFile := "../tmp/laptop.json"

	err = WriteProtobufToJSONFile(laptop1, jsonFile)
	if err != nil{
		fmt.Println(err)
		t.Fatalf("test case failed")
	}
}

