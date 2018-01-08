package model

import "strconv"

type Block struct {
	Transactions []Transaction `json:"transactions"`
	Timestamp    string        `json:"timestamp"`
	Number       string        `json:"number"`
}

func (b Block) GetNumber() BlockNumber {
	i, _ := strconv.ParseInt(b.Number[2:], 16, 64)
	return BlockNumber(i)
}

func (b Block) GetTimestamp() int64 {
	timestamp, _ := strconv.ParseInt(b.Timestamp[2:], 16, 64)
	return timestamp
}
