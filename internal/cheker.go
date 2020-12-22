package internal

func Check(idn, alg []int) bool {
	chekSum := 0
	for i := 0; i < 11; i++ {
		chekSum = (idn[i] * alg[i]) + chekSum
	}

	chekSumMod := (chekSum % 11)

	if chekSumMod == idn[11] {
		return true
	} else {
		return false
	}
}
