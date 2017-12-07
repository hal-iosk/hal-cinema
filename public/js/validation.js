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
function validation() {
  var validationMsg = {
    email: '',
    password: '',
    confirmPass: '',
    first_name: '',
    last_name: '',
    first_name_read: '',
    last_name_read: '',
    birth: '',
    phone: '',
    address: '',
    credit_card_number: '',
    security_code: '',
    limit: ''
  };

  var email = $('#email').val();
  var e_email = document.getElementById('e_email').innerHTML;
  var password = $('#password').val();
  var e_password = document.getElementById('e_password').innerHTML;
  var confirmPass = $('#confirmPass').val();
  var e_confirmPass = document.getElementById('e_confirmPass').innerHTML;
  var first_name = $('#first_name').val();
  var e_first_name = document.getElementById('e_first_name').innerHTML;
  var last_name = $('#last_name').val();
  var e_last_name = document.getElementById('e_last_name').innerHTML;
  var first_name_read = $('#first_name_read').val();
  var e_first_name_read = document.getElementById('e_first_name_read').innerHTML;
  var last_name_read = $('#last_name_read').val();
  var e_last_name_read = document.getElementById('e_last_name_read').innerHTML;
  var year = $('#year').val();
  var month = $('#month').val();
  var date = $('#date').val();
  var phone = $('#phone').val();
  var e_phone = document.getElementById('e_phone').innerHTML;
  var address = $('#address').val();
  var e_address = document.getElementById('e_address').innerHTML;
  var credit_card_number = $('#credit_card_number').val();
  var e_credit_card_number = document.getElementById('e_credit_card_number').innerHTML;
  var security_code = $('#security_code').val();
  var e_security_code = document.getElementById('e_security_code').innerHTML;
  var limitMonth = $('#limitMonth').val();
  //   var e_limitMonth = document.getElementById('e_limitMonth').innerHTML;
  var limitYear = $('#limitYear').val();
  //   var e_limitYear = document.getElementById('e_limitYear').innerHTML;
  validationMsg["birth"] = "";
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

  if (password != confirmPass) {
    validationMsg["confirmPass"] = "パスワードが一致しませんでした。";
  } else {
    validationMsg["confirmPass"] = "";
  }

  if (first_name.length == 0) {
    validationMsg["first_name"] = "お名前(姓)を入力してください";
  } else if (!e_first_name.length == 0) {
    validationMsg["first_name"] = e_first_name;
  } else {
    validationMsg["first_name"] = "";
  }

  if (last_name.length == 0) {
    validationMsg["last_name"] = "お名前(名)を入力してください";
  } else if (!e_last_name.length == 0) {
    validationMsg["last_name"] = e_last_name;
  } else {
    validationMsg["last_name"] = "";
  }

  if (first_name_read.length == 0) {
    validationMsg["first_name_read"] = "フリガナ(セイ)を入力してください";
  } else if (!e_first_name_read.length == 0) {
    validationMsg["first_name_read"] = e_first_name_read;
  } else {
    validationMsg["first_name_read"] = "";
  }

  if (last_name_read.length == 0) {
    validationMsg["last_name_read"] = "フリガナ(メイ)を入力してください";
  } else if (!e_last_name_read.length == 0) {
    validationMsg["last_name_read"] = e_last_name_read;
  } else {
    validationMsg["last_name_read"] = "";
  }

  if (year.length == 0 || month.length == 0 || date.length == 0) {
    validationMsg["birth"] = "生年月日を入力してください";
  }
  if (moment([year, month, date], 'YYYY-MM-DD').isValid()) {
    validationMsg["birth"] = "";
  } else {
    validationMsg["birth"] = "正しい生年月日を入力してください";
  }

  if (phone.length == 0) {
    validationMsg["phone"] = "電話番号を入力してください";
  } else if (!e_phone.length == 0) {
    validationMsg["phone"] = e_phone;
  } else {
    validationMsg["phone"] = "";
  }
  if (address.length == 0) {
    validationMsg["address"] = "郵便番号を入力してください";
  } else if (!e_address.length == 0) {
    validationMsg["address"] = e_address;
  } else {
    validationMsg["address"] = "";
  }
  if (credit_card_number.length == 0) {
    validationMsg["credit_card_number"] = "クレジットカードの番号を入力してください";
  } else if (!e_credit_card_number.length == 0) {
    validationMsg["credit_card_number'"] = e_credit_card_number;
  } else {
    validationMsg["credit_card_numbr"] = "";
  }
  if (security_code.length == 0) {
    validationMsg["security_code"] = "セキュリティコードを入力してください";
  } else if (!e_security_code.length == 0) {
    validationMsg["security_code"] = e_security_code;
  } else {
    validationMsg["security_code"] = "";
  }
  if (limitMonth.length == 0 || limitYear.length == 0) {
    validationMsg["limit"] = "有効期限を入力してください";
  } else {
    validationMsg["limit"] = "";
  }

  if (!validationMsg["email"].length == 0 || !validationMsg["password"].length == 0 || !validationMsg["confirmPass"].length == 0 || !validationMsg["first_name"].length == 0 || !validationMsg["last_name"].length == 0 || !validationMsg["first_name_read"].length == 0 || !validationMsg["last_name_read"].length == 0 || !validationMsg["birth"].length == 0 || !validationMsg["phone"].length == 0 || !validationMsg["address"].length == 0 || !validationMsg["credit_card_number"].length == 0 || !validationMsg["security_code"].length == 0 || !validationMsg["limit"].length == 0) {
    document.getElementById("e_email").innerHTML = validationMsg["email"];
    document.getElementById("e_password").innerHTML = validationMsg["password"];
    document.getElementById("e_confirmPass").innerHTML = validationMsg["confirmPass"];
    document.getElementById("e_first_name").innerHTML = validationMsg["first_name"];
    document.getElementById("e_last_name").innerHTML = validationMsg["last_name"];
    document.getElementById("e_first_name_read").innerHTML = validationMsg["first_name_read"];
    document.getElementById("e_last_name_read").innerHTML = validationMsg["last_name_read"];
    document.getElementById("e_birth").innerHTML = validationMsg["birth"];
    document.getElementById("e_phone").innerHTML = validationMsg["phone"];
    document.getElementById("e_address").innerHTML = validationMsg["address"];
    document.getElementById("e_credit_card_number").innerHTML = validationMsg["credit_card_number"];
    document.getElementById("e_security_code").innerHTML = validationMsg["security_code"];
    document.getElementById("e_limit").innerHTML = validationMsg["limit"];
    return false;
  } else {
    return true;
  }
};
