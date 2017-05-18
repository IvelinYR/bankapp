package persistence_test

import (
	"github.com/iliyanmotovski/bankv1/bank/domain"
	"github.com/iliyanmotovski/bankv1/bank/persistence"
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

	expected := domain.UserRegistrationRequest{Username: "testuser", Name: "testname", Email: "testemail", Age: 10}

	_, err := sessionStore.RegisterUser(expected)
	if err != nil {
		t.Error(err.Error())
	}

	var user domain.User
	err = session.DB("bankTest").C("users").Find(bson.M{"username": expected.Username}).One(&user)
	if err != nil {
		t.Error(err.Error())
	}

	if user.Username != expected.Username || user.Name != expected.Name || user.Email != expected.Email || user.Age != expected.Age {
		t.Errorf("expected userFromDB Username='%s' Name='%s' Email='%s' Age='%d'\n"+
			"got '%s' '%s' '%s' '%d'", expected.Username, expected.Name, expected.Email, expected.Age,
			user.Username, user.Name, user.Email, user.Age)
	}

	sessionStore.Session.DB("bankTest").C("users").RemoveAll(nil)
}
