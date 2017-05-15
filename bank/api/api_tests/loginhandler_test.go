package api_tests

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/iliyanmotovski/bankv3/bank/api"
	"github.com/iliyanmotovski/bankv3/bank/domain"
	"github.com/iliyanmotovski/bankv3/bank/domain/mock_domain"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestUserLogsInSuccessfully(t *testing.T) {
	r := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(nil, false)
	mockUserStore.EXPECT().Authenticate(gomock.Any()).Return(new(domain.User), nil)
	mockSessionStore.EXPECT().StartSession(gomock.Any(), gomock.Any()).Return(&domain.Session{SessionID: "12345"}, nil)

	api.LoginHandler(mockUserStore, mockSessionStore, time.Second).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status: expected %d, got %d", http.StatusOK, w.Code)
	}
}

func TestIfUserIsAlreadyLoggedIn(t *testing.T) {
	r := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(nil, true)

	api.LoginHandler(mockUserStore, mockSessionStore, time.Second).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Login Failed","Resource":"user","Field":"user_status","Code":"already_logged_in"}`)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status: expected %d, got %d", http.StatusBadRequest, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestIfUserTriesToLoginWithInvalidUsername(t *testing.T) {
	r := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(nil, false)
	mockUserStore.EXPECT().Authenticate(gomock.Any()).Return(nil, domain.ErrUsernameDoesntExist)

	api.LoginHandler(mockUserStore, mockSessionStore, time.Second).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Login Failed","Resource":"user","Field":"username","Code":"username_does_not_exist"}`)

	if w.Code != http.StatusNotFound {
		t.Errorf("wrong status: expected %d, got %d", http.StatusNotFound, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestIfUserIsAlreadyLoggedInAndPersonFromAnotherLocationTriesToLogin(t *testing.T) {
	r := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(nil, false)
	mockUserStore.EXPECT().Authenticate(gomock.Any()).Return(nil, domain.ErrAlreadyLoggedIn)

	api.LoginHandler(mockUserStore, mockSessionStore, time.Second).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Login Failed","Resource":"user","Field":"user_status","Code":"already_logged_in"}`)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status: expected %d, got %d", http.StatusBadRequest, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestIfUsernameIsCorrectButPasswordIsWrong(t *testing.T) {
	r := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(nil, false)
	mockUserStore.EXPECT().Authenticate(gomock.Any()).Return(nil, domain.ErrWrongPassword)

	api.LoginHandler(mockUserStore, mockSessionStore, time.Second).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Login Failed","Resource":"user","Field":"password","Code":"wrong_password"}`)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("wrong status: expected %d, got %d", http.StatusUnauthorized, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestIfThereIsAnErrorWhileStartingTheSession(t *testing.T) {
	r := httptest.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUserStore := mock_domain.NewMockUserStore(ctrl)
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(nil, false)
	mockUserStore.EXPECT().Authenticate(gomock.Any()).Return(new(domain.User), nil)
	mockSessionStore.EXPECT().StartSession(gomock.Any(), gomock.Any()).Return(nil, errors.New("persistence error while starting the session"))

	api.LoginHandler(mockUserStore, mockSessionStore, time.Second).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Session Initializing Failed","Resource":"error","Field":"unexpected_error","Code":"persistence error while starting the session"}`)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status: expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}
