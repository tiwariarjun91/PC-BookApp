package serializer

import(
	"github.com/golang/protobuf/proto"
	"fmt"
	"io/ioutil"
)

// Writes protocol buffer message to JSON file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil{
		return fmt.Errorf("Cannot convert proto message to JSON: %w",err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0064) // data is a string //we also use []byte(string) while unmarshaling as Unmarshal requires byte arry as Marshal returns a byte array
	if err != nil{ // the second argument in WriteFile needs to be an array of bytes
		return fmt.Errorf("Cannot write to JSON data to file: %w",err)
	}

	return nil
}
// Writes protocol buffer message to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error{
	data, err := proto.Marshal(message) // converts proto message to binary
	if err != nil{
		return fmt.Errorf("Cannot marshal proto message to binary: %w",err)
	}

	err = ioutil.WriteFile(filename, data, 0644) // creates a new file with binary data
	if err != nil{
		return fmt.Errorf("Cannot write binary data to file: %w", err)
	}

	return nil
}

// Reads protocol buffer message from binary file
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error{ // m in message should be uppercase 
	data, err := ioutil.ReadFile(filename)
	//fmt.Println(string(data))
	if err != nil{
		return fmt.Errorf("Cannot read binary data from file: %w", err)
	}

	err = proto.Unmarshal(data, message)//did not pass reference to message unlike xml/json . Unmarshal // usually unmarshal is used to convert binary data to suitable types like structs, protobuf messages
	if err != nil{
		return fmt.Errorf("Cannot unmarshal binary to proto message: %w",err)
	}

	return nil
}