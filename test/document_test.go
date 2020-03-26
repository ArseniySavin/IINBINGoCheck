package test

import (
	"github.com/ArseniySavin/IINBINGoCheck/types"
	"testing"
)

func TestDocType_Len_Array(t *testing.T) {
	a := types.DocType(3)
	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestDocType_ValueType_Less_Unk(t *testing.T) {
	a := types.DocType(-2)

	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestDocType_Equal_Unk(t *testing.T) {
	a := types.DocType(types.Unk)
	res := a.String()

	if res != "Unk" {
		t.Errorf("want to \"Unk\" but result %s", res)
	}

	a = -1

	if res != "Unk" {
		t.Errorf("want to \"Unk\" but result %s", res)
	}
}

func TestDocType_Bin(t *testing.T) {
	a := types.DocType(types.Bin)

	res := a.String()

	if res != "Bin" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestDocType_Iin(t *testing.T) {
	a := types.DocType(types.Iin)

	res := a.String()

	if res != "Iin" {
		t.Errorf("want to Unk but result %s", res)
	}
}
