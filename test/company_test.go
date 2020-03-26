package test

import (
	"github.com/ArseniySavin/IINBINGoCheck/types"
	"testing"
)

func TestCompanyType_ValueType_Less_Unk(t *testing.T) {
	a := types.CompanyType(-2)

	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestCompanyType_Equal_Unk(t *testing.T) {
	a := types.CompanyType(types.Unk)
	res := a.String()

	if res != "Unk" {
		t.Errorf("want to \"Unk\" but result %s", res)
	}

	a = -1

	if res != "Unk" {
		t.Errorf("want to \"Unk\" but result %s", res)
	}
}

func TestCompanyType_LegalResident(t *testing.T) {
	a := types.CompanyType(types.LegalResident)

	res := a.String()

	if res != "LegalResident" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestCompanyType_LegalNoResident(t *testing.T) {
	a := types.CompanyType(types.LegalNoResident)

	res := a.String()

	if res != "LegalNoResident" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestCompanyType_Individual(t *testing.T) {
	a := types.CompanyType(types.Individual)

	res := a.String()

	if res != "Individual" {
		t.Errorf("want to Unk but result %s", res)
	}
}
