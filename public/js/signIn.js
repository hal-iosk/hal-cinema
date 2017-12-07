$(document).ready(function () {
    // 入力チェックの初期化
    $.el_valid.init({
      "password": {
        "email": {
            "rule": [
              "required",
              "email"
            ]
          },
        "rule": [
          "required",
          "min_length[8]",
          "max_length[12]",
          "alpha_numeric"
        ]
      }
    });
  });
  
  // function errorcheck() {
  //   password = document.getElementsById("password");
  //   confirmPass = document.getElementsById("confirmPass");
  //   if (password.value != confirmPass.value) {
  //     $("#e_confirmPass").text("パスワードが一致しません。")
  //     return false;
  //   }
  // }
  
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

  
    console.log(Object(validationMsg["email"]));
    console.log(Object(validationMsg["email"].length));
    console.log(Object(e_email));
    if (!validationMsg["email"].length == 0 || !validationMsg["password"].length == 0) {
      document.getElementById("e_email").innerHTML = validationMsg["email"];
      document.getElementById("e_password").innerHTML = validationMsg["password"];
      return false;
    } else {
      return submit();
    }
  });
  