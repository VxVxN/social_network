$(function () {
   $.getScript("/static/js/language.js");

   var changeLangOnPage = function (language) {
      document.title = lang(language, "TITLE_AUTHORIZATION")
      $('#main').html('<h1>' + lang(language, "MAIN_AUTHORIZATION") + '</h1>')
      $('#main_btn').val(lang(language, "MAIN_BTN_AUTHORIZATION"))
      $('#email_lbl').text(lang(language, "EMAIL_LBL_AUTHORIZATION"))
      $('#password_lbl').text(lang(language, "PASSWORD_LBL_AUTHORIZATION"))
      $('#log_in_btn').val(lang(language, "LOG_IN_BTN_AUTHORIZATION"))
   };

   $.ajax({
      type: "GET",
      url: "/ajax/language",
   }).done(function (data) {
      data = JSON.parse(data);
      changeLangOnPage(data.Language);
   });
});

$('#log_in_btn').click(function () {
   var password = (document.getElementById("password")).value;
   var email = (document.getElementById("email")).value;
   var error = { err: document.getElementById("error") };
   if (email == "" || password == "") {
      error.err.innerHTML = "The fields must be filled in.";
      return
   }
   error.err.innerHTML = "";
   $('#form').trigger('submit');
});

$(document).on('keypress', function (e) {
   if (e.which == 13) {
      $('#log_in_btn').click();
   }
});