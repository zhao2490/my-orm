package schema

import (
	"testing"

	"github.com/zhao2490/my-orm/dialect"
)

type User struct {
	Name string `gorm:"PRIMARY KEY"`
	Age  int
}

var TestDial, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	t.Log(schema.FieldMap)
	if schema.GetField("Name").Tag != "PRIMARY KEY" {
		t.Fatal("failed to parse TAG")
	}
}
