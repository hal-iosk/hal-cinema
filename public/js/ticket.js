//parse 復号化
var obj = QueryString.parse(null,null,null,true);
//作品名
document.getElementById("movie").textContent = obj.movie;
//日
document.getElementById("date").textContent = obj.date;
//時間
document.getElementById("time").textContent = obj.time;
//シアター
document.getElementById("theater").textContent = obj.theater;

//parameter
var param = QueryString.stringify({"movie":obj.movie,"date":obj.date,"time":obj.time,"theater":obj.theater},null, null,true);

document.getElementById("back").innerHTML = "<a href='seat?"+param+"' class='back'>戻る</a>";

document.getElementById("next").innerHTML = "<a href='payment?"+param+"' class='next'>次へ</a>";


//sessionStorageから席番号を取得
var seats = sessionStorage.getItem('seats').split(',');

//取得した席数分html生成
for (i = 0; i < seats.length; i++) {
	$('#seats').append("<div id ='seats_"+i+"'></div>");
  $("#seats_"+i+"").html("<div class='border'>"
                          +"<p>"
                            +"<img src='image/seat.png' width='50' height='70'>"
                            +"<span class='seats'>"+seats[i]+"</span>"
                            +"<button class='ticket' data-remodal-target='modal'>券種を選択してください。</button>"
                          +"</p>"
                          +"<div class='remodal' data-remodal-id='modal' id='modal-area'>"
                            +"<h2 class='title'>券種を選択してください。</h2>"
                            +"<ul class='list'>"
                            +"<li>"
                              +"<a href='#' onClick='getType(1," + i + ")'>"
                                +"<span class='name'>一般</span>"
                                +"<span class='price'>1,800円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#'>"
                                +"<span class='name'>大・専・高</span>"
                                +"<span class='price'>1,500円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#'>"
                                +"<span class='name'>中・小</span>"
                                +"<span class='price'>800円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#'>"
                                +"<span class='name'>幼児(3〜6歳)</span>"
                                +"<span class='price'>500円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#'>"
                                  +"<span class='name'>毎月1日</span>"
                                  +"<span class='price'>1,100円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#'>"
                                +"<span class='name'>レディースデー</span>"
                                +"<span class='price'>1,100円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#'>"
                                +"<span class='name'>シニアデー</span>"
                                +"<span class='price'>1,000円</span>"
                              +"</a>"
                            +"</li>"
                            +"<li>"
                              +"<a href='#'>"
                              +"<span class='name'>夫婦デー(2名)</span>"
                              +"<span class='price'>2,000円</span>"
                              +"</a>"
                            +"</li>"
                          +"</div>"
                        +"</div>");
}

function getType(type,seat) {
		switch (type) {
			case 1:
          $("#seats_"+seat+"").html("<div class='border'>"
                                      +"<p>"
                                        +"<img src='image/seat.png' width='50' height='70'>"
                                        +"<span class='seats'>"+seats[seat]+"</span>"
                                        +"<button class='selected' data-remodal-target='modal'>"
                                          +"<span class='name'>一般</span>"
                                          +"<span class='price'>1,800円</span>"
                                        +"</button>"
                                      +"</p>"
                                      +"<div class='remodal' data-remodal-id='modal' id='modal-area'>"
                                      +"<h2 class='title'>券種を選択してください。</h2>"
                                      +"<ul class='list'>"
                                        +"<li>"
                                          +"<a href='#'>"
                                            +"<span class='name'>一般</span>"
                                            +"<span class='price'>1,800円</span>"
                                          +"</a>"
                                        +"</li>"
                                        +"<li>"
                                          +"<a href='#'>"
                                            +"<span class='name'>大・専・高</span>"
                                            +"<span class='price'>1,500円</span>"
                                          +"</a>"
                                        +"</li>"
                                        +"<li>"
                                        +"<a href='#'>"
                                          +"<span class='name'>中・小</span>"
                                          +"<span class='price'>800円</span>"
                                        +"</a>"
                                      +"</li>"
                                      +"<li>"
                                        +"<a href='#'>"
                                          +"<span class='name'>幼児(3〜6歳)</span>"
                                          +"<span class='price'>500円</span>"
                                        +"</a>"
                                      +"</li>"
                                      +"<li>"
                                        +"<a href='#'>"
                                          +"<span class='name'>毎月1日</span>"
                                          +"<span class='price'>1,100円</span>"
                                        +"</a>"
                                      +"</li>"
                                      +"<li>"
                                        +"<a href='#'>"
                                          +"<span class='name'>レディースデー</span>"
                                          +"<span class='price'>1,100円</span>"
                                        +"</a>"
                                      +"</li>"
                                      +"<li>"
                                        +"<a href='#'>"
                                          +"<span class='name'>シニアデー</span>"
                                          +"<span class='price'>1,000円</span>"
                                        +"</a>"
                                      +"</li>"
                                      +"<li>"
                                        +"<a href='#'>"
                                          +"<span class='name'>夫婦デー(2名)</span>"
                                          +"<span class='price'>2,000円</span>"
                                        +"</a>"
                                      +"</li>"
                                    +"</div>"
                                  +"</div>"
                                );
				break;

			default:
				alert('error');
				break;
		}
}
