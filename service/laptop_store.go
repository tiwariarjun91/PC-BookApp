package service

import(
	"sync"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"errors"
	"github.com/jinzhu/copier"
	"fmt"
)


var	ErrAlreadyExists = errors.New("Record already exists") // this type cant be in var ( var name var type) format


// LaptopStore is an interface to store laptop
type LaptopStore interface{
	// save saves the laptop to the store
	Save(in *pb.Laptop) error
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
	other := &pb.Laptop{}
	err := copier.Copy(other,laptop)
	if err!= nil{
		return fmt.Errorf("Cannnot copy laptop data: %w",err)
	}

	store.data[other.Id] = other
	return nil
}

// DBLaptopStorage stores the laptop to database
/*type DBLaptopStorage struct{

}*/