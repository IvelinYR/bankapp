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
)

func TestUserLogsOutSuccessfully(t *testing.T) {
	r := httptest.NewRequest("POST", "/logout", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(new(domain.Session), true)
	mockSessionStore.EXPECT().DeleteSession(gomock.Any()).Return(nil)

	api.LogoutHandler(mockSessionStore).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status: expected %d, got %d", http.StatusOK, w.Code)
	}
}

func TestIfUserIsAlreadyLoggedOut(t *testing.T) {
	r := httptest.NewRequest("POST", "/logout", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(nil, false)

	api.LogoutHandler(mockSessionStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Logout Failed","Resource":"user","Field":"user_status","Code":"already_logged_out"}`)

	if w.Code != http.StatusBadRequest {
		t.Errorf("wrong status: expected %d, got %d", http.StatusBadRequest, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}

func TestIfThereIsAnErrorWhileDeletingTheSession(t *testing.T) {
	r := httptest.NewRequest("POST", "/logout", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(new(domain.Session), true)
	mockSessionStore.EXPECT().DeleteSession(gomock.Any()).Return(errors.New("persistance error while deleting the session"))

	api.LogoutHandler(mockSessionStore).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Logout Failed","Resource":"error","Field":"unexpected_error","Code":"persistance error while deleting the session"}`)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status: expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}
