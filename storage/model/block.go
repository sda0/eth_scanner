package model

import (
	"strconv"
	"time"
)

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

func (b Block) GetDate() string {
	timestamp, _ := strconv.ParseInt(b.Timestamp[2:], 16, 64)
	unixTimeUTC := time.Unix(timestamp, 0)
	return unixTimeUTC.Format("2006-01-02")
}
