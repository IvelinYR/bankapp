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

func TestUserWithdrawsSuccessfully(t *testing.T) {
	r := httptest.NewRequest("PATCH", "/withdraw", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().Withdraw(gomock.Any()).Return(nil)

	context.Set(r, "session", &domain.Session{})
	api.UserAccountWithdraw(mockAccountStore).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status: expected %d, got %d", http.StatusOK, w.Code)
	}
}

func TestIfUserTriesToWithdrawMoreThanAvailableInTheAccount(t *testing.T) {
	r := httptest.NewRequest("PATCH", "/withdraw", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().Withdraw(gomock.Any()).Return(domain.ErrWithdrawMoreThanHave)

	context.Set(r, "session", &domain.Session{})
	api.UserAccountWithdraw(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"User Account Withdraw Failed","Resource":"account","Field":"money_availability","Code":"cannot_withdraw_more_than_have_available"}`)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status: expected %d, got %d", http.StatusBadRequest, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestIfThereIsAnErrorWhileWithdrawing(t *testing.T) {
	r := httptest.NewRequest("PATCH", "/withdraw", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().Withdraw(gomock.Any()).Return(errors.New("persistence error while withdrawing"))

	context.Set(r, "session", &domain.Session{})
	api.UserAccountWithdraw(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"User Account Withdraw Failed","Resource":"error","Field":"unexpected_error","Code":"persistence error while withdrawing"}`)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status: expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}
