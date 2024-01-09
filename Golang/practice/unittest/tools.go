package unittest

import "errors"

var (
	ErrIndexOutRange = errors.New("exists an error a the values inputs")
)

func GetPositionSlice(index int, n ...int64) (result int64, err error) {
	if len(n) == 0 || index >= len(n) {
		err = ErrIndexOutRange
		return
	}
	result = n[index]
	return
}
