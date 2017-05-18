package persistence_test

import (
	"github.com/iliyanmotovski/bankv1/bank/domain"
	"github.com/iliyanmotovski/bankv1/bank/persistence"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestItWritesAndReadsAccountDataFromDB(t *testing.T) {
	session, _ := mgo.DialWithTimeout("localhost:27016", time.Second*10)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	sessionStore := persistence.NewMongoSessionStore(*session, "bankTest")

	testAccount := domain.Account{UserID: "1234", Currency: "testC", Amount: 10, Type: "VISA"}

	accountID, err := sessionStore.InsertAccount(testAccount.UserID, testAccount)
	if err != nil {
		t.Error(err.Error())
	}
	testAccount.AccountID = accountID

	a, err := sessionStore.GetAccounts(testAccount.UserID)
	if err != nil {
		t.Error(err.Error())
	}
	account := *a

	if account[0] != testAccount {
		t.Errorf("expected account: %v got: %v", testAccount, account[0])
	}

	sessionStore.Session.DB(sessionStore.DBName).C("accounts").RemoveAll(nil)
}

func TestItUpdatesAndDeletesAccountDataFromDB(t *testing.T) {
	session, _ := mgo.DialWithTimeout("localhost:27016", time.Second*10)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	sessionStore := persistence.NewMongoSessionStore(*session, "bankTest")

	session.DB(sessionStore.DBName).C("accounts").Insert(&domain.Account{UserID: "1234", AccountID: "4321", Currency: "testC", Amount: 10})

	testAccountDeposit := domain.Account{UserID: "1234", Currency: "testC", Amount: 5}
	testAccountWithdraw := domain.Account{UserID: "1234", Currency: "testC", Amount: 2}
	expected := domain.Account{UserID: "1234", AccountID: "4321", Currency: "testC", Amount: 13}

	err := sessionStore.Deposit(testAccountDeposit)
	if err != nil {
		t.Error(err.Error())
	}

	err = sessionStore.Withdraw(testAccountWithdraw)
	if err != nil {
		t.Error(err.Error())
	}

	account := []domain.Account{}
	session.DB(sessionStore.DBName).C("accounts").Find(bson.M{"userid": "1234", "currency": "testC"}).All(&account)

	if account[0] != expected {
		t.Errorf("expected account: %v got: %v", expected, account[0])
	}

	err = sessionStore.DeleteAccount("1234", "4321")
	if err != nil {
		t.Error(err.Error())
	}

	session.DB(sessionStore.DBName).C("accounts").Find(bson.M{"userid": "1234", "currency": "testC"}).All(&account)

	if len(account) != 0 {
		t.Error("expected to have 0 accounts in the DB, but there was more")
	}

	sessionStore.Session.DB(sessionStore.DBName).C("accounts").RemoveAll(nil)
}
