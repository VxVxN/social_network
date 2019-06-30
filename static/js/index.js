$(function () {
    $.getScript("/static/js/language.js");

    var changeLangOnPage = function (language) {
        document.title = lang(language, "TITLE_INDEX")
        $('#main').html('<h1>' + lang(language, "MAIN_INDEX") + '</h1>')
        $('#log_in_btn').val(lang(language, "LOG_IN_INDEX"))
        $('#sign_up_btn').val(lang(language, "SIGN_UP_INDEX"))
        $('#select_lang').val(language)
    };

    $.ajax({
        type: "GET",
        url: "/ajax/language",
    }).done(function (data) {
        data = JSON.parse(data);
        changeLangOnPage(data.Language);
    });

    $("#select_lang").change(function () {
        var data
        if ($(this).val() == "EN") {
            data = {
                language: "EN"
            }
        }
        if ($(this).val() == "RU") {
            data = {
                language: "RU"
            }
        }
        $.ajax({
            type: "POST",
            url: "/ajax/language",
            dataType: "json",
            data: JSON.stringify(data)
        }).done(function () {
            changeLangOnPage(data);
            document.location.reload(true);
        });
    });
});