var langEN = {};
var langRU = {};

var lang = function (language, text) {
    if (language == "EN") {
        return langEN[text].toString()
    }
    if (language == "RU") {
        return langRU[text].toString()
    }
    return langEN[text].toString()
};

langEN["TITLE_INDEX"] = "Main"
langRU["TITLE_INDEX"] = "Главная"
langEN["MAIN_INDEX"] = "Main"
langRU["MAIN_INDEX"] = "Главная страница"
langEN["LOG_IN_INDEX"] = "Log in"
langRU["LOG_IN_INDEX"] = "Войти"
langEN["SIGN_UP_INDEX"] = "Sign up"
langRU["SIGN_UP_INDEX"] = "Зарегестрироваться"

langEN["TITLE_AUTHORIZATION"] = "Log in"
langRU["TITLE_AUTHORIZATION"] = "Авторизация"
langEN["MAIN_AUTHORIZATION"] = "Log in"
langRU["MAIN_AUTHORIZATION"] = "Авторизация"
langEN["MAIN_BTN_AUTHORIZATION"] = "Main"
langRU["MAIN_BTN_AUTHORIZATION"] = "Главная"
langEN["EMAIL_LBL_AUTHORIZATION"] = "Email: "
langRU["EMAIL_LBL_AUTHORIZATION"] = "Емайл: "
langEN["PASSWORD_LBL_AUTHORIZATION"] = "Password: "
langRU["PASSWORD_LBL_AUTHORIZATION"] = "Пароль: "
langEN["LOG_IN_BTN_AUTHORIZATION"] = "Log in"
langRU["LOG_IN_BTN_AUTHORIZATION"] = "Войти"


langEN["TITLE_REGISTRATION"] = "Sign up"
langRU["TITLE_REGISTRATION"] = "Регистрация"
langEN["MAIN_REGISTRATION"] = "Sign up"
langRU["MAIN_REGISTRATION"] = "Регистрация"
langEN["MAIN_BTN_REGISTRATION"] = "Main"
langRU["MAIN_BTN_REGISTRATION"] = "Главная страница"
langEN["USERNAME_LBL_REGISTRATION"] = "Username*: "
langRU["USERNAME_LBL_REGISTRATION"] = "Ник*: "
langEN["FNAME_LBL_REGISTRATION"] = "First name: "
langRU["FNAME_LBL_REGISTRATION"] = "Имя: "
langEN["LNAME_LBL_REGISTRATION"] = "Last name: "
langRU["LNAME_LBL_REGISTRATION"] = "Фамилия: "
langEN["EMAIL_LBL_REGISTRATION"] = "Email*: "
langRU["EMAIL_LBL_REGISTRATION"] = "Емайл*: "
langEN["PASSWORD_LBL_REGISTRATION"] = "Password*: "
langRU["PASSWORD_LBL_REGISTRATION"] = "Пароль*: "
langEN["SIGN_UP_BTN_REGISTRATION"] = "Sign up"
langRU["SIGN_UP_BTN_REGISTRATION"] = "Зарегистрироваться"