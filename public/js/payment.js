//parse 復号化
var obj = QueryString.parse(null,null,null,true);

//parameter
var param = QueryString.stringify({"movie":obj.movie,"date":obj.date,"time":obj.time,"theater":obj.theater},null, null,true);

document.getElementById("back").innerHTML = "<a href='ticket?"+param+"' class='back'>戻る</a>";

document.getElementById("next").innerHTML = "<a href='confirm?"+param+"' class='next' onclick='paymentSelect()'>次へ</a>";

function paymentSelect() {
	var element = document.getElementById('payment').payment;

	if (element.value == 1){
	var payment  =  "<div class='border'><p><span class='payment'>クレジットカード</span><span class='detail'>xxxx-xxxx-xxxx-1111</span></p></div>";
	sessionStorage.setItem('payment',payment);
	} 
	else{
	var payment  =  "<div class='border'><p><span class='payment'>現金</span><span class='detail'>当日、窓口にてお支払いください。</span></p></div>";
	sessionStorage.setItem('payment',payment);
	}
}


