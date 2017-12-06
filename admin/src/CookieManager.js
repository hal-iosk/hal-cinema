import cookie from 'cookie'

if(location.pathname === "/admin") {
  const admin = cookie.parse("halCinemaAdmin");
  if(!admin || admin === "") location.href = "/admin/login";
} else if(location.pathname === "/admin/login") {
  const admin = cookie.parse("halCinemaAdmin");
  if(admin && admin !== "") location.href = "/admin";
} else {
  const user = cookie.parse("halCinemaUser");
  if(!user || user === "") location.href = "/";
}

