package types

type CompanyType int

const (
	LegalResident   CompanyType = 4
	LegalNoResident CompanyType = 5
	Individual      CompanyType = 6
)

func (i CompanyType) String() string {
	switch i {
	case LegalResident:
		return "LegalResident"
	case LegalNoResident:
		return "LegalNoResident"
	case Individual:
		return "Individual"
	default:
		return "Unk"
	}

}
