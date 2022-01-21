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
	"time"
	"io"
	"os"
	"path/filepath"


)


func createLaptop(laptopClient pb.LaptopServiceClient){
	laptop := sample.NewLaptop()
	laptop.Id = ""

	req := &pb.CreateLaptopRequest{
		Laptop : laptop,
	}

	ctx,cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res,err := laptopClient.CreateLaptop(ctx,req)
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

func searchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter){
	log.Print("Seach filter %v", filter)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	req := &pb.SearchLaptopRequest{Filter : filter}
	stream,err := laptopClient.SearchLaptop(ctx, req)
	if err != nil{
		log.Fatal("Cannot search laptop: ",err)
	}

	for{
		res,err := stream.Recv()
		if err == io.EOF{
			return
		}
		if err != nil{
			log.Fatal("Cannot receice response : ",err)
		}

		laptop := res.GetLaptop()
		log.Print("  - found: ",laptop.GetId())
		log.Print("  + brand: ",laptop.GetBrand())
		log.Print("  + name: ",laptop.GetName())
		log.Print("  + cpu cores: ", laptop.GetCpu().GetNumberCores())
		log.Print("  + cpu min ghz: ", laptop.GetCpu().GetMinGhz())
		log.Print("  + ram: ", laptop.GetRam())
		log.Print("  + price: ", laptop.GetPriceUsd())
	}
}

func uploadImage(laptopClient pb.LaptopServiceClient )  {
	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop : laptop,
	}
	_,err := laptopClient.CreateLaptop(ctx,req)
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

	file, err := os.Open("tmp/download.png")
	if err != nil{
		log.Printf("Cannot open the file, %s", err)
		return 
	}

	defer file.Close()

	ctx, cancel := Context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := laptopClient.UploadImage(ctx)
	if err != nil(
		log.Printf("Cannot create the stream %s", err)
		return 
	)

	req := &pb.UploadImageRequest{
		Data : pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId : laptop.Id,
				ImageType : filepath.Ext("tmp/download.png"),
			},
		},

	}

	err = stream.Send(req)
	if err != nil{
		log.Printf("Error sending image info with the stream %s", err)
		
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for{
		n, err := reader.read(buffer)
		if err == io.EOF{
			break
		}

		
	}

	
}

func main(){
	serverAddress := flag.String("address","","The server address") // go run main.go --address "actual address"
	flag.Parse()
	log.Printf("dial server %s",*serverAddress)

	conn,err := grpc.Dial(*serverAddress,grpc.WithInsecure())
	defer conn.Close()
	if err != nil{
		log.Fatalf("Cannot dial server : ",err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	for i:= 0 ; i < 10; i++{
		createLaptop(laptopClient)
	}

	filter := &pb.Filter{
		MaxPriceUsd : 3000,
		MinCpuCores : 4,
		MinCpuGhz : 2.5,
		MinRam : &pb.Memory{Value : 8, Unit : pb.Memory_GIGABYTE},
	}

	searchLaptop(laptopClient, filter)
	}
