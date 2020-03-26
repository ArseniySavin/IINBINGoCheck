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
	enums := [...]string{"HeadOffice", "DepartOffice", "Agency", "Other", "Farm"}

	if -1 >= i || int(i) >= len(enums) {
		return "Unk"
	}

	return enums[i]
}
