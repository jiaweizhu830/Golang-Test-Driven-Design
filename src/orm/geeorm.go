package geeorm

import (
	"database/sql"
	"tdd-go/src/orm/log"
	"tdd-go/src/orm/session"
)

/*
	User interface
*/

type Engine struct {
	db *sql.DB
}

// driver: db driver
// source: db name
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)

	if err != nil {
		log.Error(err)
		return
	}

	// Send a ping to make sure database connection is alive
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	e = &Engine{db: db}
	log.Info("Connect to database successfully")
	return
}

// Close: method of Engine
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database successfully")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
