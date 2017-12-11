window.onload= function () {
    
        var selectYear = document.getElementById("year");
        var selectMonth = document.getElementById("month");
        var selectDate = document.getElementById("date");
        var crMonth = document.getElementById("limitMonth");
        var crYear = document.getElementById("limitYear");
        var date = new Date();
        var year = date.getFullYear();
        // 生年月日（年）
        for(var minYear = 1900;minYear<=year;minYear++){
            var optionAddYear = document.createElement("option");
            optionAddYear.setAttribute("value", minYear);
            optionAddYear.innerHTML = minYear;
            selectYear.appendChild(optionAddYear);
        }
        // 生年月日（月）
        for(var minMonth = 1;minMonth<=12;minMonth++){
            var optionAddMonth= document.createElement("option");
            optionAddMonth.setAttribute("value", minMonth);
            optionAddMonth.innerHTML = minMonth;
            selectMonth.appendChild(optionAddMonth);
        }
        // 生年月日（日）
        for(var minDate = 1;minDate<=31;minDate++){
            var optionAddDate= document.createElement("option");
            optionAddDate.setAttribute("value", minDate);
            optionAddDate.innerHTML = minDate;
            selectDate.appendChild(optionAddDate);
        }
        
    }
    