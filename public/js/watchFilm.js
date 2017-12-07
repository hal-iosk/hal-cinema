window.onload = function () {
  var month = new Date();
  var date = new Date();
  console.log(month.getMonth() + 1);
  console.log(date.getDate());

  var scheduleTable = document.getElementById("scheduleTable");

  for (var scheduleWeek = 0; scheduleWeek <= 6; scheduleWeek++) {
    var scheduleDate = document.createElement("li");
    scheduleDate.classList.add("scheduleDate");
    if(scheduleWeek<=2){
        scheduleDate.classList.add("three");
    }else{
        scheduleDate.classList.add("other"); 
    }
    scheduleDate.innerHTML = (month.getMonth() + 1) + "/" + (date.getDate() + scheduleWeek);
    scheduleTable.appendChild(scheduleDate);
  }
}
