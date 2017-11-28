window.onload= function () {
    
        var selectYear = document.getElementById("year");
        var selectMonth = document.getElementById("month");
        var selectDate = document.getElementById("date");
        var date = new Date();
        var year = date.getFullYear();
        for(var minYear = 1900;minYear<=year;minYear++){
            var optionAddYear = document.createElement("option");
            optionAddYear.setAttribute("value", minYear);
            optionAddYear.innerHTML = minYear;
            selectYear.appendChild(optionAddYear);
        }
        for(var minMonth = 1;minMonth<=12;minMonth++){
            var optionAddMonth= document.createElement("option");
            optionAddMonth.setAttribute("value", minMonth);
            optionAddMonth.innerHTML = minMonth;
            selectMonth.appendChild(optionAddMonth);
        }
        for(var minDate = 1;minDate<=31;minDate++){
            var optionAddDate= document.createElement("option");
            optionAddDate.setAttribute("value", minDate);
            optionAddDate.innerHTML = minDate;
            selectDate.appendChild(optionAddDate);
        }
    }