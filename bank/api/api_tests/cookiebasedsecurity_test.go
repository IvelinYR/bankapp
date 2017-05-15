package api_tests

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/context"
	"github.com/iliyanmotovski/bankv3/bank/api"
	"github.com/iliyanmotovski/bankv3/bank/domain"
	"github.com/iliyanmotovski/bankv3/bank/domain/mock_domain"
	"github.com/justinas/alice"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestSecurityPassesAndSessionIsUpdated(t *testing.T) {
	testHandler := func() http.HandlerFunc {
		fn := func(w http.ResponseWriter, r *http.Request) {
			session := context.Get(r, "session").(*domain.Session)
			if session.SessionID != "12345" {
				t.Error("expected SessionID to be '12345', but it wasn't")
			}
		}
		return http.HandlerFunc(fn)
	}

	account := []domain.Account{}
	account = append(account, domain.Account{UserID: "12345", AccountID: "54321", Currency: "BGN", Amount: 100})

	r := httptest.NewRequest("POST", "/deposit", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(&domain.Session{SessionID: "12345"}, true)
	mockSessionStore.EXPECT().UpdateSession(gomock.Any(), gomock.Any()).Return(nil)

	security := alice.New(api.CookieBasedSecurity(mockSessionStore, time.Second))

	security.Then(testHandler()).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("wrong status: expected %d, got %d", http.StatusOK, w.Code)
	}
}

func TestSessionWasNotValid(t *testing.T) {
	testHandler := func() http.HandlerFunc {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t.Error("should not have entered testHandler")
		}
		return http.HandlerFunc(fn)
	}

	r := httptest.NewRequest("POST", "/deposit", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(&domain.Session{}, false)

	security := alice.New(api.CookieBasedSecurity(mockSessionStore, time.Second))
	security.Then(testHandler()).ServeHTTP(w, r)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("wrong status: expected %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestUpdateSessionFailed(t *testing.T) {
	testHandler := func() http.HandlerFunc {
		fn := func(w http.ResponseWriter, r *http.Request) {
		}
		return http.HandlerFunc(fn)
	}

	account := []domain.Account{}
	account = append(account, domain.Account{UserID: "12345", AccountID: "54321", Currency: "BGN", Amount: 100})

	r := httptest.NewRequest("POST", "/deposit", nil)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSessionStore := mock_domain.NewMockSessionStore(ctrl)

	mockSessionStore.EXPECT().FindSessionAvailableAt(gomock.Any(), gomock.Any()).Return(&domain.Session{}, true)
	mockSessionStore.EXPECT().UpdateSession(gomock.Any(), gomock.Any()).Return(errors.New("persistence error while updating session"))

	security := alice.New(api.CookieBasedSecurity(mockSessionStore, time.Second))

	security.Then(testHandler()).ServeHTTP(w, r)

	expected := []byte(`{"Message":"Session Update Failed","Resource":"error","Field":"unexpected_error","Code":"persistence error while updating session"}`)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("wrong status: expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
	if !reflect.DeepEqual(w.Body.Bytes()[:len(w.Body.Bytes())-1], expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), string(expected))
	}
}
