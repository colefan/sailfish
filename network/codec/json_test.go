package codec

import (
	"fmt"
	"testing"
)

type TestMsg struct {
	F1 int
	F2 int
}

func TestJsonEncode(t *testing.T) {
	// var stream bytes.Buffer
	// var decoder *json.Decoder
	// var encoder *json.Encoder
	// decoder = json.NewDecoder(&stream)
	// encoder = json.NewEncoder(&stream)
	// for i := 0; i < 10; i++ {
	// 	var t1 TestMsg
	// 	t1.F1 = i
	// 	t1.F2 = i
	// 	fmt.Printf(" t1 %v \n", t1)
	// 	marshByte, _ := json.Marshal(t1)
	// 	fmt.Printf("Marshal() i %d\n", i)
	// 	fmt.Print((marshByte))
	// 	fmt.Print("\n")
	// 	marshByte2, _ := json.Marshal(&t1)
	// 	fmt.Print("Marsh &\n")
	// 	fmt.Print((marshByte2))
	// 	fmt.Print("\n")
	// 	encoder.Encode(t1)
	// 	var t2 TestMsg
	// 	decoder.Decode(&t2)
	// 	fmt.Printf("t2 = %v\n", t2)

	// }

	xxx := make([]byte, 16, 16)
	fmt.Println(xxx)

}
