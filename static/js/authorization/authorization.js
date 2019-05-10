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