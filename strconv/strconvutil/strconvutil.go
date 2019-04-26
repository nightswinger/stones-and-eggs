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
