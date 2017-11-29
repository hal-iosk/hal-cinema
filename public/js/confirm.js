//sessionStorageから席番号を取得
var seats = sessionStorage.getItem('seats').split(',');

//取得した席数分html生成
for (i = 0; i < seats.length; i++) {
	$('#seats').append("<div id ='seats_"+i+"'></div>");
  $("#seats_"+i+"").html("<div class='border'><p><img src='image/seat.png' width='50' height='70'><span class='seats'>"+seats[i]+"</span><span class='ticket'>券種：一般</span></p></div>");
}

//sessionStorageからpaymentを取得
var payment = sessionStorage.getItem('payment');

$("#payment").html(payment);