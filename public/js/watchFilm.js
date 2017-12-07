window.onload = function () {
  var month = new Date();
  var date = new Date();
  console.log(month.getMonth() + 1);
  console.log(date.getDate());
  
  var scheduleTable = document.getElementById("scheduleTable");

  for (var scheduleWeek = 0; scheduleWeek <= 6; scheduleWeek++) {
    var scheduleDate = document.createElement("li");
    scheduleDate.classList.add("scheduleDate");
    scheduleDate.classList.add((month.getMonth() + 1) + "/" + (date.getDate() + scheduleWeek));
    scheduleDate.innerHTML = (month.getMonth() + 1) + "/" + (date.getDate() + scheduleWeek);
    
    scheduleTable.appendChild(scheduleDate);
    var day = document.getElementsByClassName((month.getMonth() + 1) + "/" + (date.getDate() + scheduleWeek).innerHTML);
    scheduleDate.onclick = function(){ 
        paging(day);
        };
  }

}
var de = document.getElementsByClassName("scheduleDate").innerHTML;
console.log(de);
function paging(day){
 alert(day);
}