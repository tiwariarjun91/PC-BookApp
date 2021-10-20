package main //go run  main.go -port 8080

import (
	"fmt"
	"net"
	"log"
	"flag"
	"github.com/tiwariarjun91/PC-BookApp/service"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"google.golang.org/grpc"

)

func main(){
	port := flag.Int("port",0,"The Server Port")
	flag.Parse()
	log.Printf("Start server on port %d",*port)

	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStorage())

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer,laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d",*port)

	listener, err := net.Listen("tcp",address)
	if err != nil{
		log.Fatal("Cannot Start Server ",err)
	}

	err = grpcServer.Serve(listener) //go run main.go -port 8080 //sample port
	if err != nil{
		log.Fatal("Cannot Start Server ",err)
	}
	// http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){ }) // returns a handler for this wrapped function
}