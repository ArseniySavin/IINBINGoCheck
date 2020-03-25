package types

type GenderType int

const (
	Male   GenderType = 0
	Female GenderType = 1
)

func (i GenderType) String() string {
	if i == -1 {
		return "Unk"
	}

	return [...]string{"Male", "Female"}[i]
}
