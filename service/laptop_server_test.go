package service

import(
	"testing"
	"github.com/tiwariarjun91/PC-BookApp/sample"
	"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"context"
)

func TestServerCreateLaptop(t *testing.T) {


	withoutId := sample.NewLaptop()
	withoutId.Id = ""
	laptopduplicate := sample.NewLaptop()
	storeduplicate := NewInMemoryLaptopStorage()
	err := storeduplicate.Save(laptopduplicate)
	if err != nil{
		t.Fatalf("Test case failed")
	}
	invalidId := sample.NewLaptop()
	invalidId.Id = "invalidId654"
	testcases := []struct {
		name string
		laptop *pb.Laptop
		store  LaptopStore // need to read about this concept once again
		code codes.Code}{

		{
			name : "Success with ID",
			laptop : sample.NewLaptop(),
			store : NewInMemoryLaptopStorage(), // need to read about this concept once again
			code : codes.OK,

		},{
			name : "Success Without ID",
			laptop : withoutId,
			store : NewInMemoryLaptopStorage(),
			code : codes.OK,
		},{
			name : "Fail with duplicate ID",
			laptop : laptopduplicate,
			store : storeduplicate,
			code: codes.AlreadyExists,
		},{
			name : "Fail with Invalid ID",
			laptop : invalidId,
			store : NewInMemoryLaptopStorage(),
			code : codes.InvalidArgument,
		},



	}

	for i := range testcases{
		tc := testcases[i]

		req := &pb.CreateLaptopRequest{
			Laptop : tc.laptop,
		}
		server := NewLaptopServer(tc.store)

		res,err := server.CreateLaptop(context.Background(), req)
		
		if tc.code == codes.OK{
			if err != nil || res.Id != req.Laptop.Id{
				t.Fatalf("Testcase failed")
			}
		} else{
			if err == nil{
				t.Fatalf("Testcase failed")
			}
		}
	}
}