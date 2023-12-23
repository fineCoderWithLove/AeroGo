package aorm

import (
	"aorm/log"
	"aorm/sqlsession"
	"database/sql"
)

// 定义engine类
type Engine struct {
	db *sql.DB
}

// 新建一个engine
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}
	//ping
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db}
	log.Info("Connect to database")
	return
}

// 关闭连接
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

// 创建新的连接
func (engine *Engine) NewSession() *sqlsession.Session {
	return sqlsession.New(engine.db)
}
