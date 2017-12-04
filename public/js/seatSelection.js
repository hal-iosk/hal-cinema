//選択座席を取得
function getSeatSelected() {
	//parse 符号化 
	var obj = QueryString.parse(null,null,null,true);

	//parameter
	var param = QueryString.stringify({"movie":obj.movie,"date":obj.date,"time":obj.time,"theater":obj.theater}, null, null, true);

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


