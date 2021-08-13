package sample

import(
	"math/rand"
	"github.com/tiwariarjun91/PC-BookApp/pb"
	"github.com/google/uuid"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano()) // rand has a fixed seed if this bit of code were not here random file would generate the same values except the id which is generated because of uuid
}

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
		"Core i3-1005G1", //needs comma on last line aswell
	)
	}

	return randomStringFromSet(
		"Ryzen 7 Pro 2700U",
		"Ryzen 5 Pro 3500U",
		"Ryzen 3 Pro 3200GE", //needs comma on last line as well
	)
	
}

func randomInt(min,max int) int{
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min,max float64) float64{
	return min + rand.Float64()*(max-min)
}
func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
func randomGPUBrand() string{
	return randomStringFromSet("Nvidia","AMD")
}

func randomId() string {
	return uuid.New().String()
}

func randomGPUName(brand string) string{
	if brand == "Nvidia"{
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}
/* 
func randomStringFromSet(a ...string) string{
	n := len(a)
	if a == 0{
		return ""
	}

	return a[rand.Intn(n)]
}
*/

/*
func RandomInt(min,max int) int{
	return min + rand.Intn(max-min)
}*/

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080,4320)
	width := height * 16 /9

	screenResolution := &pb.Screen_Resolution{
		Height : uint32(height),
		Width : uint32(width),
	}

	return screenResolution
}

func randomScreenPanel() pb.Screen_Panel{
	if rand.Intn(2) == 1{
		return pb.Screen_IPS
	}else{
		return pb.Screen_OLED
	}
}

func randomLaptopBrand() string{
	return randomStringFromSet("Apple","Dell", "Lenovo")
}

func randomLaptopName(brand string) string{
	switch brand{
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	}
}