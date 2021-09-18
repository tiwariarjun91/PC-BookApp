package service

import(
	"sync"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"errors"
	"github.com/jinzhu/copier"
	"fmt"
	"context"
	"log"
	"time"
)


var	ErrAlreadyExists = errors.New("Record already exists") // this type cant be in var ( var name var type) format


// LaptopStore is an interface to store laptop
type LaptopStore interface{
	// save saves the laptop to the store
	Save(in *pb.Laptop) error
	Find(Id string) (*pb.Laptop,error)
	Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
}

// InMemoryLaptopStorage stores the laptop in memory
type InMemoryLaptopStorage struct{
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// NewInMemoryLaptopStorage() returns a new NewInMemoryLaptopStorage
func NewInMemoryLaptopStorage() *InMemoryLaptopStorage{
	return &InMemoryLaptopStorage{
		data : make(map[string]*pb.Laptop),
	}
}

//
func (store *InMemoryLaptopStorage) Save(laptop *pb.Laptop) error{
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil{ // to check if "laptop" exists in the memory
		return ErrAlreadyExists
	}

	//Deep copy of laptop
	other, err := deepCopy(laptop)
	if err!= nil{
		return fmt.Errorf("Cannnot copy laptop data: %w",err)
	}

	store.data[other.Id] = other
	return nil
}


func (store *InMemoryLaptopStorage) Find(Id string) (*pb.Laptop,error){
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[Id]

	if laptop == nil{
		return nil,nil
	}

	return deepCopy(laptop)

}


func (store *InMemoryLaptopStorage) Search(
	ctx context.Context,
	filter *pb.Filter,
	found func(laptop *pb.Laptop) error,
) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	for _, laptop := range store.data{

		time.Sleep(time.Second)
		log.Print("checking laptop id: ", laptop.GetId())
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is cancelled")
			return nil
		}
		if isQualified(filter, laptop){
			other, err := deepCopy(laptop)
			if err != nil{
				return err
			}
			err = found(other)
			if err != nil{
				return err
			}
		}
	}
	return nil
}



func isQualified(filter *pb.Filter, laptop *pb.Laptop) bool{
	if laptop.GetPriceUsd() > filter.GetMaxPriceUsd(){
		return false
	}
	if laptop.GetCpu().GetNumberCores() < filter.GetMinCpuCores(){
		return false
	}
	if laptop.GetCpu().GetMinGhz() < filter.GetMinCpuGhz(){
		return false
	}
	if toBit(laptop.GetRam()) < toBit(filter.GetMinRam()){ // as ram is of different units this function will convert it to same minimum unit
		return false
	}

	return true

}

func toBit(memory *pb.Memory) uint64{
	value := memory.GetValue()

	switch memory.GetUnit(){
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3 // 2^3 = 8
	case pb.Memory_KILOBYTE:
		return value << 13 // 2^10 = 1024 & 2^3 = 8 therefore 2^13 = 1024*8
	case pb.Memory_MEGABYTE:
		return value << 23
	case pb.Memory_GIGABYTE:
		return value << 33
	case pb.Memory_TERABYTE:
		return value << 43
	default:
		return 0
	
	}
}


func deepCopy(laptop *pb.Laptop) (*pb.Laptop, error){
	other := &pb.Laptop{}

	err := copier.Copy(other,laptop)
	if err != nil{
		return nil,fmt.Errorf("Cannot copy data %w",err)
	}

	return other,nil
}

// DBLaptopStorage stores the laptop to database
/*type DBLaptopStorage struct{

}*/