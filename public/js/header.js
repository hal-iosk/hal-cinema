$(document).ready(function() {
    $(".drawer").drawer();
});

//ログイン制御なう
var isLogin = AuthManager.isLogin();

if(isLogin){
	var navLogin = document.getElementById("navLogin");
	navLogin.innerHTML = "<a href='/view/mypage.html'><img src='/image/mypage.png' width='auto' height='auto'><span>マイページ</span></a>";

	var navPaper = document.getElementById("navPaper");
	navPaper.innerHTML = "<a href='/view/mypage.html'><img src='/image/paper.png' width='auto' height='auto'><span>予約履歴</span></a>";

	var navP = document.getElementById("navP");
	navP.innerHTML = "<a href='/view/mypage.html'><img src='/image/p.png' width='auto' height='auto'><br><span>チェックイン</span></a>";
}