package sample
import(
	"github.com/tiwariarjun91/PC-BookApp/pb"
)
func NewKeyboard() *pb.Keyboard{
	keyboard := &pb.Keyboard{
		Layout: randomKeyBoardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

func NewCPU() *pb.CPU{
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2,8)
	numberThreads := randomInt(numberCores, 12)
	cpu := &pb.CPU{
		Brand: brand,
		Name: name,
		NumberCores: uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		

	}
	return cpu
}