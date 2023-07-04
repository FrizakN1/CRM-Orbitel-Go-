function checkAbonentForm(uri) {
    let name = document.querySelector('#Name');
    let registeredAddress = document.querySelector("#RegisteredAddress");
    let actualAddress = document.querySelector("#ActualAddress");
    let ipAddress = document.querySelector("#IPAddress");
    let phone = document.querySelector("#Phone");
    let passportSeries = document.querySelector("#PassportSeries");
    let passportNumber = document.querySelector("#PassportNumber");
    let contractNumber = generateContractNumber();

    let status = true;

    if (name.value.length < 1) {
        status = false
        name.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
    } else {
        name.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
    }


    if (registeredAddress.value.length < 3) {
        status = false
        registeredAddress.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
    } else {
        registeredAddress.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
    }

    if (actualAddress.value.length < 3) {
        status = false
        actualAddress.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
    } else {
        actualAddress.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
    }

    if ((ipAddress.value.split(".")).length !== 4) {
        status = false
        ipAddress.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
    } else {
        ipAddress.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
    }

    if (phone.value.length < 3) {
        status = false
        phone.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
    } else {
        phone.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
    }

    if (passportSeries.value.length !== 4) {
        status = false
        passportSeries.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
    } else {
        passportSeries.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
    }

    if (passportNumber.value.length !== 6) {
        status = false
        passportNumber.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
    } else {
        passportNumber.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
    }

    if (status) {
        sendData(uri,name.value,registeredAddress.value,actualAddress.value,ipAddress.value,phone.value,passportSeries.value,passportNumber.value,contractNumber)
    }
}

function sendData(uri,name,registeredAddress,actualAddress,ipAddress,phone,passportSeries,passportNumber,contractNumber ) {
    let createData = {
        Name: name.trim(),
        RegisteredAddress: registeredAddress,
        ActualAddress: actualAddress,
        IPAddress: ipAddress,
        Phone: phone,
        passportSeries: passportSeries,
        passportNumber: passportNumber,
        ContractNumber: contractNumber
    }

    Send("PUT", uri, createData, (response) => {
        if (response) {
            window.location.href = "/abonents";
        }
    })
}

let abonentCreate = document.querySelector("#createAbonent");
if (abonentCreate) {
    abonentCreate.onclick = function () {
        checkAbonentForm('/abonents/create')
    }
}

let abonentUpdate = document.querySelector("#updateAbonent");
if (abonentUpdate) {
    abonentUpdate.onclick = function () {
        let abonentID = window.location.href.split("edit/")[1];
        checkAbonentForm('/abonents/update/'+abonentID)
    }
}


function generateContractNumber() {
    let currentDate = new Date();
    let year = currentDate.getFullYear().toString().slice(-2);
    let month = (currentDate.getMonth() + 1).toString().padStart(2, '0');
    let day = currentDate.getDate().toString().padStart(2, '0');
    let hours = currentDate.getHours().toString().padStart(2, '0');

    let randomTwoDigitNumber = Math.floor(Math.random() * 90) + 10;

    let contractNumber = year + month + day + hours +randomTwoDigitNumber;

    return contractNumber;
}