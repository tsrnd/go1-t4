package database

import (
	"fmt"
	"strings"
	"database/sql"

)

func MappingPointer(colsName []string, colsScan map[string]interface{}) (result []interface{}){
	for _, colName := range colsName {
		for key, val := range colsScan {
			if key == colName {
				result = append(result, val)
			}
		}
	}
	return
}

func (db *DB) ClearAttr() {
	db.Condition = ""
	db.ColsSelect = ""
	for i := 0; i<len(db.ColsScan); i++ {
		db.ColsScan[i] = ""
	}
}

func ScanInto(rows *sql.Rows, m []interface{}) (err error) {
	for rows.Next() {
		err = rows.Scan(m...); if err != nil {
			return
		}
	}
	return
}

func (db *DB) Where(stmt string, cond ...interface{}) (*DB) {
	for _, v := range cond {
		stmt = strings.Replace(stmt, "?", fmt.Sprintf("%v", v), 1)
	}
	db.Condition = "WHERE " + stmt
	return db
}

