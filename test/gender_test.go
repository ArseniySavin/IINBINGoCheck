package test

import (
	"github.com/ArseniySavin/IINBINGoCheck/types"
	"testing"
)

func TestGenderType_Len_Array(t *testing.T) {
	a := types.GenderType(5)
	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestGenderType_ValueType_Less_Unk(t *testing.T) {
	a := types.GenderType(-2)

	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestGenderType_Equal_Unk(t *testing.T) {
	a := types.GenderType(types.Unk)
	res := a.String()

	if res != "Unk" {
		t.Errorf("want to \"Unk\" but result %s", res)
	}

	a = -1

	if res != "Unk" {
		t.Errorf("want to \"Unk\" but result %s", res)
	}
}

func TestGenderType_Female(t *testing.T) {
	a := types.GenderType(types.Female)

	res := a.String()

	if res != "Female" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestGenderType_Male(t *testing.T) {
	a := types.GenderType(types.Male)

	res := a.String()

	if res != "Male" {
		t.Errorf("want to Unk but result %s", res)
	}
}
