if(CookieDoc.hasItem("halCinemaUser")){
    location.href = "/login"
}

// logout
// CookieDoc.removeItem("halCinemaUser")
// location.href = "/"

// http request
// var token = CookieDoc.getItem("halCinemaUser")
// `/api/xxx?token=${token}`

// user model GET
// `/api/user?token=${token}`

// reserved
// `/api/user/reserved?token=${token}`