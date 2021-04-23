package schema

import (
	"go/ast"
	"reflect"

	"github.com/zhao2490/my-orm/dialect"
)

// Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag  string
}

// Schema represents a table of database
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	FieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.FieldMap[name]
}

// 将任意的对象解析为 Schema 实例
func Parse(dest interface{}, dialect dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		FieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				Type: dialect.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("gorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, field.Name)
			schema.FieldMap[field.Name] = field
		}
	}
	return schema
}
