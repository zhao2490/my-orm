package session

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zhao2490/my-orm/dialect"
	"github.com/zhao2490/my-orm/log"
)

type User struct {
	Name string
	Age  int
}

var (
	user1 = &User{"tom", 18}
	user2 = &User{"sam", 25}
	user3 = &User{"jack", 25}
)

func NewSession() *Session {
	driver := "sqlite3"
	db, err := sql.Open(driver, "/Users/edz/Documents/work/gopath/src/github.com/zhao2490/my-orm/myorm.db")
	if err != nil {
		log.Error(err)
		return nil
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return nil
	}
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("can't got dialect for %s", driver)
		return nil
	}
	return New(db, dial)
}

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	affected, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	//t.Log(err1)
	t.Log("insert num ", affected)
	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil {
		t.Fatal(err)
	}
	t.Log(users)
}

func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Limit(1).Find(&users)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result = %+v", users)
}

func TestSession_Update(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("Name = ?", "tom").Update("Age", 30)
	u := &User{}
	_ = s.OrderBy("Age DESC").First(u)

	if affected != 1 || u.Age != 30 {
		t.Fatal("failed to update")
	}
}

func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("Name = ?", "tom").Delete()
	if affected != 1 {
		t.Fatal("failed to delete")
	}
	count, _ := s.Count()
	if count != 1 {
		t.Fatal("failed to count")
	}
}
