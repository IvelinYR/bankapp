package api

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/iliyanmotovski/bankv3/bank/domain"
	"net/http"
	"time"
)

func SignUpHandler(userStore domain.UserStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request domain.UserRegistrationRequest
		json.NewDecoder(r.Body).Decode(&request)
		w.Header().Set("Content-Type", "application/json")

		u, err := userStore.RegisterUser(request)
		if err == nil {
			responseUser := &domain.ResponseUser{UserID: u.UserID, Username: u.Username, Email: u.Email, Name: u.Name, Age: u.Age, DateCreated: time.Now().Local()}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(responseUser)
			return
		}

		switch err {
		case domain.ErrUserAlreadyExists:
			errorResponse(w, http.StatusBadRequest, "SignUp Failed", "user", "username", "already_exist")
		case domain.ErrEmailAlreadyExists:
			errorResponse(w, http.StatusBadRequest, "SignUp Failed", "user", "email", "already_exist")
		default:
			errorResponse(w, http.StatusInternalServerError, "SignUp Failed", "error", "unexpected_error", err.Error())
		}
	})
}

func LoginHandler(userStore domain.UserStore, sessionStore domain.SessionStore, userSessionDuration time.Duration) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("SID")
		if err != nil {
			cookie = &http.Cookie{Value: "rfBd56ti2SMtY"}
		}

		w.Header().Set("Content-Type", "application/json")

		_, ok := sessionStore.FindSessionAvailableAt(cookie.Value, time.Now().Local())
		if ok {
			errorResponse(w, http.StatusBadRequest, "Login Failed", "user", "user_status", "already_logged_in")
			return
		}

		var request domain.UserLoginRequest
		json.NewDecoder(r.Body).Decode(&request)

		user, err := userStore.Authenticate(request)
		if err != nil {
			switch err {
			case domain.ErrUsernameDoesntExist:
				errorResponse(w, http.StatusNotFound, "Login Failed", "user", "username", "username_does_not_exist")
				return
			case domain.ErrAlreadyLoggedIn:
				errorResponse(w, http.StatusBadRequest, "Login Failed", "user", "user_status", "already_logged_in")
				return
			case domain.ErrWrongPassword:
				errorResponse(w, http.StatusUnauthorized, "Login Failed", "user", "password", "wrong_password")
				return
			default:
				errorResponse(w, http.StatusInternalServerError, "Login Failed", "error", "unexpected_error", err.Error())
				return
			}
		}

		session, err := sessionStore.StartSession(*user, time.Now().Local().Add(userSessionDuration))
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, "Session Initializing Failed", "error", "unexpected_error", err.Error())
			return
		}
		//json.NewEncoder(w).Encode(session)
		*cookie = http.Cookie{Name: "SID", Path: "/", Value: session.SessionID, MaxAge: 300}
		http.SetCookie(w, cookie)
	})
}

func LogoutHandler(sessionStore domain.SessionStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("SID")
		if err != nil {
			cookie = &http.Cookie{Value: "rfBd56ti2SMtY"}
		}

		w.Header().Set("Content-Type", "application/json")

		session, ok := sessionStore.FindSessionAvailableAt(cookie.Value, time.Now().Local())
		if !ok {
			errorResponse(w, http.StatusBadRequest, "Logout Failed", "user", "user_status", "already_logged_out")
			return
		}

		err = sessionStore.DeleteSession(session.SessionID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, "Logout Failed", "error", "unexpected_error", err.Error())
			return
		}

		//json.NewEncoder(w).Encode(session)
		*cookie = http.Cookie{Name: "SID", Path: "/", Value: "deleted", MaxAge: -1}
		http.SetCookie(w, cookie)

		w.WriteHeader(http.StatusOK)
	})
}

func GetUserAccounts(accountStore domain.AccountStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := context.Get(r, "session").(*domain.Session)

		result, err := accountStore.GetAccounts(session.UserID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, "Fetch User Accounts Failed", "error", "unexpected_error", err.Error())
			return
		}

		json.NewEncoder(w).Encode(result)
	})
}

func NewUserAccount(accountStore domain.AccountStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := context.Get(r, "session").(*domain.Session)

		var account domain.Account
		json.NewDecoder(r.Body).Decode(&account)

		_, err := accountStore.InsertAccount(session.UserID, account)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, "Create User Account Failed", "error", "unexpected_error", err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
	})
}

func DeleteUserAccount(accountStore domain.AccountStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := context.Get(r, "session").(*domain.Session)

		var account domain.Account

		value := r.URL.Query()
		accountID := value["id"]

		if len(accountID) == 0 {
			errorResponse(w, http.StatusBadRequest, "Delete User Account Failed", "request", "URL_parameters", "need_accountID_to_be_specified_in_URL")
			return
		}
		account.UserID = session.UserID
		account.AccountID = accountID[0]

		err := accountStore.DeleteAccount(account.UserID, account.AccountID)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, "Delete User Accounts Failed", "error", "unexpected_error", err.Error())
			return
		}
	})
}

func UserAccountDeposit(accountStore domain.AccountStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := context.Get(r, "session").(*domain.Session)

		var account domain.Account
		json.NewDecoder(r.Body).Decode(&account)
		account.UserID = session.UserID

		err := accountStore.Deposit(account)
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, "User Account Deposit Failed", "error", "unexpected_error", err.Error())
			return
		}
	})
}

func UserAccountWithdraw(accountStore domain.AccountStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := context.Get(r, "session").(*domain.Session)

		var account domain.Account
		json.NewDecoder(r.Body).Decode(&account)
		account.UserID = session.UserID

		err := accountStore.Withdraw(account)
		if err != nil {
			switch err {
			case domain.ErrWithdrawMoreThanHave:
				errorResponse(w, http.StatusBadRequest, "User Account Withdraw Failed", "account", "money_availability", "cannot_withdraw_more_than_have_available")
				return
			default:
				errorResponse(w, http.StatusInternalServerError, "User Account Withdraw Failed", "error", "unexpected_error", err.Error())
				return
			}
		}
	})
}

func UserTransactionHistory(accountStore domain.AccountStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := context.Get(r, "session").(*domain.Session)

		var historyRequest domain.History
		json.NewDecoder(r.Body).Decode(&historyRequest)

		historyRequest.UserID = session.UserID

		result, err := accountStore.GetHistory(historyRequest)
		if err != nil {
			switch err {
			case domain.ErrUnauthorized:
				errorResponse(w, http.StatusUnauthorized, "Fetch User Account History Failed", "account", "account_authorization", "unauthorized_to_see_this_account_history")
				return
			default:
				errorResponse(w, http.StatusInternalServerError, "Fetch User Account History Failed", "error", "unexpected_error", err.Error())
				return
			}
		}

		json.NewEncoder(w).Encode(result)
	})
}

func errorResponse(w http.ResponseWriter, HTTPstatus int, message string, resource string, field string, code string) {
	w.WriteHeader(HTTPstatus)
	resp := domain.NewErrorResponse(message, resource, field, code)
	json.NewEncoder(w).Encode(resp)
}
