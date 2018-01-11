package model

const maxCapacity = 1000

type LastTransaction struct {
	json          string
	Date          string  `json:"date"`
	To            string  `json:"address"`
	Amount        float64 `json:"amount"`
	Confirmations int     `json:"confirmations"`
	block         BlockNumber
	index         int //for heap.Interface
}

func (ls *LastTransaction) SetBlock(id BlockNumber) {
	ls.block = id
}

func (ls *LastTransaction) GetBlockNumber() BlockNumber {
	return ls.block
}

// heap.Interface
type SortedTransactions []*LastTransaction

func (st SortedTransactions) Len() int { return len(st) }

func (st SortedTransactions) Less(i, j int) bool {
	return st[i].block < st[j].block
}

func (st SortedTransactions) Swap(i, j int) {
	st[i], st[j] = st[j], st[i]
	st[i].index = i
	st[j].index = j
}

func (st *SortedTransactions) Push(x interface{}) {
	n := len(*st)
	item := x.(*LastTransaction)
	item.index = n
	*st = append(*st, item)

	if len(*st) > maxCapacity {
		st.Pop()
	}
}

func (st *SortedTransactions) Pop() interface{} {
	old := *st
	n := len(old)
	item := old[n-1]
	item.index = -1
	*st = old[0 : n-1]
	return item
}

func (st *SortedTransactions) PopLimited(max BlockNumber) *LastTransaction {
	old := *st
	n := len(old)
	if n == 0 {
		return nil
	}
	item := old[n-1]
	if item.block < max {
		item.index = -1
		*st = old[0 : n-1]
		return item
	}
	return nil
}
