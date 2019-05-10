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