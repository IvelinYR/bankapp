let axios = require('axios');

export function axiosLoad() {
    const accountList = [];
    const transactionList = [];
    const users = [{Username: "ivo", Password: "123"}];
    let MockAdapter = require('axios-mock-adapter');
    let mock = new MockAdapter(axios);

    mock.onGet('/v1/users/login').reply(function (config) {
        let obj = config.data;
        let profil = JSON.parse(obj);
        let SID = 'XXSDFSSAAA23423SDDFFSASD3434ASAD';
        if (JSON.stringify([profil]) === JSON.stringify(users)) {
            return [200, [SID]]
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

    mock.onGet('/v1/users/me/accounts').reply(200,
        accountList
    );

    mock.onPost('/v1/users/me/new-account').reply(function (config) {
        console.log(config);
        let obj = config.data;
        let newAccount = JSON.parse(obj)
        newAccount.id = accountList.length + 1 + "dsddssdd";
        newAccount.title = "ACCOUNT-" + (accountList.length + 1);
        accountList.push(newAccount);
        return [201, newAccount]
    });

    mock.onPost('/v1/users/me/account-history').reply(function (config) {
        let obj = config.data;
        let newTransaction = JSON.parse(obj);
        let idNumber = newTransaction.id;
        let result = transactionList.filter(function (obj) {
            return obj.id === idNumber;
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
        newTransaction.date = today;
        transactionList.push(newTransaction);
        let result = accountList.filter(function (obj) {
            return obj.id === newTransaction.id;
        });
        let account = result[0];
        let amount = account.Total;
        let accountDeposit = newAccount.Amount;
        let sum = Number(amount) + Number(accountDeposit);
        account.Total = sum;
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
        newTransaction.date = today;
        transactionList.push(newTransaction);
        let result = accountList.filter(function (obj) {
            return obj.id === newTransaction.id;
        });
        let account = result[0];
        let amount = account.Total;
        console.log(amount)
        let accountDeposit = newAccount.Amount;
        let sum = Number(amount) - Number(accountDeposit);
        account.Total = sum;
        return [201, account];
    });
}