package iinbindata

const (
	Unk GenderType = -1
)

type BaseType struct {
	SequenceNumber string
	Rank           int
	DocType
}
