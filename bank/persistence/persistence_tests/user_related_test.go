package persistance_test

import (
	"github.com/iliyanmotovski/bankv3/bank/domain"
	"github.com/iliyanmotovski/bankv3/bank/persistence"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestItWritesCorrectUserDataInDB(t *testing.T) {
	session, _ := mgo.DialWithTimeout("localhost:27016", time.Second*10)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	sessionStore := persistence.NewMongoSessionStore(*session, "bankTest")

	testUserRegRequest := domain.UserRegistrationRequest{Username: "testuser", Name: "testname", Email: "testemail", Age: 10}

	_, err := sessionStore.RegisterUser(testUserRegRequest)
	if err != nil {
		t.Error(err.Error())
	}

	var userFromDB domain.User
	err = session.DB("bankTest").C("users").Find(bson.M{"username": testUserRegRequest.Username}).One(&userFromDB)
	if err != nil {
		t.Error(err.Error())
	}

	if userFromDB.Username != testUserRegRequest.Username || userFromDB.Name != testUserRegRequest.Name || userFromDB.Email != testUserRegRequest.Email || userFromDB.Age != testUserRegRequest.Age {
		t.Errorf("expected userFromDB Username='%s' Name='%s' Email='%s' Age='%d'\n"+
			"got '%s' '%s' '%s' '%d'", testUserRegRequest.Username, testUserRegRequest.Name, testUserRegRequest.Email, testUserRegRequest.Age,
			userFromDB.Username, userFromDB.Name, userFromDB.Email, userFromDB.Age)
	}

	sessionStore.Session.DB("bankTest").C("users").RemoveAll(nil)
}
