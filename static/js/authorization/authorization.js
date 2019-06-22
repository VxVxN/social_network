$(function () {
   $.getScript("/static/js/language.js");

   var changeLangOnPage = function (language) {
      document.title = lang(language, "TITLE_AUTHORIZATION")
      document.getElementById('main').innerHTML = '<h1>' + lang(language, "MAIN_AUTHORIZATION") + '</h1>'
      document.getElementById('main_btn').value = lang(language, "MAIN_BTN_AUTHORIZATION")
      document.getElementById('email_lbl').textContent = lang(language, "EMAIL_LBL_AUTHORIZATION")
      document.getElementById('password_lbl').textContent = lang(language, "PASSWORD_LBL_AUTHORIZATION")
      document.getElementById('log_in_btn').value = lang(language, "LOG_IN_BTN_AUTHORIZATION")
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
   var error = { err: document.getElementById("error") };
   if (email == "" || password == "") {
      error.err.innerHTML = "The fields must be filled in.";
      return
   }
   error.err.innerHTML = "";
   $('#form').trigger('submit');
});