package iinbindata

type DocType int

const (
	IIN DocType = 0
	BIN DocType = 1
)

func (i DocType) String() string {
	enums := [...]string{"Iin", "Bin"}

	if -1 >= i || int(i) >= len(enums) {
		return "Unk"
	}

	return enums[i]
}
