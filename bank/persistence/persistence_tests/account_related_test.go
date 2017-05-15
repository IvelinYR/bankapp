package persistance_test

import (
	"github.com/iliyanmotovski/bankv3/bank/domain"
	"github.com/iliyanmotovski/bankv3/bank/persistence"
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

	testAccount := domain.Account{UserID: "1234", Currency: "testC", Amount: 10}

	accountID, err := sessionStore.InsertAccount(testAccount.UserID, testAccount)
	if err != nil {
		t.Error(err.Error())
	}

	a, err := sessionStore.GetAccounts(testAccount.UserID)
	if err != nil {
		t.Error(err.Error())
	}
	accountFromDB := *a

	if accountFromDB[0].AccountID != accountID || accountFromDB[0].UserID != testAccount.UserID || accountFromDB[0].Currency != testAccount.Currency || accountFromDB[0].Amount != 10 {
		t.Errorf("expected accountFromDB AccoundID='%s' UserID='%s' Currency='%s' Amount='%f'\n"+
			"got '%s' '%s' '%s' '%f'", accountID, testAccount.UserID, testAccount.Currency,
			13.0, accountFromDB[0].AccountID, accountFromDB[0].UserID, accountFromDB[0].Currency, accountFromDB[0].Amount)
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

	err := sessionStore.Deposit(testAccountDeposit)
	if err != nil {
		t.Error(err.Error())
	}

	err = sessionStore.Withdraw(testAccountWithdraw)
	if err != nil {
		t.Error(err.Error())
	}

	accounts := []domain.Account{}
	session.DB(sessionStore.DBName).C("accounts").Find(bson.M{"userid": "1234", "currency": "testC"}).All(&accounts)

	if accounts[0].AccountID != "4321" || accounts[0].UserID != "1234" || accounts[0].Currency != "testC" || accounts[0].Amount != 13 {
		t.Errorf("expected accountFromDB AccoundID='%s' UserID='%s' Currency='%s' Amount='%f'\n"+
			"got '%s' '%s' '%s' '%f'", "4321", "1234", "testC",
			13.0, accounts[0].AccountID, accounts[0].UserID, accounts[0].Currency, accounts[0].Amount)
	}

	err = sessionStore.DeleteAccount("1234", "4321")
	if err != nil {
		t.Error(err.Error())
	}

	session.DB(sessionStore.DBName).C("accounts").Find(bson.M{"userid": "1234", "currency": "testC"}).All(&accounts)

	if len(accounts) != 0 {
		t.Error("expected to have 0 accounts in the DB, but there was more")
	}

	sessionStore.Session.DB("bankTest").C("accounts").RemoveAll(nil)
}
