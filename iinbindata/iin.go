package iinbindata

import (
	"fmt"
	"strconv"
	"time"

	internal "github.com/ArseniySavin/IINBINGoCheck/internal"
)

// Iin it is document
type Iin struct {
	idn []int
	BaseType
	BirthDate time.Time
	Gender    GenderType
	Century   int
}

// NewIin it represented a new type Iin
func NewIin(idn string) (*Iin, error) {
	idnArray, err := internal.Parse(idn)
	if err != nil {
		return nil, err
	}
	return Iin{}.SetIdn(idnArray), nil
}

// Init set set parsed identity
func (i Iin) SetIdn(idn []int) *Iin {
	i.idn = idn
	return &i
}

// Data document data
func (i Iin) Data() Iin {
	i.BirthDate = internal.SetBirthDateUsingCentury(i.idn[6], i.idn[0], i.idn[1], i.idn[2], i.idn[3], i.idn[4], i.idn[5])
	i.Gender = GenderType(internal.SetGender(i.idn[6]))
	i.Century = internal.SetCentury(i.idn[6])
	i.BaseType.Rank = i.idn[11]
	i.BaseType.SequenceNumber = fmt.Sprintf("%d%d%d%d", i.idn[7], i.idn[8], i.idn[9], i.idn[10])

	return i
}

func (i Iin) String() string {
	iin := ""
	for _, d := range i.idn {
		iin = iin + strconv.Itoa(d)
	}

	return iin
}
