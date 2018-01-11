package model

import (
	"container/heap"
	"fmt"
	"math"
	"testing"
)

func TestSortedTransactions_PopLimited(t *testing.T) {
	items := map[string]BlockNumber{
		"banana": 3, "apple": 2, "pear": 4,
		"banana1": 3, "apple1": 2, "pear1": 4,
		"banana2": 3, "apple2": 2, "pear2": 4,
		"banana3": 3, "apple3": 2, "pear3": 4,
		"banana4": 5, "apple4": 6, "pear4": 7,
		"banana5": 5, "apple5": 6, "pear5": 7,
		"banana6": 5, "apple6": 6, "pear6": 7,
	}

	set := make(SortedTransactions, len(items))
	i := 0
	for value, priority := range items {
		set[i] = &LastTransaction{
			To:     value,
			Amount: math.Sqrt(float64(priority)),
			Date:   "20061212",
			block:  priority,
			index:  i,
		}
		i++
	}
	heap.Init(&set)

	// Insert a new item and then modify its Block.
	item := &LastTransaction{
		To:     "0x123123123123123213",
		Amount: 1.23,
		Date:   "20061212",
		block:  8,
	}
	heap.Push(&set, item)

	for set.Len() > 0 {
		item := heap.Pop(&set).(*LastTransaction)
		fmt.Printf("%.2d:%s %f", item.block, item.To, item.Amount)
	}
}
