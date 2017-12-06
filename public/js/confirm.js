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

document.getElementById("back").innerHTML = "<a href='ticket?"+param+"' class='back'>戻る</a>";

document.getElementById("next").innerHTML = "<a href='complete?"+param+"' class='next'>予約する</a>";


//sessionStorageから席番号を取得
var seats = sessionStorage.getItem('seats').split(',');

//取得した席数分html生成
for (i = 0; i < seats.length; i++) {
	$('#seats').append("<div id ='seats_"+i+"'></div>");
  $("#seats_"+i+"").html("<div class='border'><p><img src='image/seat.png' width='50' height='70'><span class='seats'>"+seats[i]+"</span><span class='ticket'>券種：一般</span></p></div>");
}




// $.ajax({
//   type: "POST",
//   url: "",
//   data: 
//   { 
//   	movie: movie, 
//   	date: date, 
//   	time: time, 
//   	theater: theater,
//   	seats : seats
//   }
// }).done(function() {
//   console.log('success!');
// });
