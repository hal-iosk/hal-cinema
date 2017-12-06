//parse 復号化
var obj = QueryString.parse(null,null,null,true);

//URLパラメータを変更された時にリダイレクト
if ("" in obj === true ) {
	alert("パラメータが不正です。");
	location.href = "/";
}

var movie = obj.movie.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var date = obj.date.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var time = obj.time.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');
var theater = obj.theater.replace(/<("[^"]*"|'[^']*'|[^'">])*>/g,'');

if (movie == "" || date == "" || time == "" || theater == ""){
	alert("パラメータが不正です。");
	location.href = "/";
}

//作品名
document.getElementById("movie").textContent = movie;
//日
document.getElementById("date").textContent = date;
//時間
document.getElementById("time").textContent = time;
//シアター
document.getElementById("theater").textContent = theater;

//parameter
var param = QueryString.stringify({"movie":movie,"date":date,"time":time,"theater":theater},null, null,true);

document.getElementById("next").innerHTML = "<a href='seatSelection?"+param+"' class='next'>座席選択へ</a>";