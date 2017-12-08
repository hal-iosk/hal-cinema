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

//dataの初期化
var data = [];

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

document.getElementById("back").innerHTML = "<a href='seat?"+param+"' class='back'>戻る</a>";

document.getElementById("next").innerHTML = "<a href='#' class='next' onClick=typeSelectCheck()>次へ</a>";


//sessionStorageから席番号を取得
var seats = sessionStorage.getItem('seats').split(',');

//取得した席数分html生成
for (i = 0; i < seats.length; i++) {
	$('#seats').append("<div id ='seats_"+i+"'></div>");
  $("#seats_"+i+"").html("<div class='border'>"
                          +"<p>"
                            +"<img src='image/seat.png' width='50' height='70'>"
                            +"<span class='seats'>"+seats[i]+"</span>"
                            +"<button class='ticket' data-remodal-target='modal" + i + "'>券種を選択してください。</button>"
                          +"</p>"
                          +"<div class='remodal' data-remodal-id='modal" + i + "' id='modal-area'>"
                            +"<h2 class='title'>券種を選択してください。</h2>"
                            +"<ul class='list'>"
                            +"<li>"
                              +"<a href='#' onClick='getType(1," + i + ")'>"
                                +"<span class='name'>一般</span>"
                                +"<span class='price'>1,800円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#' onClick='getType(2," + i + ")'>"
                                +"<span class='name'>大・専・高</span>"
                                +"<span class='price'>1,500円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#' onClick='getType(3," + i + ")'>"
                                +"<span class='name'>中・小</span>"
                                +"<span class='price'>800円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#' onClick='getType(4," + i + ")'>"
                                +"<span class='name'>幼児(3〜6歳)</span>"
                                +"<span class='price'>500円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#' onClick='getType(5," + i + ")'>"
                                  +"<span class='name'>毎月1日</span>"
                                  +"<span class='price'>1,100円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#' onClick='getType(6," + i + ")'>"
                                +"<span class='name'>レディースデー</span>"
                                +"<span class='price'>1,100円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#' onClick='getType(7," + i + ")'>"
                                +"<span class='name'>シニアデー</span>"
                                +"<span class='price'>1,000円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#' onClick='getType(8," + i + ")'>"
                              +"<span class='name'>夫婦デー(2名)</span>"
                              +"<span class='price'>2,000円</span>"
                              +"</a>"
                            +"</li>"
                          +"</div>"
                        +"</div>"
                        );
}

function getType(type,seat) {
		switch (type) {
			case 1:
          $("#seats_"+seat+"").html("<div class='border'>"
                                      +"<p>"
                                        +"<img src='image/seat.png' width='50' height='70'>"
                                        +"<span class='seats'>"+seats[seat]+"</span>"
                                        +"<button class='selected' data-remodal-target='modal" + seat + "'>"
                                          +"<span class='name'>一般</span>"
                                          +"<span id='price'>1,800円</span>"
                                        +"</button>"
                                      +"</p>"
                                  +"</div>"
                                );

          var newData = data.filter(function(item, index){
          if (item.seat != seats[seat]) return true;
          });

          data = newData;

          var addData = { seat : seats[seat], price : 1800 }
          data.push(addData);
				break;

      case 2:
          $("#seats_"+seat+"").html("<div class='border'>"
                                      +"<p>"
                                        +"<img src='image/seat.png' width='50' height='70'>"
                                        +"<span class='seats'>"+seats[seat]+"</span>"
                                        +"<button class='selected' data-remodal-target='modal" + seat + "'>"
                                          +"<span class='name'>大・専・高</span>"
                                          +"<span id='price'>1,500円</span>"
                                        +"</button>"
                                      +"</p>"
                                  +"</div>"
                                );

          var newData = data.filter(function(item, index){
          if (item.seat != seats[seat]) return true;
          });

          data = newData;

          var addData = { seat : seats[seat], price : 1500 }
          data.push(addData);
        break;

        case 3:
          $("#seats_"+seat+"").html("<div class='border'>"
                                      +"<p>"
                                        +"<img src='image/seat.png' width='50' height='70'>"
                                        +"<span class='seats'>"+seats[seat]+"</span>"
                                        +"<button class='selected' data-remodal-target='modal" + seat + "'>"
                                          +"<span class='name'>中・小</span>"
                                          +"<span class='price'>800円</span>"
                                        +"</button>"
                                      +"</p>"
                                  +"</div>"
                                );
          var newData = data.filter(function(item, index){
          if (item.seat != seats[seat]) return true;
          });

          data = newData;

          var addData = { seat : seats[seat], price : 800 }
          data.push(addData);
        break;

        case 4:
          $("#seats_"+seat+"").html("<div class='border'>"
                                      +"<p>"
                                        +"<img src='image/seat.png' width='50' height='70'>"
                                        +"<span class='seats'>"+seats[seat]+"</span>"
                                        +"<button class='selected' data-remodal-target='modal" + seat + "'>"
                                          +"<span class='name'>幼児(3〜6歳)</span>"
                                          +"<span class='price'>500円</span>"
                                        +"</button>"
                                      +"</p>"
                                  +"</div>"
                                );
          var newData = data.filter(function(item, index){
          if (item.seat != seats[seat]) return true;
          });

          data = newData;

          var addData = { seat : seats[seat], price : 500 }
          data.push(addData);
        break;

        case 5:
          $("#seats_"+seat+"").html("<div class='border'>"
                                      +"<p>"
                                        +"<img src='image/seat.png' width='50' height='70'>"
                                        +"<span class='seats'>"+seats[seat]+"</span>"
                                        +"<button class='selected' data-remodal-target='modal" + seat + "'>"
                                          +"<span class='name'>毎月1日</span>"
                                          +"<span class='price'>1,100円</span>"
                                        +"</button>"
                                      +"</p>"
                                  +"</div>"
                                );
          var newData = data.filter(function(item, index){
          if (item.seat != seats[seat]) return true;
          });

          data = newData;

          var addData = { seat : seats[seat], price : 1100 }
          data.push(addData);
        break;

        case 6:
          $("#seats_"+seat+"").html("<div class='border'>"
                                      +"<p>"
                                        +"<img src='image/seat.png' width='50' height='70'>"
                                        +"<span class='seats'>"+seats[seat]+"</span>"
                                        +"<button class='selected' data-remodal-target='modal" + seat + "'>"
                                          +"<span class='name'>レディースデー</span>"
                                          +"<span class='price'>1,100円</span>"
                                        +"</button>"
                                      +"</p>"
                                  +"</div>"
                                );
          var newData = data.filter(function(item, index){
          if (item.seat != seats[seat]) return true;
          });

          data = newData;

          var addData = { seat : seats[seat], price : 1100 }
          data.push(addData);
        break;

        case 7:
          $("#seats_"+seat+"").html("<div class='border'>"
                                      +"<p>"
                                        +"<img src='image/seat.png' width='50' height='70'>"
                                        +"<span class='seats'>"+seats[seat]+"</span>"
                                        +"<button class='selected' data-remodal-target='modal" + seat + "'>"
                                          +"<span class='name'>シニアデー</span>"
                                          +"<span class='price'>1,000円</span>"
                                        +"</button>"
                                      +"</p>"
                                  +"</div>"
                                );
          var newData = data.filter(function(item, index){
          if (item.seat != seats[seat]) return true;
          });

          data = newData;

          var addData = { seat : seats[seat], price : 1000 }
          data.push(addData);
        break;

        case 8:
          $("#seats_"+seat+"").html("<div class='border'>"
                                      +"<p>"
                                        +"<img src='image/seat.png' width='50' height='70'>"
                                        +"<span class='seats'>"+seats[seat]+"</span>"
                                        +"<button class='selected' data-remodal-target='modal" + seat + "'>"
                                          +"<span class='name'>夫婦デー(2名)</span>"
                                          +"<span class='price'>2,000円</span>"
                                        +"</button>"
                                      +"</p>"
                                  +"</div>"
                                );
          var newData = data.filter(function(item, index){
          if (item.seat != seats[seat]) return true;
          });

          data = newData;

          var addData = { seat : seats[seat], price : 2000 }
          data.push(addData);
        break;

			default:
				alert('error');
				break;
		}
}

function typeSelectCheck() {
    if(data.length != seats.length){
      alert("券種を選択してください。");
    } else {
        for (i = 0; i < data.length; i++) {
          sessionStorage.setItem('a',data[i].seat);
          sessionStorage.setItem('b',data[i].price);
        }
      location.href = "confirm?"+param+"";
    }
}

