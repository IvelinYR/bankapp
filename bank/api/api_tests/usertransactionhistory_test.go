package api_tests

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/context"
	"github.com/iliyanmotovski/bankv3/bank/api"
	"github.com/iliyanmotovski/bankv3/bank/domain"
	"github.com/iliyanmotovski/bankv3/bank/domain/mock_domain"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestUserIsAuthenticatedAndRequestsHistory(t *testing.T) {
	r := httptest.NewRequest("POST", "/history", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	instant, _ := time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 3, 2013 at 7:54pm (EET)")
	history := []domain.History{}
	history = append(history, domain.History{UserID: "12345", AccountID: "54321", TransactionType: "deposit", Amount: 100, Date: instant})

	mockAccountStore.EXPECT().GetHistory(gomock.Any()).Return(&history, nil)

	context.Set(r, "session", &domain.Session{})
	api.UserTransactionHistory(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`[{"AccountID":"54321","UserID":"12345","TransactionType":"deposit","Currency":"","Amount":100,"Date":"2013-02-03T19:54:00+02:00"}]`)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status: expected %d, got %d", http.StatusOK, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestUserIsNotAuthorizedToSeeTheAccountHistory(t *testing.T) {
	r := httptest.NewRequest("POST", "/history", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().GetHistory(gomock.Any()).Return(nil, domain.ErrUnauthorized)

	context.Set(r, "session", &domain.Session{})
	api.UserTransactionHistory(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Fetch User Account History Failed","Resource":"account","Field":"account_authorization","Code":"unauthorized_to_see_this_account_history"}`)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("wrong status: expected %d, got %d", http.StatusUnauthorized, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestIfThereIsAnErrorWhileFetchingTheRequestedAccountHistory(t *testing.T) {
	r := httptest.NewRequest("POST", "/history", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().GetHistory(gomock.Any()).Return(nil, errors.New("persistence error while fetching the account history"))

	context.Set(r, "session", &domain.Session{})
	api.UserTransactionHistory(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Fetch User Account History Failed","Resource":"error","Field":"unexpected_error","Code":"persistence error while fetching the account history"}`)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status: expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}
