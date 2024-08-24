package istckimlikgo

import (
	"errors"
	"strconv"
	"strings"
)

func validate(tckimlik string) (valid bool, err error) {

	if tckimlik[0] == '0' {
		err = errors.New("turkish ID can not start with a zero")
		return
	}

	if len(tckimlik) != 11 {
		err = errors.New("turkish ID Number should be 11 characters")
		return
	}

	sum := 0
	splittedID := strings.SplitN(tckimlik, "", 11)
	lastDigit, erz := strconv.Atoi(splittedID[10])
	if erz != nil {
		return
	}

	for i := 0; i < len(tckimlik)-1; i++ {
		toInt, eR := strconv.Atoi(splittedID[i])
		if eR != nil {
			return
		}
		sum += toInt
	}

	if sum%10 != lastDigit {
		err = errors.New("your Turkish ID number is invalid")
		return
	}

	return true, nil
}
