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
)

func TestUserRequestsHisAccountsAndThereAreNoErrors(t *testing.T) {
	r := httptest.NewRequest("GET", "/accounts", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	account := []domain.Account{}
	account = append(account, domain.Account{UserID: "12345", AccountID: "54321", Currency: "BGN", Amount: 100, Type: "VISA"})
	account = append(account, domain.Account{UserID: "54321", AccountID: "12345", Currency: "NGB", Amount: 200, Type: "AMERICAN"})

	mockAccountStore.EXPECT().GetAccounts(gomock.Any()).Return(&account, nil)

	context.Set(r, "session", &domain.Session{})
	api.GetUserAccounts(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`[{"AccountID":"54321","UserID":"12345","Currency":"BGN","Amount":100,"Type":"VISA"},{"AccountID":"12345","UserID":"54321","Currency":"NGB","Amount":200,"Type":"AMERICAN"}]`)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status: expected %d, got %d", http.StatusOK, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestIfThereIsAnErrorWhileFetchingTheAccounts(t *testing.T) {
	r := httptest.NewRequest("GET", "/accounts", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().GetAccounts(gomock.Any()).Return(nil, errors.New("persistence error while fitching user accounts"))

	context.Set(r, "session", &domain.Session{})
	api.GetUserAccounts(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Fetch User Accounts Failed","Resource":"error","Field":"unexpected_error","Code":"persistence error while fitching user accounts"}`)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status: expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}
