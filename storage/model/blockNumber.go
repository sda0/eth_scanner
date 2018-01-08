package model

type BlockNumber int64

func (b BlockNumber) String() string {
	return string(b)
}
