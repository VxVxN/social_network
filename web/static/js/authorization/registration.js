$(function () {
    $.getScript("/static/js/language.js");

    var changeLangOnPage = function (language) {
        document.title = lang(language, "TITLE_REGISTRATION")
        $('#main').html('<h1>' + lang(language, "MAIN_REGISTRATION") + '</h1>')
        $('#main_btn').val(lang(language, "MAIN_BTN_REGISTRATION"))
        $('#username_lbl').text(lang(language, "USERNAME_LBL_REGISTRATION"))
        $('#fname_lbl').text(lang(language, "FNAME_LBL_REGISTRATION"))
        $('#lname_lbl').text(lang(language, "LNAME_LBL_REGISTRATION"))
        $('#email_lbl').text(lang(language, "EMAIL_LBL_REGISTRATION"))
        $('#password_lbl').text(lang(language, "PASSWORD_LBL_REGISTRATION"))
        $('#sign_up_btn').val(lang(language, "SIGN_UP_BTN_REGISTRATION"))
    };

    $.ajax({
        type: "GET",
        url: "/ajax/language",
    }).done(function (data) {
        data = JSON.parse(data);
        changeLangOnPage(data.data);
    });
});

$('#sign_up_btn').click(function () {
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

$(document).on('keypress', function (e) {
    if (e.which == 13) {
        $('#sign_up_btn').click();
    }
});