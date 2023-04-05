let loginBtn = document.querySelector('#btn');
if (loginBtn) {
    loginBtn.onclick = function (event) {
        let login = document.querySelector('#login').value;
        let password = document.querySelector('#password').value;

        if (login && password) {
            const loginData = {
                Login: login,
                Password: password,
            }

            Send('POST', '/users/login', loginData, (response) => {
                if (response) {
                    window.location.href = "/";
                }
            })
        } else {
            let loginAlert = document.querySelector('#login').closest('label').querySelector('p');
            let passwordAlert = document.querySelector('#password').closest('label').querySelector('p');
            if (!login) {
                loginAlert.textContent = "Необходимо заполнить поле «Логин»";
            } else {
                loginAlert.textContent = "";
            }
            if (!password) {
                passwordAlert.textContent = "Необходимо заполнить поле «Пароль»";
            } else {
                passwordAlert.textContent = "";
            }
        }
    }
}

let userCreate = document.querySelector("#createUser");
if (userCreate) {
    userCreate.onclick = function (event) {
        let login = document.querySelector('#Login').value
        let name = document.querySelector('#Name').value
        let password = document.querySelector('#Password').value
        let role = document.querySelector('#Role').getAttribute("value")
        let department = document.querySelector('#Department').value

        let createData = {
            Login: String(login),
            Name: String(name),
            Password: String(password),
            Role: Number(role),
            Department: Number(department),
        }

        Send('PUT', '/users/create', createData, (response) => {
            if (response) {
                window.location.href = '/users';
            }
        })
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


let btnsBlock = document.querySelectorAll(".btn-block-unblock");
if (btnsBlock) {
    for (let btn of btnsBlock) {
        btn.onclick = () => {
            let data = {
                ID: Number(btn.getAttribute("data-user_id"))
            }
            Send("POST", "/users/user-blocked-switch", data, (res) => {
                if (res) {
                    console.log(res)
                    if (btn.innerHTML === "Заблокировать") {
                        btn.innerHTML = "Разблокировать";
                        btn.className = "btn unblock"
                    } else {
                        btn.innerHTML = "Заблокировать";
                        btn.className = "btn block"
                    }
                }
            })
        }
    }
}