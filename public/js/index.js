$(function(){
	// #で始まるリンクをクリックしたら実行されます
	$('a[href^="#"]').click(function() {
	  // スクロールの速度
	  var speed = 400; // ミリ秒で記述
	  var href= $(this).attr("href");
	  var target = $(href == "#" || href == "" ? 'html' : href);
	  var position = target.offset().top;
	  $('body,html').animate({scrollTop:position}, speed, 'swing');
	  return false;
	});
  });

  $(function() {
    $('.center-item').slick({
          infinite: true,
          dots:true,
          slidesToShow: 1,
          centerMode: true, //要素を中央寄せ
          centerPadding:'130px', //両サイドの見えている部分のサイズ
          autoplay:true, //自動再生
          responsive: [{
               breakpoint: 480,
                    settings: {
                         centerMode: false,
               }
          }]
     });
});

function mapInit() {
    var centerPosition = new google.maps.LatLng(34.810666, 135.558743);
    var option = {
        zoom : 15,
        center : centerPosition,
        mapTypeId : google.maps.MapTypeId.ROADMAP
    };


    //地図本体描画
    var googlemap = new google.maps.Map(document.getElementById("mapField"), option);

    //マーカーオプションの指定
    var markerOption = {
        position : centerPosition, //マーカーを表示させる座標
        map : googlemap,  //マーカーを表示させる地図
    };

    //情報ウインドウオプションの指定
    var infoWindowOption = {
    position : centerPosition,  //中心座標
    content : "HALCINEMA"  //ウインドウ内に表示する文字列
    };
 
    //情報ウインドウ追加
    var infoWindow = new google.maps.InfoWindow(infoWindowOption);
    infoWindow.open(googlemap);

    //マーカー追加
    var marker = new google.maps.Marker(markerOption);

}
 
mapInit();

var isLogin = AuthManager.isLogin();

if(isLogin){
  var login = document.getElementById("islogin");
  login.innerHTML = "<a href='/mypage' class='login'><span>マイページ</span></a>";

}
