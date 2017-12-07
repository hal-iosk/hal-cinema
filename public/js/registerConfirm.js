//parse 復号化
var obj = QueryString.parse(null,null,null,true);

//URLパラメータを変更された時にリダイレクト
if ("" in obj === true ) {
    alert("パラメータが不正です。");
    location.href = "/";
}

var email = obj.email.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var password = obj.password.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var first_name = obj.first_name.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var last_name = obj.last_name.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var first_name_read = obj.first_name_read.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var last_name_read = obj.last_name_read.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var year = obj.year.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var month = obj.month.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var date = obj.date.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var phone = obj.phone.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var address = obj.address.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var credit_card_number = obj.credit_card_number.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var security_code = obj.security_code.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var limitMonth = obj.limitMonth.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var limitYear = obj.limitYear.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');

if (email == "" || password == "" || first_name == "" || last_name == "" || first_name_read == "" || last_name_read == "" || year == "" || month == "" || date == "" || phone == "" || address == "" || credit_card_number == "" || security_code == "" || limitMonth == "" || limitYear == ""){
    alert("パラメータが不正です。");
    location.href = "/";
}
//生年月日のフォーマット
var birth = moment(year+'-'+month+'-'+date);
birth.format("YYYY-MM-DD");

//APIに投げるデータ
//メールアドレス
document.getElementsByClassName("email").value = email;
//パスワード
document.getElementsByClassName("password").value = password;
//性
document.getElementsByClassName("first_name").value = first_name;
//名
document.getElementsByClassName("last_name").value = last_name;
//フリガナ（セイ）
document.getElementsByClassName("first_name_read").value = first_name_read;
//フリガナ（メイ）
document.getElementsByClassName("last_name_read").value = last_name_read;
//生年月日(年)
document.getElementsByClassName("year").value = year;
//生年月日(月)
document.getElementsByClassName("month").value = month;
//生年月日(日)
document.getElementsByClassName("date").value = date;
//生年月日送るやつ
document.getElementsByClassName("birth").value = birth;
//電話番号
document.getElementsByClassName("phone").value = phone;
//郵便番号
document.getElementsByClassName("address").value = address;
//クレジットカード番号
document.getElementsByClassName("credit_card_number").value = credit_card_number;
//セキュリティコード
document.getElementsByClassName("security_code").value = security_code;
//有効期限(月)
document.getElementsByClassName("limitmonth").value = limitMonth;
//有効期限(年)
document.getElementsByClassName("limitYear").value = limitYear;


//表示データ
//メールアドレス
document.getElementsByClassName("v_email").innerHTML = email;
//パスワード
document.getElementsByClassName("v_password").innerHTML = password;
//性
document.getElementsByClassName("v_first_name").innerHTML = first_name;
//名
document.getElementsByClassName("v_last_name").innerHTML = last_name;
//フリガナ（セイ）
document.getElementsByClassName("v_first_name_read").innerHTML = first_name_read;
//フリガナ（メイ）
document.getElementsByClassName("v_last_name_read").innerHTML = last_name_read;
//生年月日(年)
document.getElementsByClassName("v_year").innerHTML = year;
//生年月日(月)
document.getElementsByClassName("v_month").innerHTML = month;
//生年月日(日)
document.getElementsByClassName("v_date").innerHTML = date;
//生年月日送るやつ
document.getElementsByClassName("v_birth").innerHTML = birth;
//電話番号
document.getElementsByClassName("v_phone").innerHTML = phone;
//郵便番号
document.getElementsByClassName("v_address").innerHTML = address;
//クレジットカード番号
document.getElementsByClassName("v_credit_card_number").innerHTML = credit_card_number;
//セキュリティコード
document.getElementsByClassName("v_security_code").innerHTML = security_code;
//有効期限(月)
document.getElementsByClassName("v_limitmonth").innerHTML = limitMonth;
//有効期限(年)
document.getElementsByClassName("v_limitYear").innerHTML = limitYear;















//parameter
// var param = QueryString.stringify({"movie":movie,"date":date,"time":time,"theater":theater},null, null,true);

// document.getElementById("back").innerHTML = "<a href='ticket?"+param+"' class='back'>戻る</a>";

// document.getElementById("next").innerHTML = "<a href='complete?"+param+"' class='next'>予約する</a>";


//sessionStorageから席番号を取得
// var seats = sessionStorage.getItem('seats').split(',');

//取得した席数分html生成
// for (i = 0; i < seats.length; i++) {
//  $('#seats').append("<div id ='seats_"+i+"'></div>");
//   $("#seats_"+i+"").html("<div class='border'><p><img src='image/seat.png' width='50' height='70'><span class='seats'>"+seats[i]+"</span><span class='ticket'>券種：一般</span></p></div>");
// }




// $.ajax({
//   type: "POST",
//   url: "",
//   data: 
//   { 
//      movie: movie, 
//      date: date, 
//      time: time, 
//      theater: theater,
//      seats : seats
//   }
// }).done(function() {
//   console.log('success!');
// });

