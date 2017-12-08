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

// var param = QueryString.stringify({"movie":movie,"date":date,"time":time,"theater":theater},null, null,true);

if (email == "" || password == "" || first_name == "" || last_name == "" || first_name_read == "" || last_name_read == "" || year == "" || month == "" || date == "" || phone == "" || address == "" || credit_card_number == "" || security_code == "" || limitMonth == "" || limitYear == ""){
    alert("パラメータが不正です。");
    location.href = "/";
}
if(month.length == 1){
    month = "0"+month;
}
if(date.length == 1){
    date = "0"+date;
}
if(limitYear.length == 1){
    limitYear = "0"+limitYear;
}
if(limitMonth.length == 1){
    limitMonth = "0"+limitMonth;
}

//生年月日のフォーマット
var birth =year+'/'+month+'/'+date;

//生年月日のフォーマット
var credit_card_limit = limitYear+"/"+limitMonth
//APIに投げるデータ
//メールアドレス
document.getElementById("email").value= email;
//パスワード
document.getElementById("password").value = password;
//性
document.getElementById("first_name").value = first_name;
//名
document.getElementById("last_name").value = last_name;
//フリガナ（セイ）
document.getElementById("first_name_read").value = first_name_read;
//フリガナ（メイ）
document.getElementById("last_name_read").value = last_name_read;
//生年月日送るやつ
document.getElementById("birth").value = birth;
//電話番号
document.getElementById("phone").value = phone;
//郵便番号
document.getElementById("address").value = address;
//クレジットカード番号
document.getElementById("credit_card_number").value = credit_card_number;
//セキュリティコード
document.getElementById("security_code").value = security_code;
//有効期限(月)
document.getElementById("limitMonth").value = limitMonth;
//有効期限(年)
document.getElementById("limitYear").value = limitYear;
//有効期限(年)
document.getElementById("credit_card_limit").value = credit_card_limit;

//表示データ
//メールアドレス
document.getElementById("v_email").textContent = email;
//パスワード
document.getElementById("v_password").textContent= password;
//性
document.getElementById("v_first_name").textContent = first_name;
//名
document.getElementById("v_last_name").textContent = last_name;
//フリガナ（セイ）
document.getElementById("v_first_name_read").textContent = first_name_read;
//フリガナ（メイ）
document.getElementById("v_last_name_read").textContent = last_name_read;
//生年月日(年)
document.getElementById("year").textContent = year;
//生年月日(月)
document.getElementById("month").textContent = month;
//生年月日(日)
document.getElementById("date").textContent = date;
//生年月日送るやつ
// document.getElementsByClassName("v_birth").innerHTML = birth;
//電話番号
document.getElementById("v_phone").textContent = phone;
//郵便番号
document.getElementById("v_address").textContent = address;
//クレジットカード番号
document.getElementById("v_credit_card_number").textContent = credit_card_number;
//セキュリティコード
document.getElementById("v_security_code").textContent = security_code;
//有効期限(月)
document.getElementById("v_limitMonth").textContent = limitMonth;
//有効期限(年)
document.getElementById("v_limitYear").textContent = limitYear;















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

