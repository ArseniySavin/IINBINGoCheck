package test

import (
	"testing"

	"github.com/ArseniySavin/IINBINGoCheck/iinbindata"
)

func TestSpecialCompanyType_Len_Array(t *testing.T) {
	a := iinbindata.SpecialCompanyType(5)
	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_ValueType_Less_Unk(t *testing.T) {
	a := iinbindata.SpecialCompanyType(-2)

	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_Equal_Unk(t *testing.T) {
	a := iinbindata.SpecialCompanyType(iinbindata.Unk)
	res := a.String()

	if res != "Unk" {
		t.Errorf("want to \"Unk\" but result %s", res)
	}

	a = -1

	if res != "Unk" {
		t.Errorf("want to \"Unk\" but result %s", res)
	}
}

func TestSpecialCompanyType_HeadOffice(t *testing.T) {
	a := iinbindata.SpecialCompanyType(iinbindata.HeadOffice)

	res := a.String()

	if res != "HeadOffice" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_DepartOffice(t *testing.T) {
	a := iinbindata.SpecialCompanyType(iinbindata.DepartOffice)

	res := a.String()

	if res != "DepartOffice" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_Agency(t *testing.T) {
	a := iinbindata.SpecialCompanyType(iinbindata.Agency)

	res := a.String()

	if res != "Agency" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_Other(t *testing.T) {
	a := iinbindata.SpecialCompanyType(iinbindata.Other)

	res := a.String()

	if res != "Other" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_Farm(t *testing.T) {
	a := iinbindata.SpecialCompanyType(iinbindata.Farm)

	res := a.String()

	if res != "Farm" {
		t.Errorf("want to Unk but result %s", res)
	}
}
