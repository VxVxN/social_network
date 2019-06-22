$(function () {
    $.getScript("/static/js/language.js");

    var changeLangOnPage = function (language) {
        document.title = lang(language, "TITLE_REGISTRATION")
        document.getElementById('main').innerHTML = '<h1>' + lang(language, "MAIN_REGISTRATION") + '</h1>'
        document.getElementById('main_btn').value = lang(language, "MAIN_BTN_REGISTRATION")
        document.getElementById('username_lbl').textContent = lang(language, "USERNAME_LBL_REGISTRATION")
        document.getElementById('fname_lbl').textContent = lang(language, "FNAME_LBL_REGISTRATION")
        document.getElementById('lname_lbl').textContent = lang(language, "LNAME_LBL_REGISTRATION")
        document.getElementById('email_lbl').textContent = lang(language, "EMAIL_LBL_REGISTRATION")
        document.getElementById('password_lbl').textContent = lang(language, "PASSWORD_LBL_REGISTRATION")
        document.getElementById('sign_up_btn').value = lang(language, "SIGN_UP_BTN_REGISTRATION")
    };

    $.ajax({
        type: "GET",
        url: "/ajax/language",
    }).done(function (data) {
        data = JSON.parse(data);
        changeLangOnPage(data.Language);
    });
});

$('#buttonCheck').click(function () {
    var password = (document.getElementById("password")).value;
    var email = (document.getElementById("email")).value;
    var username = (document.getElementById("username")).value;
    var error = { err: document.getElementById("error") };

    if (username == "") {
        error.err.innerHTML = "Username is a required field.";
        return
    }
    if (email == "") {
        error.err.innerHTML = "Email is a required field.";
        return
    }
    if (password == "") {
        error.err.innerHTML = "Password is a required field.";
        return
    }
    if (!validateEmail(email)) {
        error.err.innerHTML = "Email is invalid.";
        return
    }
    error.err.innerHTML = "";
    $('#form').trigger('submit');
});

var validateEmail = function (email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
};