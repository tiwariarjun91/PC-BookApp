package sample
import(
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"github.com/golang/protobuf/ptypes"
)

//NewKeyboard() returns a new keyboard sample
func NewKeyboard() *pb.Keyboard{
	keyboard := &pb.Keyboard{
		Layout: randomKeyBoardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

//NewCpu() returns a new CPU sample
func NewCPU() *pb.CPU{
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2,8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat64(2.0,3.5)
	maxGhz := randomFloat64(minGhz, 5.0)
	cpu := &pb.CPU{		
		Brand: brand,
		Name: name,
		NumberCores: uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz : minGhz,
		MaxGhz : maxGhz,
	}
	return cpu
}

//NewGPU() will return a sample GPU

func NewGPU() *pb.GPU{
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0,1.5)
	maxGhz := randomFloat64(minGhz,2.0)
	memory := &pb.Memory{
		Value : uint64(randomInt(2,6)),
		Unit : pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand : brand,
		Name : name,
		MinGhz : minGhz,
		MaxGhz : maxGhz, 
		Memory : memory,
	}
	
	return gpu
}

//NewRAM() returns a new sample RAM
func NewRAM() *pb.Memory{
	ram := &pb.Memory{
		Value : uint64(randomInt(4,64)),
		Unit : pb.Memory_GIGABYTE,
	}

	return ram
}

//NewSSD() returns a new sample SSD storage
func NewSSD() *pb.Storage{
	ssd := &pb.Storage{
		Driver : pb.Storage_SSD,
		Memory : &pb.Memory{
		Value : uint64(randomInt(128,1024)),
		Unit : pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

//NewHDD() returns a new sample HDD storage
func NewHDD() *pb.Storage{
	hdd := &pb.Storage{
		Driver : pb.Storage_HDD,
		Memory : &pb.Memory{
		Value : uint64(randomInt(1,6)),
		Unit : pb.Memory_TERABYTE,
		},
	}

	return hdd
}

//NewScreen() returns a new sample screen

func NewScreen() *pb.Screen{
	screen := &pb.Screen{
		SizeInch : randomFloat32(13,17),
		Resolution : randomScreenResolution(),
		Panel : randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

//NewLaptop() returns a new sample laptop

func NewLaptop() *pb.Laptop {

	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	newLaptop := &pb.Laptop{
		Id : randomId(),
		Brand : brand,
		Name : name,
		Cpu : NewCPU(),
		Ram : NewRAM(),
		Gpus : []*pb.GPU{NewGPU()},
		Storages :[]*pb.Storage{NewSSD(), NewSSD()} ,
		Screen : NewScreen(),
		Keyboard : NewKeyboard(),
		Weight : &pb.Laptop_WeightKg{
			WeightKg : randomFloat64(1.0,3.0),
		},
		PriceUsd :randomFloat64(1500,3500) ,
		ReleaseYear :uint32(randomInt(2015,2019)) ,
		UpdatedAt : ptypes.TimestampNow(),
	}

	return newLaptop
}

func RandomLaptopScore() float64 {
	return float64(randomInt(1, 10))
}