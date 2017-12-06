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




//選択座席を取得
function getSeatSelected() {
	//parse 符号化 
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

	//parameter
	var param = QueryString.stringify({"movie":movie,"date":date,"time":time,"theater":theater},null, null,true);

	var elements = document.getElementsByClassName('row__seat--selected');
	if (elements.length != 0 && elements.length <=6) {
		var seatNumbers = Object.keys(elements).map((index) => {return elements[index].dataset.tooltip});
		sessionStorage.setItem('seats',seatNumbers);
		location.href = "ticket?"+param+"";
	}
	else if(elements.length == 0) {
		alert("座席を選択してください。");
	}
	else {
		alert("座席を一度に選択できるのは6席までです。");
	}
}


