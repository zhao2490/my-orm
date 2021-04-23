package main

import (
	"fmt"
	"testing"

	"github.com/zhao2490/my-orm/engine"
)

type User struct {
	Name string `gorm:"PRIMARY KEY"`
	Age  int
}

func TestSession_CreateTable(t *testing.T) {
	e, err := engine.NewEngine("sqlite3", "myorm.db")
	if err != nil {
		t.Fatal(err)
	}

	s := e.NewSession().Model(&User{})
	fmt.Println("hasTable", s.HasTable())
	_ = s.DropTable()
	fmt.Println("hasTable", s.HasTable())
	_ = s.CreateTable()
	fmt.Println("hasTable", s.HasTable())
}
