package types

const (
	Unk GenderType = -1
)

type BaseModel struct {
	SequenceNumber string
	Rank           int
	DocType
}
