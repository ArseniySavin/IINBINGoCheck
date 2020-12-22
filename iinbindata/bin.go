package iinbindata

import (
	"fmt"
	"strconv"
	"time"

	internal "github.com/ArseniySavin/IINBINGoCheck/internal"
)

// Bin it is document
type Bin struct {
	idn []int
	BaseType
	Type               CompanyType
	RegistrationDate   time.Time
	SpecialCompanyType SpecialCompanyType
}

// NewBin it represented a new type Bin
func NewBin(idn string) (*Bin, error) {
	idnArray, err := internal.Parse(idn)
	if err != nil {
		return nil, err
	}
	return Bin{}.SetIdn(idnArray), nil
}

// SetIdn set parsed identity
func (b Bin) SetIdn(idn []int) *Bin {
	b.idn = idn
	return &b
}

// Data document data
func (b Bin) Data() Bin {
	b.RegistrationDate = internal.SetRegistrationDate(b.idn[0], b.idn[1], b.idn[2], b.idn[3])
	b.Type = CompanyType(b.idn[4])
	b.SpecialCompanyType = SpecialCompanyType(b.idn[0])
	b.BaseType.SequenceNumber = fmt.Sprintf("%d%d%d%d%d", b.idn[6], b.idn[7], b.idn[8], b.idn[9], b.idn[10])
	b.BaseType.Rank = b.idn[11]
	return b
}

func (b Bin) String() string {
	iin := ""
	for _, d := range b.idn {
		iin = iin + strconv.Itoa(d)
	}

	return iin
}
