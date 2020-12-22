package internal

import (
	"fmt"
	"strconv"
)

func Parse(idn string) ([]int, error) {
	ids := make([]int, 0)
	for _, i := range idn {
		toi, err := strconv.Atoi(string(i))
		if err != nil {
			return nil, fmt.Errorf("Cannot parse iin or bin: %s", err)
		}
		ids = append(ids, toi)
	}

	return ids, nil
}
