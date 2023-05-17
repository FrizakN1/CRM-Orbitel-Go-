let ths = document.querySelectorAll('th');
if (ths) {
    [].forEach.call(ths, function(th, index) {
        th.addEventListener('click', function() {
            [].forEach.call(ths, function (th) {
                th.style.color = "#b0b0b0";
            });
            th.style.color = "#0177fd";
            sortColumn(index);
        });
    });
}

let columnFlags = [0, 0, 0, 0, 0, 0];
let tbody = document.querySelector('tbody');

function sortColumn(index) {
    let criticalRows = tbody.querySelectorAll('.critical');
    let newCriticalRows = Array.from(criticalRows);

    let highRows = tbody.querySelectorAll('.high');
    let newHighRows = Array.from(highRows);

    let normalRows = tbody.querySelectorAll('.normal');
    let newNormalRows = Array.from(normalRows);

    let lowRows = tbody.querySelectorAll('.low');
    let newLowRows = Array.from(lowRows);

    newCriticalRows = sortColumnByPriority(newCriticalRows)
    newHighRows = sortColumnByPriority(newHighRows)
    newNormalRows = sortColumnByPriority(newNormalRows)
    newLowRows = sortColumnByPriority(newLowRows)

    function sortColumnByPriority(newRows) {
        newRows.sort(function(rowA, rowB) {
            let cellA = rowA.querySelectorAll('td')[index].innerHTML;
            let cellB = rowB.querySelectorAll('td')[index].innerHTML;

            if (columnFlags[index] === 1) {
                switch (true) {
                    case cellB > cellA: return 1;
                    case cellB < cellA: return -1;
                    case cellB === cellA: return 0;
                }
            } else {
                switch (true) {
                    case cellA > cellB: return 1;
                    case cellA < cellB: return -1;
                    case cellA === cellB: return 0;
                }
            }
        });
        return newRows
    }

    if (columnFlags[index] === 1) {
        columnFlags = [0, 0, 0, 0, 0, 0]
    } else {
        columnFlags = [0, 0, 0, 0, 0, 0]
        columnFlags[index] = 1;
    }

    tbody.innerHTML = '';

    newCriticalRows.forEach(function(newRow) {
        tbody.append(newRow);
    });
    newHighRows.forEach(function(newRow) {
        tbody.append(newRow);
    });
    newNormalRows.forEach(function(newRow) {
        tbody.append(newRow);
    });
    newLowRows.forEach(function(newRow) {
        tbody.append(newRow);
    });
}

function createApplicationsRows(title, applications) {
    console.log(4)
    tbody.innerHTML = "";
    let h2 = document.querySelector("h2")
    h2.innerHTML = title+" заявки - "+applications.length
    let criticalApplications = []
    let highApplications = []
    let normalApplications = []
    let lowApplications = []
    for (let row of applications) {
        let tr = document.createElement("tr")
        let column5Value
        if (row.Executor.ID) {
            column5Value = row.Executor.Name
        } else {
            column5Value = row.Status.Name
        }
        tr.innerHTML = "<td class=\"table-column1\">"+ row.ID +"</td>\n" +
            "<td class=\"table-column2\">"+ row.Abonent.Address +"</td>\n" +
            "<td class=\"table-column3\">"+ row.Abonent.Name +"</td>\n" +
            "<td class=\"table-column4\">"+ row.Description +"</td>\n" +
            "<td class=\"table-column5\">"+ column5Value +"</td>\n" +
            "<td class=\"table-column6\">"+ row.Date.slice(0,16) +"</td>"
        switch (row.Priority.ID) {
            case 1: tr.className = "critical"; criticalApplications.push(tr); break;
            case 2: tr.className = "high"; highApplications.push(tr); break;
            case 3: tr.className = "normal"; normalApplications.push(tr); break;
            case 4: tr.className = "low"; lowApplications.push(tr); break;
        }
    }
    for (let tr of criticalApplications) {
        tbody.append(tr)
    }
    for (let tr of highApplications) {
        tbody.append(tr)
    }
    for (let tr of normalApplications) {
        tbody.append(tr)
    }
    for (let tr of lowApplications) {
        tbody.append(tr)
    }
}

let applications = []
Send("GET", "/application/get-data", null, (res) => {
    if (res) {
        console.log(res)
        applications = res;
        switchTabs(null)
    }
})

let mainHead = document.querySelector(".main-head");
if (mainHead) {
    let tabs = mainHead.querySelector(".tabs").querySelectorAll("li");

    function defaultStyleTab() {
        for (let tab of tabs) {
            tab.classList.remove("active")
        }
    }

    for (let tab of tabs) {
        tab.onclick = (event) => {
            // defaultStyleTab()
            // switch (event.target.id) {
            //     case "status-1": createApplicationsRows("Текущие", applications.filter(item => item.Status.ID !== 2)); event.target.classList.add("active"); break
            //     case "status-2": createApplicationsRows("Новые", applications.filter(item => item.Status.ID === 1)); event.target.classList.add("active"); break
            //     case "status-3": createApplicationsRows("Закрытые", applications.filter(item => item.Status.ID === 2)); event.target.classList.add("active"); break
            //     case "status-4": createApplicationsRows("Все", applications); event.target.classList.add("active"); break
            // }
            switchTabs(event.target)
        }
    }

    function switchTabs(target) {
        console.log(2)
        if (target) {
            console.log(3)
            defaultStyleTab()
            let department = document.querySelector(".select").querySelector("div").getAttribute("value");
            if (department === "1") {
                switch (target.id) {
                    case "status-1": createApplicationsRows("Текущие", applications.filter(item => item.Status.ID !== 2)); target.classList.add("active"); break
                    case "status-2": createApplicationsRows("Новые", applications.filter(item => item.Status.ID === 1)); target.classList.add("active"); break
                    case "status-3": createApplicationsRows("Закрытые", applications.filter(item => item.Status.ID === 2)); target.classList.add("active"); break
                    case "status-4": createApplicationsRows("Все", applications); target.classList.add("active"); break
                }
            } else {
                switch (target.id) {
                    case "status-1": createApplicationsRows("Текущие", applications.filter(item => item.Status.ID !== 2 && item.Department.ID === Number(department))); target.classList.add("active"); break
                    case "status-2": createApplicationsRows("Новые", applications.filter(item => item.Status.ID === 1 && item.Department.ID === Number(department))); target.classList.add("active"); break
                    case "status-3": createApplicationsRows("Закрытые", applications.filter(item => item.Status.ID === 2 && item.Department.ID === Number(department))); target.classList.add("active"); break
                    case "status-4": createApplicationsRows("Все", applications.filter(item => item.Department.ID === Number(department))); target.classList.add("active"); break
                }
            }
        } else {
            for (let tab of tabs) {
                if (tab.className.includes("active")) {
                    switchTabs(tab)
                    console.log(tab)
                    break
                }
            }
        }
    }
}
//
// let selects = document.querySelectorAll(".select")
// if (selects) {
//     for (let select of selects) {
//         let selectButton = select.querySelector("div")
//         select.onclick = () => {
//             select.classList.toggle("active");
//         }
//         for (let i = 1; i <= select.querySelector("ul").childNodes.length-1; i=i+2) {
//             let item = select.querySelector("ul").childNodes[i]
//             item.onclick = () => {
//                 selectButton.innerHTML = item.innerHTML;
//                 selectButton.setAttribute("value",item.value);
//                 if (!selectButton.innerHTML.includes("Выбрать")) {
//                     selectButton.style.backgroundColor = '#0177fd';
//                     selectButton.style.color = "#ffffff"
//                 }
//
//             }
//         }
//     }
//     window.onclick = (event) => {
//         switch (event.target.parentNode) {
//             case selects[0]: disableSelect(selects[0], selects[1]);break;
//             case selects[1]: disableSelect(selects[1], selects[0]); break;
//             default: switch (event.target.parentNode.parentNode) {
//                 case selects[0]: disableSelect(selects[0], selects[1]); break;
//                 case selects[1]: disableSelect(selects[1], selects[0]); break;
//                 default:
//                     selects[0].classList.remove("active");
//                     selects[1].classList.remove("active");
//             }
//         }
//
//         function disableSelect (select1, select2) {
//             if (!select1.className.includes("active")) {
//                 select1.classList.remove("active");
//             }
//             select2.classList.remove("active")
//         }
//     }
// }