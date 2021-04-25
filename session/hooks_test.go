package session

import (
	"github.com/zhao2490/my-orm/log"
	"testing"
)

type Account struct {
	ID       int `gorm:"PRIMARY KEY"`
	Password string
}

func (a *Account) BeforeInsert(s *Session) error {
	log.Info("before insert", a)
	a.ID += 1000
	return nil
}

func (a *Account) AfterQuery(s *Session) error {
	log.Info("after query", a)
	a.Password = "********"
	return nil
}

func TestSession_CallMethod(t *testing.T) {
	s := NewSession().Model(&Account{})
	_ = s.DropTable()
	_ = s.CreateTable()
	_, _ = s.Insert(&Account{1, "12345"}, &Account{2, "qwerty"})

	u := &Account{}
	err := s.First(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("u=%+v", u)
}
