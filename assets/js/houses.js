let search = document.querySelector("#search");
if (search) {
    setTimeout(function () {
        let section = document.querySelector("section")
        let addresses = section.querySelectorAll("a");
        search.oninput = (event) => {
            if (event.target.value === "") {
                addresses.forEach(function(address) {
                    section.append(address);
                });
            } else {
                section.innerHTML = ""
                let newAddresses = Array.from(addresses).filter(item => item.innerHTML.toLowerCase().includes(event.target.value.toLowerCase()))
                newAddresses.forEach(function(newAddress) {
                    section.append(newAddress);
                });
            }
        }
    }, 100)
}

let housesList = document.querySelector("#houses_list");
if (housesList) {
    Send("GET", "/houses/get-all", null, (response) => {
        if (response) {
            for (let house of response) {
                let a = document.createElement("a");

                a.href = "/houses/view/"+house.ID;
                a.innerHTML = house.Name;
                housesList.append(a)
            }
        }
    })
}

function checkHouseForm(uri) {
    let street = document.querySelector("#Street")
    let number = document.querySelector("#Number")
    let internet = document.querySelector("#Internet")
    let tv = document.querySelector("#Tv")
    let telephony = document.querySelector("#Telephony")
    let nameMC = document.querySelector("#NameMC")
    let addressMC = document.querySelector("#AddressMC")
    let chairmanName = document.querySelector("#ChairmanName")
    let chairmanContact = document.querySelector("#ChairmanContact")
    let power = document.querySelector("#Power")
    let agreement = document.querySelector("#Agreement")

    let status = true

    if (street.value.length < 3) {
        status = false;
        street.parentNode.querySelector("p").style.backgroundColor = "#ff0000";
    } else {
        street.parentNode.querySelector("p").style.backgroundColor = "#0177fd";
    }

    if (number.value.length < 1) {
        status = false;
        number.parentNode.querySelector("p").style.backgroundColor = "#ff0000";
    } else {
        number.parentNode.querySelector("p").style.backgroundColor = "#0177fd";
    }

    if (internet.checked) {
        internet = 1;
    } else {
        internet = 0;
    }
    if (tv.checked) {
        tv = 1;
    } else {
        tv = 0;
    }
    if (telephony.checked) {
        telephony = 1;
    } else {
        telephony = 0;
    }
    if (agreement.checked) {
        agreement = 1;
    } else {
        agreement = 0;
    }

    if (status) {
        sendData(uri, street.value, number.value, internet, tv, telephony, nameMC.value, addressMC.value, chairmanName.value, chairmanContact.value, power.value, agreement)
    }
}

function sendData(uri, street, number, internet, tv, telephony, nameMC, addressMC, chairmanName, chairmanContact, power, agreement) {
    let createData = {
        Name: street.trim() + " " + number.trim(),
        Internet: internet,
        TV: tv,
        Telephony: telephony,
        NameMC: nameMC,
        AddressMC: addressMC,
        ChairmanName: chairmanName,
        ChairmanContact: chairmanContact,
        Power: Number(power),
        Agreement: agreement
    }

    Send("PUT", uri, createData, (response) => {
        if (response) {
            window.location.href = "/houses";
        }
    })
}

let houseForm = document.querySelector(".house-form");
if (houseForm) {
    let houseCreate = document.querySelector("#houseCreate");
    if (houseCreate) {
        houseCreate.onclick = () => {
            checkHouseForm("/houses/create")
        }
    }

    let houseUpdate = document.querySelector("#houseUpdate");
    if (houseUpdate) {
        houseUpdate.onclick = () => {
            let houseID = window.location.href.split("edit/")[1];
            checkHouseForm("/houses/edit/"+houseID)
        }
    }
}