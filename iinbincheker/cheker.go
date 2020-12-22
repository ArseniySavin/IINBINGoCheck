package iinbincheker

import (
	"github.com/ArseniySavin/IINBINGoCheck/iinbindata"
	"github.com/ArseniySavin/IINBINGoCheck/internal"
)

var direct = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
var indirect = []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}

// Validation
type Validation interface {
	DirectValidating() bool
	IndirectValidating() bool
}

type validator struct {
	idn   []int
	valid bool
}

// NewValidator it represented new object for validation
func NewValidator(idn string) (*validator, error) {
	idnArray, err := internal.Parse(idn)
	if err != nil {
		return nil, err
	}
	return &validator{idnArray, false}, nil
}

// DirectValidating it is direct algorithm for validating
func (v *validator) DirectValidating() bool {
	v.valid = internal.Check(v.idn, direct)
	return v.valid
}

// IndirectValidating it is indirect algorithm for validating
func (v *validator) IndirectValidating() bool {
	v.valid = internal.Check(v.idn, indirect)
	return v.valid
}

// Identify get new identification with parsed data
func (v *validator) Identify() (iinbindata.DocType, interface{}) {
	if !v.valid {
		return 0, nil
	}

	if v.idn[4] > 3 {
		return iinbindata.BIN, iinbindata.Bin{
			BaseType: iinbindata.BaseType{
				DocType: iinbindata.BIN,
			},
		}.SetIdn(v.idn)
	}
	return iinbindata.IIN, iinbindata.Iin{
		BaseType: iinbindata.BaseType{
			DocType: iinbindata.IIN,
		},
	}.SetIdn(v.idn)
}
