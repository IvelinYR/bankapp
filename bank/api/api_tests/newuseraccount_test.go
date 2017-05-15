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

func TestUserOpensNewAccountSuccessfully(t *testing.T) {
	r := httptest.NewRequest("POST", "/new-account", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().InsertAccount(gomock.Any(), gomock.Any()).Return("test", nil)

	context.Set(r, "session", &domain.Session{})
	api.NewUserAccount(mockAccountStore).ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Errorf("wrong status: expected %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestIfThereIsAnErrorWhileInsertingNewAccount(t *testing.T) {
	r := httptest.NewRequest("POST", "/new-account", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAccountStore := mock_domain.NewMockAccountStore(ctrl)

	mockAccountStore.EXPECT().InsertAccount(gomock.Any(), gomock.Any()).Return("test", errors.New("persistence error while inserting new user account"))

	context.Set(r, "session", &domain.Session{})
	api.NewUserAccount(mockAccountStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Create User Account Failed","Resource":"error","Field":"unexpected_error","Code":"persistence error while inserting new user account"}`)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status: expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}
