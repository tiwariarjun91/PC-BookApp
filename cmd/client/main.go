package main // go run main.go -address 0.0.0.0:8080

import (
	"flag"
	"log"
	"context"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"github.com/tiwariarjun91/PC-BookApp/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"


)

func main(){
	serverAddress := flag.String("address","","The server address")
	flag.Parse()
	log.Printf("dial server %s",*serverAddress)

	conn,err := grpc.Dial(*serverAddress,grpc.WithInsecure())
	if err != nil{
		log.Fatalf("Cannot dial server : ",err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := sample.NewLaptop()
	laptop.Id = ""

	req := &pb.CreateLaptopRequest{
		Laptop : laptop,
	}

	res,err := laptopClient.CreateLaptop(context.Background(),req)
	if err != nil{
		st,ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("Laptop already exists")
		} else{
			log.Fatalf("Cannot create laptop: ",err)
		}
		return
	}

	log.Printf("Created Laptop with Id : %s",res.Id)
}
