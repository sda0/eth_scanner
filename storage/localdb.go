package storage

import (
	"fmt"
	"github.com/sda0/eth_scanner/storage/dbconnect"
	"github.com/sda0/eth_scanner/storage/model"
	"log"
)

type LocalDB struct {
	connect *dbconnect.PostgresConnection
}

func (db LocalDB) GetLastBlockNumber() (blockNumber int64) {
	connection, err := db.connect.GetConnection()
	if err != nil {
		panic(err)
	}
	rows, err := connection.Queryx(`SELECT coalesce((SELECT max(blockNumber) FROM transactions),0)`)
	defer rows.Close()

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		values, err := rows.SliceScan()
		if err != nil {
			panic(err)
		}
		blockNumber = values[0].(int64)
		break
	}
	return
}

func (db LocalDB) Save(block model.Block) (affected int, err error) {
	connection, err := db.connect.GetConnection()
	if err != nil {
		panic(err)
	}
	tx := connection.MustBegin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	var query string

	timestamp := block.GetTimestamp()
	blockNumber := block.GetNumber()

	for _, t := range block.Transactions {

		query = `
			INSERT INTO transactions (blockNumber, timestamp, hash, transactionIndex, "from", "to", value, input, v, r, s)
			VALUES(?,?,?,?,?,?,?,?,?,?,?)
			ON CONFLICT DO NOTHING`
		_, err = tx.Exec(tx.Rebind(query), blockNumber, timestamp, t.Hash, t.GetTransactionIndex(), t.GetFrom(), t.GetTo(), t.GetValue(), t.Input, t.V, t.R, t.S)
		if err != nil {
			return
		}
		affected++
	}

	return
}

func (db LocalDB) GetIncomeTransactionsSinceBlock(start int64, finish int64) (resultJson string) {

	if start > finish-3 {
		start = finish - 3
	}
	// @todo избавиться от if для cpu prediction

	connection, err := db.connect.GetConnection()
	if err != nil {
		panic(err)
	}

	// для производительности не будем маршалить в свой тип, сразу готовим json в pg
	query := fmt.Sprintf(`
		SELECT COALESCE(array_to_json(array_agg(t)), '{}') FROM (
		  SELECT
			substr(abstime(timestamp)::text,0,11) as date,
			concat('0x', encode("to",'escape')) as address,
			value/1000000000000000000 as amount,
			(%d - blockNumber) AS confirmations
		  FROM transactions
		  WHERE blockNumber >= %d AND "to"<>''
		  ORDER BY blockNumber DESC
		  LIMIT 1000
		) as t`, finish, start)

	//log.Println(query)

	rows, err := connection.Queryx(query)
	if err != nil {
		panic(err)
	}
	log.Println(rows)
	for rows.Next() {
		values, err := rows.SliceScan()
		if err != nil {
			panic(err)
		}
		resultJson = fmt.Sprintf("%s", values[0])
	}
	return
}
