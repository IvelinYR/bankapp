let axios = require('axios');

export function axiosLoad() {
    const accountList = [];
    const transactionList = [];
    const users = {username: "ivo", password: "123"};
    let MockAdapter = require('axios-mock-adapter');
    let mock = new MockAdapter(axios);

    mock.onGet('/login').reply(function (config) {
        let obj = config.data;
        let profil = JSON.parse(obj);
        let name = profil.username
        if (JSON.stringify(profil) === JSON.stringify(users)) {
            return [200, name]
        } else {
            const error = "Username or Password was wrong!"
            return [401, error]
        }
    });

    mock.onGet('/accounts').reply(200,
        accountList
    );

    mock.onPost('/accounts').reply(function (config) {
        let obj = config.data;
        let newAccount = JSON.parse(obj)
        newAccount.id = accountList.length + 1;
        newAccount.title = "ACCOUNT-" + (accountList.length + 1);
        accountList.push(newAccount);
        return [201, newAccount]
    });

    mock.onPost('/history').reply(function (config) {
        let obj = config.data;
        let newTransaction = JSON.parse(obj);
        let idNumber = Number(newTransaction.id);
        let result = transactionList.filter(function (obj) {
            return obj.id === idNumber;
        });
        return [201, result]
    });

    mock.onPost('/transaction').reply(function (config) {
        let obj = config.data;
        let currentId = JSON.parse(obj);
        let id = currentId.id - 1
        let account = accountList[id]
        return [201, account]
    });

    mock.onPost('/transaction/deposits').reply(function (config) {
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
        transactionList.push(newTransaction);
        let accountId = newAccount.id - 1;
        let amount = accountList[accountId].amount;
        newTransaction.date = today;
        let accountDeposit = newAccount.operation;
        let result = Number(amount) + Number(accountDeposit);
        accountList[accountId].amount = result;
        return [201, accountList[accountId]];
    });

    mock.onPost('/transaction/withdrawals').reply(function (config) {
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
        transactionList.push(newTransaction);
        let accountId = newAccount.id - 1;
        let amount = accountList[accountId].amount;
        newTransaction.date = today;
        let accountDeposit = newAccount.operation;
        let result = Number(amount) - Number(accountDeposit);
        accountList[accountId].amount = result;
        return [201, accountList[accountId]];
    });
}