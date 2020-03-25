package types

import "time"

type BinModel struct {
	BaseModel
	Type               CompanyType
	RegistrationDate   time.Time
	SpecialCompanyType SpecialCompanyType
}
