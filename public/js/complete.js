var obj = QueryString.parse(null,null,null,true);

//作品名
document.getElementById("movie").textContent = obj.movie;

//日
document.getElementById("date").textContent = obj.date;

//時間
document.getElementById("time").textContent = obj.time;

//シアター
document.getElementById("theater").textContent = obj.theater;


$.ajax({
  type: "POST",
  url: "",
  data: { movie: obj.movie, date: obj.date, time: obj.time, theater: obj.theater  }
}).done(function() {
  console.log('success!');
});