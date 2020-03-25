package types

import "time"

type IinModel struct {
	BaseModel
	BirthDate time.Time
	Gender    GenderType
	Century   int
}
