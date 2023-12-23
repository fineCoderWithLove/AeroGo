package schema

import (
	"aorm/dialect"
	"go/ast"
	"reflect"
)

// Field represents a column of database
type Field struct {
	Name string //字段名
	Type string //类型
	Tag  string //约束条件
}

// Schema represents a table of database
type Schema struct {
	Model      interface{}
	Name       string            //表名信息
	Fields     []*Field          //所有字段信息的切片
	FieldNames []string          //所有字段名的切片
	fieldMap   map[string]*Field //所有列名
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// 将任意一个对象解析成一个Schema实例
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			//每个字段新建一个filed实例
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			//找到标签，含有arom就进行解析
			if v, ok := p.Tag.Lookup("aorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}
