package examples

import (
	"fmt"

	"github.com/ArseniySavin/IINBINGoCheck/iinbincheker"
	"github.com/ArseniySavin/IINBINGoCheck/iinbindata"
)

func BinExample(idn string) {
	fmt.Println("BIN-VALIDATION")

	chek, err := iinbincheker.NewValidator(idn)

	if err != nil {
		panic(err)
	}

	res := chek.DirectValidating()
	fmt.Println(res)

	tbin, doc := chek.Identify()
	fmt.Println(tbin)
	fmt.Println(doc)

	bin, ok := doc.(*iinbindata.Bin)
	if !ok {
		recover()
	}

	fmt.Print("Data: ")
	binData := bin.Data()
	fmt.Println(bin.Data())

	fmt.Print("Claims bin ")
	fmt.Println(binData.SequenceNumber)

	fmt.Print("Data to string: ")
	fmt.Println(bin.Data().String())
}
