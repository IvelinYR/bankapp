package api

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/iliyanmotovski/bankv1/bank/domain"
	"net/http"
)

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
