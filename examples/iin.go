package examples

import (
	"fmt"

	"github.com/ArseniySavin/IINBINGoCheck/iinbincheker"
	"github.com/ArseniySavin/IINBINGoCheck/iinbindata"
)

func IinExampl(idn string) {
	fmt.Println("IIN-VALIDATION")

	chek, err := iinbincheker.NewValidator(idn)

	if err != nil {
		panic(err)
	}

	res := chek.DirectValidating()
	fmt.Println(res)

	tiin, doc := chek.Identify()
	fmt.Println(tiin)
	fmt.Println(doc)

	iin, ok := doc.(*iinbindata.Iin)
	if !ok {
		recover()
	}

	fmt.Print("Data: ")
	iinData := iin.Data()
	fmt.Println(iin.Data())

	fmt.Print("Claims iin ")
	fmt.Println(iinData.SequenceNumber)

	fmt.Print("Data to string: ")
	fmt.Println(iin.Data().String())
}
