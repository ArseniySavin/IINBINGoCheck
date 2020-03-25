package types

type SpecialCompanyType int

const (
	HeadOffice   SpecialCompanyType = 0
	DepartOffice SpecialCompanyType = 1
	Agency       SpecialCompanyType = 2
	Other        SpecialCompanyType = 3
	Farm         SpecialCompanyType = 4
)

func (i SpecialCompanyType) String() string {
	if i == -1 {
		return "Unk"
	}
	return [...]string{"HeadOffice", "DepartOffice", "Agency", "Other", "Farm"}[i]
}
