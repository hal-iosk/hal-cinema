$(document).ready(function () {
    // 入力チェックの初期化
    $.el_valid.init({

        "email": {
            "rule": [
                "required",
                "email"
            ]
        },
        "password": {
            "rule": [
                "required",
                "min_length[]",
                "max_length[12]",
                "alpha_numeric"
            ]
        }
    });
});


//バリデーション
$('form').submit(function () {
    var validationMsg = {
        email: '',
        password: ''
    };

    var email = $('#email').val();
    var e_email = document.getElementById('e_email').innerHTML;
    var password = $('#password').val();
    var e_password = document.getElementById('e_password').innerHTML;


    if (email.length == 0) {
        validationMsg["email"] = "メールアドレスを入力してください。";
    } else if (!e_email.length == 0) {
        validationMsg["email"] = e_email;
    } else {
        validationMsg["email"] = "";
    }

    if (password.length == 0) {
        validationMsg["password"] = "パスワードを入力してください。";
    } else if (!e_password.length == 0) {
        validationMsg["password"] = e_password;
    } else {
        validationMsg["password"] = "";
    }
    if (!validationMsg["email"].length == 0 || !validationMsg["password"].length == 0) {
        document.getElementById("e_email").innerHTML = validationMsg["email"];
        document.getElementById("e_password").innerHTML = validationMsg["password"];
        return false;
    } else {
        return submit();
    }
});