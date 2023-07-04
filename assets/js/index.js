let optional = document.querySelector(".optional");
if (optional) {
    for (let i = 1; i < 4; i=i+2) {
        optional.childNodes[i].onclick = () => {
            for (let j = 1; j < 4; j=j+2) {
                optional.childNodes[j].classList.remove("active")
            }
            optional.childNodes[i].classList.add("active");
        }
    }
}

Send("GET", "/get-events", null, (res) => {
    console.log(res)
    let tbody = document.querySelector("tbody");
    for (let event of res) {
        let tr = document.createElement("tr");
        tr.innerHTML = "<td class=\"table-column1\">"+event.Date.slice(10,16)+"</td>\n" +
                        "<td class=\"table-column2\">"+event.User.Name+"</td>\n" +
                        "<td class=\"table-column3\">"+event.Name+"</td>\n" +
                        "<td class=\"table-column4\">№"+event.Application.ID+"</td>\n" +
                        "<td class=\"table-column5\">"+event.Application.Description+"</td>\n" +
                        "<td class=\"table-column6\">"+event.Application.Abonent.ActualAddress+"</td>"
        tr.onclick = () => {
            window.location.href = "/application/view-"+event.Application.ID
        }
        tbody.append(tr)
    }
})