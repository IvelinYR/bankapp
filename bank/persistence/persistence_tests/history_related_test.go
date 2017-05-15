package persistance_test

import (
	"github.com/iliyanmotovski/bankv3/bank/domain"
	"github.com/iliyanmotovski/bankv3/bank/persistence"
	"gopkg.in/mgo.v2"
	"testing"
	"time"
)

func TestItReadsCorrectHistoryDataFromDB(t *testing.T) {
	session, _ := mgo.DialWithTimeout("localhost:27016", time.Second*10)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	sessionStore := persistence.NewMongoSessionStore(*session, "bankTest")

	testHistory := domain.History{AccountID: "12345", UserID: "54321", TransactionType: "testT", Currency: "testC", Amount: 10}

	err := session.DB("bankTest").C("history").Insert(&domain.History{AccountID: testHistory.AccountID, UserID: testHistory.UserID, TransactionType: testHistory.TransactionType, Currency: testHistory.Currency, Amount: testHistory.Amount})
	if err != nil {
		t.Error(err.Error())
	}

	h, err := sessionStore.GetHistory(testHistory)
	if err != nil {
		t.Error(err.Error())
	}

	historyFromDB := *h

	if historyFromDB[0].AccountID != testHistory.AccountID || historyFromDB[0].UserID != testHistory.UserID || historyFromDB[0].TransactionType != testHistory.TransactionType || historyFromDB[0].Currency != testHistory.Currency || historyFromDB[0].Amount != testHistory.Amount {
		t.Errorf("expected historyFromDB AccountID='%s' UserID='%s' TransactionType='%s' Currency='%s' Amount='%f'\n"+
			"got '%s' '%s' '%s' '%s' '%f'", testHistory.AccountID, testHistory.UserID, testHistory.TransactionType, testHistory.Currency, testHistory.Amount,
			historyFromDB[0].AccountID, historyFromDB[0].UserID, historyFromDB[0].TransactionType, historyFromDB[0].Currency, historyFromDB[0].Amount)
	}

	sessionStore.Session.DB("bankTest").C("history").RemoveAll(nil)
}
