package database

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseInterface interface {
	OpenConn()
	CloseConn()
}

type DatabaseStruct struct {
	connString string
	db         *gorm.DB
	sqlDB      *sql.DB
	err        error
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func (s *DatabaseStruct) OpenConn() (*gorm.DB, error) {
	s.connString = "postgresql://postgres:@localhost:5432/swiftship?sslmode=disable"
	s.sqlDB, s.err = sql.Open("postgres", s.connString)
	CheckError(s.err)
	s.db, s.err = gorm.Open(postgres.New(postgres.Config{
		Conn: s.sqlDB,
	}), &gorm.Config{})
	CheckError(s.err)

	fmt.Println("Connected!")
	return s.db, s.err
}

func (s *DatabaseStruct) CloseConn() {
	s.sqlDB.Close()
	fmt.Println("Connection Closed")
}

func (s *DatabaseStruct) GetDbData() (*gorm.DB, string) {
	return s.db, s.connString
}
