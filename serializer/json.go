package serializer

import(
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)

// Converts protocol buffer message to json string
func ProtobufToJSON(message proto.Message) (string, error){
	marshaler := jsonpb.Marshaler{ // Marshaleris a struct 
		EnumsAsInts : false,
		EmitDefaults : true,
		Indent : " ",
		OrigName : true, // use the original field name defined in message
	}

	return marshaler.MarshalToString(message) // this method converts protobuf message to string
	// returns string and error
}