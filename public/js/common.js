function CookieManager() {}
CookieManager.prototype.getCookie = function(name) {
  var result = null;
  var cookieName = name + '=';
  var allcookies = document.cookie;

  var position = allcookies.indexOf(cookieName);
  if(position != -1) {
    var startIndex = position + cookieName.length;
    var endIndex = allcookies.indexOf( ';', startIndex );
    if(endIndex == -1) {
      endIndex = allcookies.length;
    }
    result = decodeURIComponent(allcookies.substring(startIndex, endIndex));
  }

  return result;
}

window.cookieManager = new CookieManager();

function AuthManager() {}
AuthManager.prototype.isLogin = function() {
  var cookieManager = new CookieManager();
  var value = cookieManager.getCookie("halCinemaUser");
  return !(!value || value === "")
}

window.AuthManager = new AuthManager();
