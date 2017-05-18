package persistence

import (
	"github.com/iliyanmotovski/bankv1/bank/domain"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"time"
)

type MongoSessionStore struct {
	Session mgo.Session
	DBName  string
}

func NewMongoSessionStore(session mgo.Session, dbname string) *MongoSessionStore {
	s := new(MongoSessionStore)
	s.Session = session
	s.DBName = dbname
	return s
}

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func randString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func (se *MongoSessionStore) GetHistory(historyRequest domain.History) (*[]domain.History, error) {
	session := se.Session.Clone()
	defer session.Close()

	var result []domain.History

	err := session.DB(se.DBName).C("history").Find(bson.M{"accountid": historyRequest.AccountID}).All(&result)
	if err != nil {
		return nil, err
	}

	for _, p := range result {
		if p.UserID != historyRequest.UserID {
			return nil, domain.ErrUnauthorized
		}
	}

	return &result, nil
}

func (se *MongoSessionStore) GetAccounts(userID string) (*[]domain.Account, error) {
	session := se.Session.Clone()
	defer session.Close()

	var result []domain.Account

	err := session.DB(se.DBName).C("accounts").Find(bson.M{"userid": userID}).All(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (se *MongoSessionStore) InsertAccount(UserID string, a domain.Account) (string, error) {
	session := se.Session.Clone()
	defer session.Close()

	accountID := randString(7)

	err := session.DB(se.DBName).C("accounts").Insert(&domain.Account{AccountID: accountID, UserID: UserID, Currency: a.Currency, Amount: a.Amount, Type: a.Type})
	if err != nil {
		return "", err
	}

	err = session.DB(se.DBName).C("history").Insert(&domain.History{AccountID: accountID, UserID: UserID, TransactionType: "new account", Currency: a.Currency, Amount: a.Amount, Date: time.Now().Local()})
	if err != nil {
		return "", err
	}

	return accountID, nil
}

func (se *MongoSessionStore) Withdraw(a domain.Account) error {
	session := se.Session.Clone()
	defer session.Close()

	var current domain.Account

	err := session.DB(se.DBName).C("accounts").Find(bson.M{"userid": a.UserID, "currency": a.Currency}).One(&current)
	if err != nil {
		return err
	}

	if current.Amount < a.Amount {
		return domain.ErrWithdrawMoreThanHave
	}
	current.Amount = current.Amount - a.Amount

	err = session.DB(se.DBName).C("accounts").Update(bson.M{"accountid": current.AccountID}, bson.M{"$set": bson.M{"amount": current.Amount}})
	if err != nil {
		return err
	}

	err = session.DB(se.DBName).C("history").Insert(&domain.History{AccountID: current.AccountID, UserID: a.UserID, TransactionType: "withdraw", Currency: a.Currency, Amount: a.Amount, Date: time.Now().Local()})
	if err != nil {
		return err
	}

	return nil
}

func (se *MongoSessionStore) Deposit(a domain.Account) error {
	session := se.Session.Clone()
	defer session.Close()

	var current domain.Account

	err := session.DB(se.DBName).C("accounts").Find(bson.M{"userid": a.UserID, "currency": a.Currency}).One(&current)
	if err != nil {
		return err
	}

	current.Amount = current.Amount + a.Amount

	err = session.DB(se.DBName).C("accounts").Update(bson.M{"accountid": current.AccountID}, bson.M{"$set": bson.M{"amount": current.Amount}})
	if err != nil {
		return err
	}

	err = session.DB(se.DBName).C("history").Insert(&domain.History{AccountID: current.AccountID, UserID: a.UserID, TransactionType: "deposit", Currency: a.Currency, Amount: a.Amount, Date: time.Now().Local()})
	if err != nil {
		return err
	}

	return nil
}

func (se *MongoSessionStore) DeleteAccount(userID string, accountID string) error {
	session := se.Session.Clone()
	defer session.Close()

	var result domain.Account

	err := session.DB(se.DBName).C("accounts").Find(bson.M{"userid": userID, "accountid": accountID}).One(&result)
	if err != nil {
		return err
	}

	err = session.DB(se.DBName).C("accounts").Remove(bson.M{"userid": userID, "accountid": accountID})
	if err != nil {
		return err
	}

	err = session.DB(se.DBName).C("history").Insert(&domain.History{AccountID: accountID, UserID: userID, TransactionType: "account deleted", Currency: result.Currency, Date: time.Now().Local()})
	if err != nil {
		return err
	}

	return nil
}

func (se *MongoSessionStore) StartSession(u domain.User, userSessionDuration time.Time) (*domain.Session, error) {
	session := se.Session.Clone()
	defer session.Close()

	var result domain.Session

	err := session.DB(se.DBName).C("sessions").Insert(&domain.Session{Expires: userSessionDuration, SessionID: randString(30), UserID: u.UserID})
	if err != nil {
		return nil, err
	}
	err = session.DB(se.DBName).C("sessions").Find(bson.M{"userid": u.UserID}).One(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (se *MongoSessionStore) FindSessionAvailableAt(sessionID string, instant time.Time) (*domain.Session, bool) {
	session := se.Session.Clone()
	defer session.Close()

	var result domain.Session

	session.DB(se.DBName).C("sessions").Find(bson.M{"sessionid": sessionID}).One(&result)

	if result.UserID != "" && result.SessionID == sessionID && !instant.After(result.Expires) {
		return &result, true
	}

	return &result, false
}

func (se *MongoSessionStore) DeleteSession(sessionID string) error {
	session := se.Session.Clone()
	defer session.Close()

	err := session.DB(se.DBName).C("sessions").Remove(bson.M{"sessionid": sessionID})
	if err != nil {
		return err
	}
	return nil
}

func (se *MongoSessionStore) UpdateSession(sessionID string, userSessionDuration time.Time) error {
	session := se.Session.Clone()
	defer session.Close()

	err := session.DB(se.DBName).C("sessions").Update(bson.M{"sessionid": sessionID}, bson.M{"$set": bson.M{"expires": userSessionDuration}})
	if err != nil {
		return err
	}
	return nil
}

func (se *MongoSessionStore) RegisterUser(req domain.UserRegistrationRequest) (*domain.User, error) {
	session := se.Session.Clone()
	defer session.Close()
	var err error
	var user domain.User

	session.DB(se.DBName).C("users").Find(bson.M{"username": req.Username}).One(&user)
	if req.Username == user.Username {
		return nil, domain.ErrUserAlreadyExists
	}

	session.DB(se.DBName).C("users").Find(bson.M{"email": req.Email}).One(&user)
	if req.Email == user.Email {
		return nil, domain.ErrEmailAlreadyExists
	}

	user = domain.User{Username: req.Username, UserID: randString(30), Name: req.Name, Email: req.Email, Age: req.Age}

	user.Hash, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = session.DB(se.DBName).C("users").Insert(&domain.User{Username: user.Username, UserID: user.UserID, Name: user.Name, Email: user.Email, Age: user.Age, Hash: user.Hash})
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (se *MongoSessionStore) Authenticate(req domain.UserLoginRequest) (*domain.User, error) {
	session := se.Session.Clone()
	defer session.Close()

	var result domain.User
	err := session.DB(se.DBName).C("users").Find(bson.M{"username": req.Username}).One(&result)
	if err != nil {
		return nil, domain.ErrUsernameDoesntExist
	}

	var sessionResult domain.Session
	session.DB(se.DBName).C("sessions").Find(bson.M{"userid": result.UserID}).One(&sessionResult)

	if result.UserID == sessionResult.UserID {
		// Even though there is functionality preventing multiple active sessions on the same user from different browser tabs,
		// it was possible to log in the same account twice from two different browser instances. This check eliminates this threat.
		return nil, domain.ErrAlreadyLoggedIn
	}

	// Comparing the password with the hash.
	if err = bcrypt.CompareHashAndPassword(result.Hash, []byte(req.Password)); err != nil {
		return nil, domain.ErrWrongPassword
	}
	return &result, nil
}
