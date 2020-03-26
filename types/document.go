package types

type DocType int

const (
	Iin DocType = 0
	Bin DocType = 1
)

func (i DocType) String() string {
	enums := [...]string{"Iin", "Bin"}

	if -1 >= i || int(i) >= len(enums) {
		return "Unk"
	}

	return enums[i]
}
