package service

import(
	"context"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"github.com/google/uuid"
	"log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"errors"
	//"time"

)

// LaptopServer is a struct that provides laptop services
type LaptopServer struct{
	laptopStore LaptopStore //interface object
	imageStore ImageStore
}

// Returns a new LaptopServer struct instance
func NewLaptopServer(laptop LaptopStore, image ImageStore) *LaptopServer{
	return &LaptopServer{
		laptopStore : laptop,
		imageStore : image,
	}
}

//CreateLaptop is a unary rpc to create a new laptop
func (laptopserver *LaptopServer) CreateLaptop(
	ctx context.Context, 
	req *pb.CreateLaptopRequest,
	)(*pb.CreateLaptopResponse, error){ //missing function body, non declaration statement outside function body error when this statement was on next line

		laptop := req.GetLaptop()

		log.Printf("received a create laptop request with id: %s",laptop.Id)


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


		// heavy code
		//time.Sleep(6 * time.Second)
		err := laptopserver.laptopStore.Save(laptop)
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

func (laptopServer *LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest, 
	stream pb.LaptopService_SearchLaptopServer,
) error{
	filter := req.GetFilter()
	log.Printf("received a search laptop request with filter: %v",filter)

	err := laptopServer.laptopStore.Search(
		stream.Context(),
		filter,
		func (laptop *pb.Laptop) error{
			res := &pb.SearchLaptopResponse{Laptop : laptop}

			err := stream.Send(res)

			if err != nil{
				return err
			}

			log.Printf("sent laptop with id: %s",laptop.GetId())
			return nil
		},
	)

	if err != nil{
		return status.Errorf(codes.Internal, "Unexpected error: %v",err)
	}

	return nil
}

func (laptopServer *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error{

	return nil
}
