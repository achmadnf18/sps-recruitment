package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type MSX []map[string]interface{}

type RDBMS struct {
	Name    string
	Adapter *sqlx.DB
}

type Rows struct {
	ResultSet   *sqlx.Rows
	Query       string
	QueryParams []interface{}
}

// TRANSACTION
type Tx struct {
	Trans *sqlx.Tx
}

// create new postgresql connection to localhost
func NewPgConn(user, db string) *RDBMS {
	opt := `user=` + user + ` dbname=` + db + ` sslmode=disable`
	conn := sqlx.MustConnect(`postgres`, opt)
	//conn.DB.SetMaxIdleConns(1)  // according to http://jmoiron.github.io/sqlx/#connectionPool
	conn.DB.SetMaxOpenConns(61) // TODO: change this according to postgresql.conf -3
	name := `pg::` + user + `@/` + db
	return &RDBMS{
		Name:    name,
		Adapter: conn,
	}
}


func (r *Rows) ErrorCheck(err error, msg string) {
	if len(r.QueryParams) == 0 {
		fmt.Printf("failed ",msg, r.Query)
	} else {
		fmt.Printf("failed ",msg, r.Query, r.QueryParams)
	}
}
func (r *Rows) Next() bool {
	return r.ResultSet.Next()
}
func (r *Rows) Close() {
	r.ResultSet.Close()
}
func (r *Rows) ScanMap() map[string]interface{} {
	res := map[string]interface{}{}
	err := r.ResultSet.MapScan(res)
	r.ErrorCheck(err, `MapScan`)
	for k, v := range res {
		bs, ok := v.([]uint8)
		if ok {
			res[k] = string(bs)
		}
	}
	return res
}

func (db *RDBMS) QMapArray(query string, params ...interface{}) MSX {
	rows := db.QAll(query)
	res := MSX{}
	defer rows.Close()
	for rows.Next() {
		res = append(res, rows.ScanMap())
	}
	return res
}

func (db *RDBMS) QAll(query string, params ...interface{}) (rows Rows) {
	var err error
	rs, err := db.Adapter.Queryx(query, params...)
	if err != nil{
		fmt.Printf(`failed to QAll: %s %# v`, query, params)
		return
	}

	rows = Rows{rs, query, params}
	return
}

// begin, commit transaction and rollback automatically when there are error
func (db *RDBMS) DoTransaction(lambda func(tx *Tx) string) {
	tx := db.Adapter.MustBegin()
	ok := ``
	defer func() {
		str := recover()
		if str != nil {
			// rollback when error
			err := tx.Rollback()
			if err != nil {
				fmt.Printf(`failed to end transaction %# v / %# v`, str, err)
			}
			return
		}
		if ok == `` {
			// commit when empty string or nil
			err := tx.Commit()
			if err != nil {
				fmt.Printf(`failed to commit transaction`)
			}
			return
		}
		fmt.Println(ok)
		tx.Rollback() // rollback when there is a string
	}()
	ok = lambda(&Tx{tx})
}

// execute anything that doesn't need LastInsertId or RowsAffected
func (tx *Tx) DoExec(query string, params ...interface{}) sql.Result {
	res, err := tx.Trans.Exec(query, params...)
	if err != nil {
		fmt.Println(err, query)
	}
	return res
}

func (tx *Tx) DoExecGetLastID(query string, params ...interface{}) int64{
	var last_id int64
	err := tx.Trans.QueryRow(query + ` RETURNING id`, params...).Scan(&last_id)
	if err != nil {
		fmt.Println(err, query)
		return 0
	}
	return last_id
}