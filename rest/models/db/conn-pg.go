package db

var PG *RDBMS

func init() {
	PG = NewPgConn(`spacestock`, `spacestock`)
}