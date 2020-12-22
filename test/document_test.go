package test

import (
	"testing"

	"github.com/ArseniySavin/IINBINGoCheck/iinbindata"
)

func TestDocType_Len_Array(t *testing.T) {
	a := iinbindata.DocType(3)
	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestDocType_ValueType_Less_Unk(t *testing.T) {
	a := iinbindata.DocType(-2)

	res := a.String()

	if res != "Unk" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestDocType_Equal_Unk(t *testing.T) {
	a := iinbindata.DocType(iinbindata.Unk)
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
	a := iinbindata.DocType(1)

	res := a.String()

	if res != "Bin" {
		t.Errorf("want to Unk but result %s", res)
	}
}

func TestDocType_Iin(t *testing.T) {
	a := iinbindata.DocType(0)

	res := a.String()

	if res != "Iin" {
		t.Errorf("want to Unk but result %s", res)
	}
}
