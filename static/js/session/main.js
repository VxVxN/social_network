$(function () {
    $.getScript("/static/js/online.js");
    $.getScript("/static/js/language.js");

    $("#send_message").hide();
    $("#message_input").hide();

    var changeLangOnPage = function (language) {
        // document.title = lang(language, "TITLE_AUTHORIZATION")
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
        requestListUsers()
    });
});

var requestListUsers = function () {// TODO упростить функцию и сделать её нормальной
    $.ajax({
        type: "GET",
        url: "/ajax/list_users",
    }).done(function (data) {
        data = JSON.parse(data);
        var userOnline = $('#user_online');
        userOnline.val("");
        $(".item").remove();
        if (data.Nickname != null) {
            data.Nickname.forEach(function (nickname) {
                $("#user_online").append("<li class='item'><button class='button'>" + nickname + "</button></li>");
            });
            $(".button").click(function (data) {
                nicknameInterlocutor = data.srcElement.innerText
                getMessages()
            });
        }
    });
};

$("#send_message").click(function () {
    if (nicknameInterlocutor != "") {
        data = {
            nickname: nicknameInterlocutor,
            message: $("#message_input").val()
        }
        $.ajax({
            type: "POST",
            url: "/ajax/send_message",
            data: JSON.stringify(data)
        }).done(function () {
        });
    }
    $("#message_input").val("")
    getMessages()
});

var getMessages = function () {
    if (nicknameInterlocutor != undefined) {
        $.ajax({
            type: "GET",
            url: "/ajax/get_messages",
            data: "nickname=" + nicknameInterlocutor,
        }).done(function (dataresp) {
            dataresp = JSON.parse(dataresp);
            $("#send_message").show()
            $("#message_input").show()
            $(".item_msg_nickname").remove();
            $(".item_message").remove();
            if (dataresp != null) {
                dataresp.forEach(function (data) {
                    $("#messages").append("<dt class='item_msg_nickname'>" + data.Nickname + "</dt>");
                    $("#messages").append("<dd class='item_message'>" + data.Message + "</dd>");
                });
            }
        });
    };
};

var interval = 1000 * 5; // request once per 5 seconds
var nicknameInterlocutor

setInterval(requestListUsers, interval);
setInterval(getMessages, interval);