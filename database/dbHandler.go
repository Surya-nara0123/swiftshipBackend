package database

import (
	"database/sql"
	"fmt"

	"github.com/surya-nara0123/swiftship/helperfunction"
)

type DatabaseInterface interface {
	OpenConn()
	CloseConn()
}

type DatabaseStruct struct {
	connString string
	db         *sql.DB
	err        error
}

func (s *DatabaseStruct) OpenConn() (*sql.DB, error) {
	s.connString = "postgresql://postgres:@localhost:5432/swiftship?sslmode=disable"
	s.db, s.err = sql.Open("postgres", s.connString)
	helperfunction.CheckError(s.err)
	s.err = s.db.Ping()
	helperfunction.CheckError(s.err)

	fmt.Println("Connected!")
	return s.db, s.err
}

func (s *DatabaseStruct) CloseConn() {
	s.db.Close()
}

func (s *DatabaseStruct) GetDbData() (*sql.DB, string) {
	return s.db, s.connString
}
