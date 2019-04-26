package strconvutil

import "strconv"

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return i
}

// BaseConvert convert a number between arbitrary bases
func BaseConvert(num string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(num, base, 64)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(i, toBase), nil
}

// Bindec convert a binary string to decimal
func Bindec(bin string) (int, error) {
	i, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return 0, err
	}

	return int(i), err
}
