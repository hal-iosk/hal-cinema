$(document).ready(function() {
    $(".drawer").drawer();
});

//ログイン制御なう
var isLogin = AuthManager.isLogin();

if(isLogin){//ログインしてるとき
	var navLogin = document.getElementById("navLogin");
	navLogin.innerHTML = "<a href='/mypage'><img src='/image/mypage.png' width='auto' height='auto'><span>マイページ</span></a>";

	var navPaper = document.getElementById("navPaper");
	navPaper.innerHTML = "<a href='/mypage'><img src='/image/paper.png' width='auto' height='auto'><span>予約履歴</span></a>";

	var navP = document.getElementById("navP");
	navP.innerHTML = "<a href='/mypage'><img src='/image/p.png' width='auto' height='auto'><br><span>コイン確認</span></a>";
}