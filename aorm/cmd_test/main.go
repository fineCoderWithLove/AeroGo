package main

import (
	"aorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine, _ := aorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/aero?charset=utf8mb4&parseTime=True&loc=Local")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
	//pass

}
