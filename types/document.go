package types

type DocType int

const (
	Iin DocType = 0
	Bin DocType = 1
)

func (i DocType) String() string {
	if i == -1 {
		return "Unk"
	}

	return [...]string{"Iin", "Bin"}[i]
}
