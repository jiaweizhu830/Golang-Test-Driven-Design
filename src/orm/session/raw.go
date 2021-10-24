package session

import (
	"database/sql"
	"strings"
	"tdd-go/src/orm/log"
)

/*
	Interaction with Database
	1. add custom logging
	2. after exec, clear sql and sqlVars => Session can be reused
*/

// *sql.DB: pointer returned from sql.Open() (connect to database)
type Session struct {
	db  *sql.DB
	sql strings.Builder
	// An empty interface may hold values of any type.
	sqlVars []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) DB() *sql.DB {
	return s.db
}

// method for user to change sql and sqlVars
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	// append works on nil slices.
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()

	log.Info(s.sql.String(), s.sqlVars)

	// first is initialization statement (result & err are if block scope)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()

	log.Info(s.sql.String(), s.sqlVars)

	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// rows and err are named in return => treated as variables defined at the top of the function
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()

	log.Info(s.sql.String(), s.sqlVars)

	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
