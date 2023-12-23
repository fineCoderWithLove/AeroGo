package dialect

import "reflect"

var dialectsMap = map[string]Dialect{}

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(tableName string) (string, []interface{})
}

// 用于将go语言的类型转换成数据库的类型
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

// 返回某个表是否存在的sql，参数是表名
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
