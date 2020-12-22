package iinbindata

type GenderType int

const (
	Male   GenderType = 0
	Female GenderType = 1
)

func (i GenderType) String() string {
	enums := [...]string{"Male", "Female"}

	if -1 >= i || int(i) >= len(enums) {
		return "Unk"
	}

	return enums[i]
}
