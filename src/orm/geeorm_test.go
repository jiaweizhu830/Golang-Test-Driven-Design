package geeorm

import (
	geeorm "tdd-go/src/orm"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(t *testing.T) *geeorm.Engine {
	t.Helper()
	engine, err := geeorm.NewEngine("sqlite3", "gee.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}
