package test

import (
	"github.com/ArseniySavin/IINBINGoCheck/types"
	"testing"
)

func TestSpecialCompanyType_Len_Array(t *testing.T) {
	a := types.SpecialCompanyType(5)
	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_ValueType_Less_Unk(t *testing.T) {
	a := types.SpecialCompanyType(-2)

	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_Equal_Unk(t *testing.T) {
	a := types.SpecialCompanyType(types.Unk)
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
	a := types.SpecialCompanyType(types.HeadOffice)

	res := a.String()

	if res != "HeadOffice" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_DepartOffice(t *testing.T) {
	a := types.SpecialCompanyType(types.DepartOffice)

	res := a.String()

	if res != "DepartOffice" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_Agency(t *testing.T) {
	a := types.SpecialCompanyType(types.Agency)

	res := a.String()

	if res != "Agency" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_Other(t *testing.T) {
	a := types.SpecialCompanyType(types.Other)

	res := a.String()

	if res != "Other" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestSpecialCompanyType_Farm(t *testing.T) {
	a := types.SpecialCompanyType(types.Farm)

	res := a.String()

	if res != "Farm" {
		t.Errorf("want to Unk but result %s", res)
	}
}
