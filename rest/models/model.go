package models

import (
	"sps-recruitment/rest/models/db"
	"strconv"
	"strings"
)

type Appartment struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
}


var PG *db.RDBMS

func init()  {
	PG = db.PG
}

func Replace(haystack, needle, gold string) string {
	return strings.Replace(haystack, needle, gold, -1)
}

func Z(str string) string {
	str = strings.TrimSpace(str)
	str = Replace(str, `<`, `&lt;`)
	str = Replace(str, `>`, `&gt;`)
	str = Replace(str, `'`, `&apos;`)
	str = Replace(str, `"`, `&quot;`)
	return `'` + str + `'`
}

func ZI(num int64) string {
	return `'` + strconv.FormatInt(num, 10) + `'`
}

func GetApartment() []map[string]interface{}{
	query := `-- GetApartment
SELECT x1.id
	, x1.apartment_name
FROM apartment x1
WHERE is_deleted = FALSE
`
	return PG.QMapArray(query)
}

func InsertApartment(name string) int64 {
	created_id := int64(0)
	//Create
	PG.DoTransaction(func(tx *db.Tx) string {
		query := ` -- func InsertApartment
INSERT INTO apartment(apartment_name) VALUES($1)`
		last_id := tx.DoExecGetLastID(query,Z(name))
		created_id = last_id
		return ``
	})
	return created_id
}

func UpdateApartment(id int64,name string) (int64, error) {
	var err error
	ra := int64(0)
	//Create
	PG.DoTransaction(func(tx *db.Tx) string {
		query := ` -- func UpdateApartment
UPDATE apartment SET apartment_name=`+Z(name)+`, updated_at=current_timestamp WHERE id=`+ZI(id)
		ra, err = tx.DoExec(query).RowsAffected()
		return ``
	})
	return ra, err
}

func DeleteApartment(id int64) (int64, error) {
	var err error
	ra := int64(0)
	//Create
	PG.DoTransaction(func(tx *db.Tx) string {
		query := ` -- func UpdateApartment
UPDATE apartment SET is_deleted=TRUE, deleted_at=current_timestamp WHERE id=`+ZI(id)
		ra, err = tx.DoExec(query).RowsAffected()
		return ``
	})
	return ra, err
}

