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
                divAddress.innerHTML = abonents[index].ActualAddress

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

function checkFastForm() {
    let abonentID = document.querySelector("#Abonent");
    let description = document.querySelector("#Description");
    let notes = document.querySelector("#Notes").value;
    let priority = document.querySelector("#Priority");
    let department = document.querySelector("#Department");

    let status = true;

    if (abonentID.getAttribute("data-abonent_id").length < 1) {
        status = false
        abonentID.style.border = "1px solid #ff0000";
    } else {
        abonentID.style.border = "1px solid #0177fd";
    }

    if (description.value.length < 1) {
        status = false
        description.style.border = "1px solid #ff0000";
    } else {
        description.style.border = "1px solid #0177fd";
    }

    if (priority.getAttribute("value").length < 1) {
        status = false
        priority.style.border = "1px solid #ff0000";
    } else {
        priority.style.border = "1px solid #0177fd";
    }

    if (department.getAttribute("value").length < 1 || department.getAttribute("value") === "0") {
        status = false
        department.style.border = "1px solid #ff0000";
    } else {
        department.style.border = "1px solid #0177fd";
    }

    if (status) {
        let data = {
            Description: description.value,
            Notes: notes,
            Priority: {
                ID: Number(priority.getAttribute("value"))
            },
            Department: {
                ID: Number(department.getAttribute("value"))
            },
            Abonent: {
                ID: Number(abonentID.getAttribute("data-abonent_id"))
            }
        }

        Send("PUT", "/application/create", data, (res) => {
            if (res) {
                window.location.href = "/application"
            }
        })
    }
}

function checkAlternativeForm() {
    let phone = document.querySelector("#phone");
    let name = document.querySelector("#name");
    let street = document.querySelector("#street");
    let houseNumber = document.querySelector("#house_number");
    let apartmentNumber = document.querySelector("#apartment_number");
    let description = document.querySelector("#Description");
    let notes = document.querySelector("#Notes").value;
    let priority = document.querySelector("#Priority");
    let department = document.querySelector("#Department");

    let status = true;

    if (phone.value.length < 1) {
        status = false
        name.style.border = "1px solid #ff0000";
    } else {
        name.style.border = "1px solid #0177fd";
    }

    if (name.value.length < 1) {
        status = false
        name.style.border = "1px solid #ff0000";
    } else {
        name.style.border = "1px solid #0177fd";
    }

    if (street.value.length < 1) {
        status = false
        street.style.border = "1px solid #ff0000";
    } else {
        street.style.border = "1px solid #0177fd";
    }

    if (houseNumber.value.length < 1) {
        status = false
        houseNumber.style.border = "1px solid #ff0000";
    } else {
        houseNumber.style.border = "1px solid #0177fd";
    }

    if (description.value.length < 1) {
        status = false
        description.style.border = "1px solid #ff0000";
    } else {
        description.style.border = "1px solid #0177fd";
    }

    if (priority.getAttribute("value").length < 1) {
        status = false
        priority.style.border = "1px solid #ff0000";
    } else {
        priority.style.border = "1px solid #0177fd";
    }

    if (department.getAttribute("value").length < 1) {
        status = false
        department.style.border = "1px solid #ff0000";
    } else {
        department.style.border = "1px solid #0177fd";
    }

    let addressSend;
    if (apartmentNumber.value.length > 0) {
        addressSend = street.value.trim()+" "+houseNumber.value.trim()+"-"+apartmentNumber.value.trim();
    } else {
        addressSend = street.value.trim()+" "+houseNumber.value.trim()
    }

    if (status) {
        let data = {
            Description: description.value,
            Notes: notes,
            Priority: {
                ID: Number(priority.getAttribute("value"))
            },
            Department: {
                ID: Number(department.getAttribute("value"))
            },
            Abonent: {
                Name: name.value.trim(),
                ActualAddress: addressSend,
                Phone: phone.value.trim(),
            }
        }

        Send("PUT", "/application/create-alternative", data, (res) => {
            if (res) {
                window.location.href = "/application"
            }
        })
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
                checkFastForm()
            } else {
               checkAlternativeForm()
            }
        }
    }
}