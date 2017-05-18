package persistence_test

import (
	"github.com/iliyanmotovski/bankv1/bank/domain"
	"github.com/iliyanmotovski/bankv1/bank/persistence"
	"gopkg.in/mgo.v2"
	"testing"
	"time"
)

func TestItReadsCorrectHistoryDataFromDB(t *testing.T) {
	session, _ := mgo.DialWithTimeout("localhost:27016", time.Second*10)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	sessionStore := persistence.NewMongoSessionStore(*session, "bankTest")

	expected := domain.History{AccountID: "12345", UserID: "54321", TransactionType: "testT", Currency: "testC", Amount: 10}

	err := session.DB("bankTest").C("history").Insert(&domain.History{AccountID: expected.AccountID, UserID: expected.UserID, TransactionType: expected.TransactionType, Currency: expected.Currency, Amount: expected.Amount})
	if err != nil {
		t.Error(err.Error())
	}

	h, err := sessionStore.GetHistory(expected)
	if err != nil {
		t.Error(err.Error())
	}

	history := *h

	if history[0] != expected {
		t.Errorf("expected history: %v got: %v", expected, history[0])
	}

	sessionStore.Session.DB("bankTest").C("history").RemoveAll(nil)
}
