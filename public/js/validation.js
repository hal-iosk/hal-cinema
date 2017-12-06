$(document).ready(function () {
  // 入力チェックの初期化
  $.el_valid.init({
    "password": {
      "rule": [
        "required",
        "min_length[8]",
        "max_length[12]",
        "alpha_numeric"
      ]
    },
    "first_name": {
      "rule": [
        "required",
        "zenkaku"
      ]
    },
    "last_name": {
      "rule": [
        "required",
        "zenkaku"
      ]
    },
    "first_name_read": {
      "rule": [
        "required",
        "zenkaku"
      ]
    },
    "last_name_read": {
      "rule": [
        "required",
        "zenkaku"
      ]
    },
    "phone": {
      "rule": [
        "required",
        "tel",
      ]
    },
    "email": {
      "rule": [
        "required",
        "email"
      ]
    },
    "address": {
      "rule": [
        "required",
        "min_length[7]",
        'numeric'
      ]
    },
    "credit_card_number": {
      "rule": [
        'numeric',
        "required",
        "min_length[15]",
        "max_length[16]"
      ]
    },
    "security_code": {
      "rule": [
        'numeric',
        "required",
        "min_length[3]",
      ]
    },
    "nominee": {
      "rule": [
        'alpha',
        "required",
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
      email:'',
      password:'',
      confirmPass:'',
      first_name:'',
      last_name:'',
      first_name_read:'',
      last_name_read:'',
      birth:'',
      phone:'',
      address:'',
      credit_card_number:'',
      security_code:''
  };

  var email = $('#email').val();
  var e_email = $('#e_email').val();

  var password = $('#password').val();
  var confirmPass = $('#confirmPass').val();
  var e_password = $('#e_password').val();
  var e_confirmPass = $('#e_confirmPass').val();

  var first_name = $('#first_name').val();
  var e_first_name = $('#e_first_name').val();

  var last_name = $('#last_name').val();
  var e_last_name = $('#e_last_name').val();
  
  
  if (email.length == 0) {
    validationMsg["email"] = "メールアドレスを入力してください。";
  } 
   else if(!e_email.length == 0) {
    validationMsg["email"] = $('#e_email');
  }

  if (password.length == 0) {
    validationMsg["password"] = "パスワードを入力してください。";
  } 
  else if (password.length < 8) {
    validationMsg["password"] ="パスワードは8文字以上で入力して下さい";
  } 
  if (password != confirmPass) {
    validationMsg["confirmPass"] = "パスワードが一致しませんでした。";
  } 
  if(!e_password.length == 0){
      validationMsg["password"] = e_password;
  }
  if (first_name.length == 0) {
    validationMsg["first_name"] = "お名前(姓)を入力してください";
  } 
  else if (!e_first_name.length == 0) {
    validationMsg["first_name"] = e_first_name;
  }


// $("#e_email").text(validationMsg["email"]);
// $("#e_password").text(validationMsg["password"]);
// $("#e_confirmPass").text(validationMsg["confirmPass"]);
// return false;
  });
