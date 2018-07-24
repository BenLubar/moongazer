package gamedata

import (
	"bytes"
	"strconv"
)

type FloatInt int

var pointZero = []byte{'.', '0'}

func (fi *FloatInt) MarshalJSON() ([]byte, error) {
	return append(strconv.AppendInt(nil, int64(*fi), 10), pointZero...), nil
}

func (fi *FloatInt) UnmarshalJSON(b []byte) error {
	i, err := strconv.Atoi(string(bytes.TrimSuffix(b, pointZero)))
	if err == nil {
		*fi = FloatInt(i)
	}
	return err
}
