package database

import (
	"database/sql"
	"fajar7xx/pzn-golang-restful-api/helper"
	"time"
)

// kalau struct set pointer untuk returnnya
func NewDB() *sql.DB{
	db, err := sql.Open("mysql", "fajarsiagian:tekanenter8x@tcp(127.0.0.1:3306)/pzn_golang_db")
	helper.PanicIfError(err)

	// set db connection pool
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}