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

func TestUserAccountIsDeleted(t *testing.T) {
	r := httptest.NewRequest("DELETE", "/delete-account?id=12345", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().DeleteAccount(gomock.Any(), gomock.Any()).Return(nil)

	context.Set(r, "session", &domain.Session{})
	api.DeleteUserAccount(mockAccountStore).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status: expected %d, got %d", http.StatusOK, w.Code)
	}
}

func TestIfNoUrlArguments(t *testing.T) {
	r := httptest.NewRequest("DELETE", "/delete-account", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	context.Set(r, "session", &domain.Session{})
	api.DeleteUserAccount(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Delete User Account Failed","Resource":"request","Field":"URL_parameters","Code":"need_accountID_to_be_specified_in_URL"}`)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status: expected %d, got %d", http.StatusBadRequest, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestIfThereIsAnErrorWhileDeletingTheAccount(t *testing.T) {
	r := httptest.NewRequest("DELETE", "/delete-account?id=12345", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().DeleteAccount(gomock.Any(), gomock.Any()).Return(errors.New("persistence error while deleting the account"))

	context.Set(r, "session", &domain.Session{})
	api.DeleteUserAccount(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Delete User Accounts Failed","Resource":"error","Field":"unexpected_error","Code":"persistence error while deleting the account"}`)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status: expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}
