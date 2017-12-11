if(!CookieDoc.hasItem("halCinemaUser")){
    location.href = "/login"
}

function logout() {
    CookieDoc.removeItem("halCinemaUser")
    location.href = "/"
}

var token = CookieDoc.getItem("halCinemaUser")

$.get(`/api/user?token=${token}`)
.done(function (data) {
  var name = `${data.custome.first_name} ${data.custome.last_name}`
  $("#username").text(name)
  $(".point_count").text(data.custome.point_count)

})
.fail(function (data, textStatus, jqXHR) {
    alert("error");
})

function html() {
    const reserveList = {}
    reserveHistory(reserveList)
    setTimeout(() => {
        Object.keys(reserveList).map((reserve) => {
            const reDate = moment(reserveList[reserve].reservedDate).format("YYYY/MM/DD")
            const start_time = moment(reserveList[reserve].start_time).format("MM/DD HH:mm")
            const end_time = moment(reserveList[reserve].start_time).add(reserveList[reserve].watch_time, "m").format("MM/DD HH:mm")
            $(".historyTable").append(generateHTML(reDate, reserveList[reserve].title, start_time, end_time, reserveList[reserve].watch_time))
        })
    }, 3 * 1000)
}
html()

function generateHTML(reservedDate, title, start_time, end_time, watch_time) {
    return "<tr>"
        +`<td class=''>${reservedDate}</td>`
        +`<td>${title}</td>`
        +`<td>${start_time}</td>`
        +`<td>${end_time}</td>`
        +`<td>${watch_time}分</td>`
    +"</tr>"
}

function reserveHistory(reserveList) {
    $.ajax({
        type: "GET",
        url: `/api/user/reserved?token=${token}`
    })
    .done((data) => {
        data.reserved.map((reserve) => {
            reserveList[reserve.ID] = {
                reservedDate: reserve.CreatedAt,
                start_time: null,
                watch_time: null,
                title: ""
            }
            getReserveDetail(reserve.schedule_id, reserve.ID, reserveList)
        })
    })
}

function getReserveDetail(id, reserve_id, reserveList) {
    $.ajax({
        type: "GET",
        url: `/api/schedule/${id}`
    })
    .done((res) => {
        reserveList[reserve_id].start_time = res.start_time
        getMovieDetail(res.movie_id, reserve_id, reserveList)
    })
}

function getMovieDetail(id, reserve_id, reserveList) {
    $.ajax({
        type: "GET",
        url: `/api/movie/${id}`
    })
    .done((res) => {
        reserveList[reserve_id].watch_time = res.watch_time
        reserveList[reserve_id].title = res.movie_name
    })
}
//schedule_id