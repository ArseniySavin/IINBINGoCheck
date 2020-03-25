package IINBINGoCheck

import (
	"encoding/json"
	"fmt"
	"github.com/ArseniySavin/IINBINGoCheck/types"
	"time"
)

func ExampleHello() {
	fmt.Println("hello")
	fmt.Println(types.Bin.String())
	a := types.Iin
	fmt.Println(a.String())

	d := types.Individual

	fmt.Print(d.String())

	model := types.IinModel{
		BaseModel: types.BaseModel{
			SequenceNumber: "12345",
			Rank:           10,
			DocType:        types.Iin,
		},
		BirthDate: time.Now(),
		Gender:    types.Female,
		Century:   20,
	}
	m, _ := json.Marshal(&model)
	fmt.Println(string(m))
	//--Output: hello
}
