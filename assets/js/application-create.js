let abonents = []
Send("GET", "/application/get-abonents", null, (res) => {
    if (res) {
        abonents = res
    }
})

let abonentInput = document.querySelector("#Abonent");
if (abonentInput) {
    let abonentList = document.querySelector("#abonent-list")

    abonentInput.oninput = (event) => {
        abonentList.innerHTML = ""
        if (event.target.value.length > 0) {
            let indices = abonents.reduce((result, item, index) => {
                if (item.Name.toLowerCase().includes(event.target.value.toLowerCase())) {
                    result.push(index);
                }
                return result;
            }, []);
            for (let index of indices) {
                let divAdbonent = document.createElement("div");
                let divName = document.createElement("div");
                let divAddress = document.createElement("div");

                divAdbonent.className = "abonent";

                divName.className = "name";
                divName.innerHTML = abonents[index].Name;

                divAddress.className = "address";
                divAddress.innerHTML = abonents[index].Address

                divAdbonent.append(divName, divAddress)
                abonentList.append(divAdbonent)

                divAdbonent.onclick = () => {
                    abonentInput.value = abonents[index].Name;
                    abonentInput.setAttribute("data-abonent_id", abonents[index].ID)
                    abonentList.innerHTML = "";
                }
            }
            if (abonentList.scrollHeight < 400) {
                abonentList.style.overflowY = 'hidden';
            } else {
                abonentList.style.overflowY = 'scroll';
            }
        }
    }
}

let disables = document.querySelectorAll(".disable");
if (disables) {
    disables[0].onclick = () => {
        disables[1].style.display = "block";
        disables[0].style.display = "none";
    }

    disables[1].onclick = () => {
        disables[0].style.display = "block";
        disables[1].style.display = "none";
    }

    let userCreateBtn = document.querySelector("#createUser");
    if (userCreateBtn) {
        userCreateBtn.onclick = () => {
            if (disables[0].style.display !== "block") {
                let abonentID = document.querySelector("#Abonent").getAttribute("data-abonent_id");
                let description = document.querySelector("#Description").value;
                let notes = document.querySelector("#Notes").value;
                let priority = document.querySelector("#Priority").getAttribute("value");
                let department = document.querySelector("#Department").getAttribute("value");

                let data = {
                    Description: description,
                    Notes: notes,
                    Priority: Number(priority),
                    Department: Number(department),
                    Abonent: {
                        ID: Number(abonentID)
                    }
                }

                Send("PUT", "/application/create", data, (res) => {
                    console.log(res)
                })
                console.log(123)
            }
        }
    }
}

let selects = document.querySelectorAll(".select")
if (selects) {
    for (let select of selects) {
        let selectButton = select.querySelector("div")
        select.onclick = () => {
            if (select.className.includes("active")) {
                select.classList.remove("active");
            } else {
                select.classList.add("active");
            }
        }
        for (let i = 1; i <= select.querySelector("ul").childNodes.length-1; i=i+2) {
            let item = select.querySelector("ul").childNodes[i]
            item.onclick = () => {
                selectButton.innerHTML = item.innerHTML;
                selectButton.setAttribute("value",item.value);
                if (!selectButton.innerHTML.includes("Выбрать")) {
                    selectButton.style.backgroundColor = '#0177fd';
                    selectButton.style.color = "#ffffff"
                }

            }
        }
    }
    window.onclick = (event) => {
        switch (event.target.parentNode) {
            case selects[0]: disableSelect(selects[0], selects[1]);break;
            case selects[1]: disableSelect(selects[1], selects[0]); break;
            default: switch (event.target.parentNode.parentNode) {
                case selects[0]: disableSelect(selects[0], selects[1]); break;
                case selects[1]: disableSelect(selects[1], selects[0]); break;
                default:
                    selects[0].classList.remove("active");
                    selects[1].classList.remove("active");
            }
        }

        function disableSelect (select1, select2) {
            if (!select1.className.includes("active")) {
                select1.classList.remove("active");
            }
            select2.classList.remove("active")
        }
    }
}