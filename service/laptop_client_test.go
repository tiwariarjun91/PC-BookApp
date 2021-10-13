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

func TestClientSearchLaptop(t *testing.T) {
	t.Parallel()

	filter := &pb.Filter{
		MaxPriceUsd: 2000,
		MinCpuCores: 4,
		MinCpuGhz:   2.2,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}

	laptopStore := service.NewInMemoryLaptopStore()
	expectedIDs := make(map[string]bool)

	for i := 0; i < 6; i++ {
		laptop := sample.NewLaptop()

		switch i {
		case 0:
			laptop.PriceUsd = 2500
		case 1:
			laptop.Cpu.NumberCores = 2
		case 2:
			laptop.Cpu.MinGhz = 2.0
		case 3:
			laptop.Ram = &pb.Memory{Value: 4096, Unit: pb.Memory_MEGABYTE}
		case 4:
			laptop.PriceUsd = 1999
			laptop.Cpu.NumberCores = 4
			laptop.Cpu.MinGhz = 2.5
			laptop.Cpu.MaxGhz = laptop.Cpu.MinGhz + 2.0
			laptop.Ram = &pb.Memory{Value: 16, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		case 5:
			laptop.PriceUsd = 2000
			laptop.Cpu.NumberCores = 6
			laptop.Cpu.MinGhz = 2.8
			laptop.Cpu.MaxGhz = laptop.Cpu.MinGhz + 2.0
			laptop.Ram = &pb.Memory{Value: 64, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		}

		err := laptopStore.Save(laptop)
		if err != nil{
			return err
		}
	}

	serverAddress := startTestLaptopServer(t, laptopStore, nil, nil)
	laptopClient := newTestLaptopClient(t, serverAddress)

	req := &pb.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(context.Background(), req)
	if err != nl{
		return err
	}
	found := 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil{
			return err
		}
		found += 1
	}

	if err != nil{
		return err
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
