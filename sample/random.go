package sample

import(
	"math/rand"
	"github.com/tiwariarjun91/PC-BookApp/pb"

)
func randomKeyBoardLayout() pb.Keyboard_Layout{
	switch rand.Intn(3){
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}

}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomStringFromSet(a ...string) string{
	n := len(a)
	if n==0{
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUBrand() string{
	return randomStringFromSet("Intel","AMD")
}

func randomCPUName(brand string) string{
	if brand == "Intel"{
		return randomStringFromSet(
		"Xeon E-2286M",
		"Core i9-9980HK",
		"Core i7-9750H",
		"Core i5-9400F",
		"Core i3-1005G1"
	)
	}

	return randomStringFromSet(
		"Ryzen 7 Pro 2700U",
		"Ryzen 5 Pro 3500U",
		"Ryzen 3 Pro 3200GE"
	)
	
}

func randomInt(min,max int) int{
	return min + rand.Intn(max-min+1)
}