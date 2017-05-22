let axios = require('axios');

export function axiosLoad() {
    const accountList = [];
    const transactionList = [];
    const users = [{Username: "ivo", Password: "123"}];
    let MockAdapter = require('axios-mock-adapter');
    let mock = new MockAdapter(axios);

    mock.onPost('/v1/users/login').reply(function (config) {
        let obj = config.data;
        let profil = JSON.parse(obj);
        let SID = {SessionID : 'XXSDFSSAAA23423SDDFFSASD3434ASAD'};
        if (JSON.stringify([profil]) === JSON.stringify(users)) {
            return [200, SID]
        } else {
            return [400, 'ERROR']
        }
    });

    mock.onPost('/v1/users/signup').reply(function (config) {
        let obj = config.data;
        let NewUser = JSON.parse(obj);
        users.push(NewUser);
        if(true) {
            return [200, name]
        } else {
            return [400, 'username does not exist']
        }
    });

    mock.onPost('/v1/users/logout').reply(function () {
      return [201]
    });

    mock.onGet('/v1/users/me/accounts').reply(200,
        accountList
    );

    mock.onPost('/v1/users/me/new-account').reply(function (config) {
        let obj = config.data;
        let newAccount = JSON.parse(obj)
        newAccount.AccountID = accountList.length + 1 + "dsddssdd";
        newAccount.title = "ACCOUNT-" + (accountList.length + 1);
        accountList.push(newAccount);
        let error = {Message:"Create User Account Failed"};
        return [201, error]
    });

    mock.onPost('/v1/users/me/account-history').reply(function (config) {
        let obj = config.data;
        let newTransaction = JSON.parse(obj);
        let idNumber = newTransaction.AccountID;
        let result = transactionList.filter(function (obj) {
            return obj.AccountID === idNumber;
        });
        return [201, result]
    });

    mock.onPost('/v1/users/me/deposit').reply(function (config) {
        let obj = config.data;
        let newAccount = JSON.parse(obj);
        let newTransaction = JSON.parse(obj);
        let today = new Date();
        let dd = today.getDate();
        let mm = today.getMonth() + 1;
        let yyyy = today.getFullYear();
        let hour = today.getHours();
        let minutes = today.getMinutes();
        let seconds = today.getSeconds();
        if (dd < 10) {dd = '0' + dd}
        if (mm < 10) {mm = '0' + mm}
        if (hour < 10) {hour = '0' + hour}
        if (minutes < 10) {minutes = '0' + minutes}
        if (seconds < 10) {seconds = '0' + seconds}
        today = mm + '/' + dd + '/' + yyyy + ' ' + hour + ':' + minutes + ':' + seconds;
        newTransaction.Date = today;
        transactionList.push(newTransaction);
        let result = accountList.filter(function (obj) {
            return obj.AccountID === newTransaction.AccountID;
        });
        let account = result[0];
        let amount = account.Amount;
        let accountDeposit = newAccount.Amount;
        console.log(accountDeposit)
        let sum = Number(amount) + Number(accountDeposit);
        account.Amount = sum;
        return [201, account];
    });

    mock.onPost('/v1/users/me/withdraw').reply(function (config) {
        let obj = config.data;
        let newAccount = JSON.parse(obj);
        let newTransaction = JSON.parse(obj);
        let today = new Date();
        let dd = today.getDate();
        let mm = today.getMonth() + 1;
        let yyyy = today.getFullYear();
        let hour = today.getHours();
        let minutes = today.getMinutes();
        let seconds = today.getSeconds();
        if (dd < 10) {dd = '0' + dd}
        if (mm < 10) {mm = '0' + mm}
        if (hour < 10) {hour = '0' + hour}
        if (minutes < 10) {minutes = '0' + minutes}
        if (seconds < 10) {seconds = '0' + seconds}
        today = mm + '/' + dd + '/' + yyyy + ' ' + hour + ':' + minutes + ':' + seconds;
        newTransaction.Date = today;
        transactionList.push(newTransaction);
        let result = accountList.filter(function (obj) {
            return obj.AccountID === newTransaction.AccountID;
        });
        let account = result[0];
        let amount = account.Amount;
        let accountDeposit = newAccount.Amount;
        let sum = Number(amount) - Number(accountDeposit);
        account.Amount = sum;
        return [201, account];
    });
}