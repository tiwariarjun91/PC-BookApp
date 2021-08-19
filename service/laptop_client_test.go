package service_test

import(
	"testing"
	"github.com/tiwariarjun91/PC-BookApp/service"
	"google.golang.org/grpc"
	"net"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"github.com/tiwariarjun91/PC-BookApp/sample"
	"context"


)

func TestClientCreateLaptop(t *testing.T){
	laptopServer, serverAddress := startTestingLaptopServer(t)
	laptopclient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()

	expectedId := laptop.Id

	req := &pb.CreateLaptopRequest{
		Laptop : laptop,
	}

	res, err := laptopclient.CreateLaptop(context.Background(),req)

	if err!= nil || res.Id != expectedId{
		t.Fatalf("Test case failed")
	}

	// to check if the laptop got stored in the memory

	other,err := laptopServer.Store.Find(res.Id)
	if err!=nil || other == nil{
		t.Fatalf("Test case failed")
	}
}

func startTestingLaptopServer(t *testing.T) (*service.LaptopServer, string){
	
	laptopserver := service.NewLaptopServer(service.NewInMemoryLaptopStorage())

	grpcServer :=grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer,laptopserver)

	listener,err := net.Listen("tcp",":0")
	if err != nil{
		t.Fatalf("Test case failed")
	}

	go grpcServer.Serve(listener) // non block code

	return laptopserver, listener.Addr().String()
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil{
		t.Fatalf("Test case failed")
	}
	return pb.NewLaptopServiceClient(conn)
}
