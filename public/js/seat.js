//parse 復号化
var obj = QueryString.parse(null,null,null,true);

//URLパラメータを変更された時にリダイレクト
if (obj.theater === undefined) {
	location.href = "/";
}

//作品名
document.getElementById("movie").textContent = obj.movie;
//日
document.getElementById("date").textContent = obj.date;
//時間
document.getElementById("time").textContent = obj.time;
//シアター
document.getElementById("theater").textContent = obj.theater;

//parameter
var param = QueryString.stringify({"movie":obj.movie,"date":obj.date,"time":obj.time,"theater":obj.theater},null,null,true);

document.getElementById("next").innerHTML = "<a href='seatSelection?"+param+"' class='next'>座席選択へ</a>";