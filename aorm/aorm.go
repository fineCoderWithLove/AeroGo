package aorm

import (
	"aorm/dialect"
	"aorm/log"
	"aorm/sqlsession"
	"database/sql"
)

// 定义engine类
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}
	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

// 新建立的sqlsession
func (engine *Engine) NewSession() *sqlsession.Session {
	return sqlsession.New(engine.db, engine.dialect)
}
