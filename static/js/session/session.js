$(function () {
    $.getScript("/static/js/online.js");
    $.getScript("/static/js/language.js");

    var changeLangOnPage = function (language) {
        document.title = lang(language, "TITLE_AUTHORIZATION")
        // document.getElementById('main').innerHTML = '<h1>' + lang(language, "MAIN_AUTHORIZATION") + '</h1>'
        // document.getElementById('main_btn').value = lang(language, "MAIN_BTN_AUTHORIZATION")
        // document.getElementById('email_lbl').textContent = lang(language, "EMAIL_LBL_AUTHORIZATION")
        // document.getElementById('password_lbl').textContent = lang(language, "PASSWORD_LBL_AUTHORIZATION")
        // document.getElementById('log_in_btn').value = lang(language, "LOG_IN_BTN_AUTHORIZATION")
    };

    $.ajax({
        type: "GET",
        url: "/ajax/language",
    }).done(function (data) {
        data = JSON.parse(data);
        changeLangOnPage(data.Language);
    });
});

var requestListUsers = function () {
    $.ajax({
        type: "GET",
        url: "/ajax/list_users",
    }).done(function (data) {
        data = JSON.parse(data);
        var userOnline = $('#user_online')
        userOnline.val("");
        data.Nickname.forEach(function (nickname) {
            userOnline.val(userOnline.val() + nickname + "\n")
        });
    });
};

var interval = 1000 * 5; // request once per 30 seconds

setInterval(requestListUsers, interval);