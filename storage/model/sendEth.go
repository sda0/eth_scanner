package model

import "strconv"

type SendEth struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func (t SendEth) GetWei() int64 {
	return int64(t.Amount * 1000000000000000000)
}

func (t SendEth) GetWeiHexed() string {
	return "0x" + strconv.FormatInt(t.GetWei(), 16)
}
