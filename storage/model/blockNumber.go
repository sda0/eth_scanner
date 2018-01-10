package model

import "strconv"

type BlockNumber int64

func (b BlockNumber) String() string {
	return string(b)
}

func (b BlockNumber) ToHex() string {
	return "0x" + strconv.FormatInt(int64(b), 16)
}
