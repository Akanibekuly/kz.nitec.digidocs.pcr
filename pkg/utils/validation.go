package utils

func CheckIin(iin string) bool {
	if len(iin) != 12 {
		return false
	}

	for _, number := range iin {
		if number < 48 && number > 57 {
			return false
		}
	}

	month := strToInt(iin[2:4])
	if month < 0 || month > 12 {
		return false
	}

	day := strToInt(iin[4:6])
	if day < 0 || day > 31 {
		return false
	}

	// second faset
	cent := strToInt(iin[6:7])
	if cent < 1 || cent > 6 {
		return false
	}

	// third faset
	uniq := strToInt(iin[7:11])
	if uniq < 0 {
		return false
	}

	//control sum
	contSum := strToInt(iin[11:])

	//
	b1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	b2 := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}
	c := 0
	for i, _ := range b1 {
		c += (int(iin[i]-48) * b1[i])
	}
	c = c % 11
	if c == 10 {
		c := 0
		for i, _ := range b2 {
			c += (int(iin[i]-48) * b2[i])
		}
		c = c % 11
	}
	if c == 10 || c != contSum {
		return false
	}

	return true
}

func strToInt(str string) int {
	result := 0
	for i := range str {
		result *= 10
		result += int(str[i] - 48)
	}
	return result
}
