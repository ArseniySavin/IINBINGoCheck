package internal

import (
	"fmt"
	"math"
	"time"
)

func SetBirthDateUsingCentury(c, y1, y2, m1, m2, d1, d2 int) time.Time {
	century := 0

	switch c {
	case 1, 2:
		century = 18
	case 3, 4:
		century = 19
	case 5, 6:
		century = 20
	}

	birthDate, _ := time.Parse("2006-02-01",
		fmt.Sprintf("%d%d%d-%d%d-%d%d",
			century, y1, y2, m1, m2, d1, d2))

	return birthDate
}

func SetGender(g int) int {
	res := g % 2
	math.Abs(float64(res))
	if math.Abs(float64(res)) == 1 {
		return 0
	}
	return 1
}

func SetCentury(c int) int {
	century := 0
	switch c {
	case 1, 2:
		century = 18
	case 3, 4:
		century = 19
	case 5, 6:
		century = 20
	}

	return century
}

func SetRegistrationDate(y1, y2, m1, m2 int) time.Time {
	regDate, _ := time.Parse("06-01", fmt.Sprintf("%d%d-%d%d", y1, y2, m1, m2))
	return regDate
}
