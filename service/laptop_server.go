package service

import(
	"context"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"github.com/google/uuid"
	"log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"errors"

)

// LaptopServer is a struct that provides laptop services
type LaptopServer struct{
	Store LaptopStore //interface object
}

// Returns a new LaptopServer struct instance
func NewLaptopServer(store LaptopStore) *LaptopServer{
	return &LaptopServer{store}
}

//CreateLaptop is a unary rpc to create a new laptop
func (laptopserver *LaptopServer) CreateLaptop(
	ctx context.Context, 
	req *pb.CreateLaptopRequest,
	)(*pb.CreateLaptopResponse, error){ //missing function body, non declaration statement outside function body error when this statement was on next line

		laptop := req.GetLaptop()

		if len(laptop.Id) > 0{ // if the user has sent an Id
			_,err := uuid.Parse(laptop.Id)
			if err != nil{
				return nil, status.Errorf(codes.InvalidArgument, "Laptop Id is not a valid uuid: %v",err)
			}
		}else {
			id,err := uuid.NewRandom()
			if err != nil{
				return nil,status.Errorf(codes.Internal, "Cannot generate a new laptop id: %v",err)
			}
			laptop.Id = id.String()
		}

		err := laptopserver.Store.Save(laptop)
		code := codes.Internal
		if errors.Is(err,ErrAlreadyExists){
			code = codes.AlreadyExists // codes. not code.
		}
		if err!= nil{
			return nil, status.Errorf(code, "Cannot store laptop to the store: %v",err) // unexpected new line, expecting =,:= error when missed return statement
		}

		log.Printf("saved laptop with id: %s",laptop.Id)

		res := &pb.CreateLaptopResponse{
			Id : laptop.Id,
		}

		return res,nil

}