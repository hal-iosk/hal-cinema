var obj = QueryString.parse(null,null,null,true);

//作品名
document.getElementById("movie").textContent = obj.movie;

//日
document.getElementById("date").textContent = obj.date;

//時間
document.getElementById("time").textContent = obj.time;

//シアター
document.getElementById("theater").textContent = obj.theater;

