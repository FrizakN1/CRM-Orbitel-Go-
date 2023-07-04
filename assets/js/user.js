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

function checkUserForm(uri) {
    let login = document.querySelector('#Login')
    let name = document.querySelector('#Name')
    let password = document.querySelector('#Password')
    let passwordConfirm = document.querySelector("#PasswordConfirm")
    let role = document.querySelector('#Role')
    let department;

    let status = true;

    if (login.value.length < 3) {
        status = false;
        login.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
    } else {
        login.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
    }

    if (name.value.length < 3) {
        status = false;
        name.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
    } else {
        name.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
    }

    if (window.location.href.includes("edit")) {
        if (password.value !== "") {
            if (passwordConfirm.value !== password.value) {
                status = false;
                passwordConfirm.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
            } else {
                passwordConfirm.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
            }
        }
    } else {
        if (password.value.length < 3) {
            status = false;
            password.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
        } else {
            password.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
        }

        if (passwordConfirm.value !== password.value) {
            status = false;
            passwordConfirm.parentNode.querySelector("i").style.backgroundColor = "#ff0000";
        } else {
            passwordConfirm.parentNode.querySelector("i").style.backgroundColor = "#0177fd";
        }
    }

    if (role.getAttribute("value") === "") {
        status = false;
        role.style.border = "1px solid #ff0000"
    } else {
        role.style.border = "1px solid #0177fd"
    }

    switch (role.getAttribute("value")) {
        case "1": department = 1; break;
        case "2": department = 1; break;
        case "3": department = 2; break;
        case "4": department = 3; break;
        case "5": department = 4; break;
        case "6": department = 1; break;
    }

    if (status) {
        sendData(uri, login.value, name.value, password.value, role.getAttribute("value"), department)
    }
}

function sendData(uri, login, name, password, role, department) {
    let createData = {
        Login: String(login),
        Name: String(name),
        Password: String(password),
        Role: {
            ID: Number(role),
        },
        Department: {
            ID: Number(department),
        },
    }

    Send('PUT', uri, createData, (response) => {
        if (response) {
            window.location.href = '/users';
        }
    })
}

let userCreate = document.querySelector("#createUser");
if (userCreate) {
    userCreate.onclick = function (event) {
        checkUserForm("/users/create")
    }
}

let userUpdate = document.querySelector("#updateUser");
if (userUpdate ) {
    userUpdate.onclick = function (event) {
        let userID = window.location.href.split("edit/")[1];
        checkUserForm('/users/update/'+userID)
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