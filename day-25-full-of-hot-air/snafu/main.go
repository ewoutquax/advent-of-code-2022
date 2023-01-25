package snafu

func snafuBase(char string) (value int) {
	switch char {
	case "2":
		value = 2
	case "1":
		value = 1
	case "0":
		value = 0
	case "-":
		value = -1
	case "=":
		value = -2
	}

	return
}
