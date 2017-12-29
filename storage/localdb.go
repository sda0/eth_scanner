package storage

import "./dbconnect"

type LocalDB struct {
	connect *dbconnect.PostgresConnection
}
