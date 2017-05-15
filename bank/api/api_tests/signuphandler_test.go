package api_tests

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iliyanmotovski/bankv3/bank/api"
	"github.com/iliyanmotovski/bankv3/bank/domain"
	"github.com/iliyanmotovski/bankv3/bank/domain/mock_domain"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestUserSignsUpSuccessfully(t *testing.T) {
	r := httptest.NewRequest("POST", "/signup", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)

	mockUserStore.EXPECT().RegisterUser(gomock.Any()).Return(&domain.User{UserID: "12345", Username: "Test"}, nil)

	api.SignUpHandler(mockUserStore).ServeHTTP(w, r)

	responseUser := domain.ResponseUser{}

	json.NewDecoder(w.Body).Decode(&responseUser)

	if w.Code != http.StatusCreated {
		t.Errorf("wrong status: expected %d, got %d", http.StatusCreated, w.Code)
	}
	if responseUser.UserID != "12345" || responseUser.Username != "Test" {
		t.Errorf("handler returned unexpected body: Got id: %s username: %s .... Want id: %s username: %s",
			responseUser.UserID, responseUser.Username, "12345", "Test")
	}
}

func TestBadUsername(t *testing.T) {
	r := httptest.NewRequest("POST", "/signup", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)

	mockUserStore.EXPECT().RegisterUser(gomock.Any()).Return(nil, domain.ErrUserAlreadyExists)

	api.SignUpHandler(mockUserStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"SignUp Failed","Resource":"user","Field":"username","Code":"already_exist"}`)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status: expected %d, got %d", http.StatusBadRequest, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestBadEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/signup", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)

	mockUserStore.EXPECT().RegisterUser(gomock.Any()).Return(nil, domain.ErrEmailAlreadyExists)

	api.SignUpHandler(mockUserStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"SignUp Failed","Resource":"user","Field":"email","Code":"already_exist"}`)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status: expected %d, got %d", http.StatusBadRequest, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestUnexpectedPersistenceError(t *testing.T) {
	r := httptest.NewRequest("POST", "/signup", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)

	mockUserStore.EXPECT().RegisterUser(gomock.Any()).Return(nil, errors.New("some persistence error"))

	api.SignUpHandler(mockUserStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"SignUp Failed","Resource":"error","Field":"unexpected_error","Code":"some persistence error"}`)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status: expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}
