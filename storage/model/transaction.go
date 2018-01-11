package model

import "strconv"

type Transaction struct {
	Hash             string `json:"hash"`
	Block            string `json:"blockNumber"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Input            string `json:"input"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

func (t Transaction) GetTransactionIndex() int {
	index, _ := strconv.ParseInt(t.TransactionIndex[2:], 16, 64)
	return int(index)
}

func (t Transaction) GetValue() (value int64) {
	value, _ = strconv.ParseInt(t.Value[2:], 16, 64)
	return
}

func (t Transaction) GetValueEth() (value float64) {
	return float64(t.GetValue()) / 1000000000000000000
}

func (t Transaction) GetFrom() (from string) {
	if len(t.From) > 2 && t.From[0:2] == "0x" {
		from = t.From[2:]
	} else {
		from = ""
	}
	return
}

func (t Transaction) GetTo() (to string) {
	if len(t.To) > 2 && t.To[0:2] == "0x" {
		to = t.To[2:]
	} else {
		to = ""
	}
	return
}
