function getUserData () {
    let userID = document.querySelector("h2").getAttribute("user-id");
    Send("GET", "/users/get-user/"+userID, null, (res) => {
        if (res) {
            console.log(res)
            let login = document.querySelector("#login");
            let name = document.querySelector("#name");
            let role = document.querySelector("#select-button");
            login.value = res.Login;
            name.value = res.Name;
            let roles = document.querySelector("#role-list");
            for (let i = 1; i <= 21; i=i+2) {
                if (Number(roles.childNodes[i].getAttribute("value")) === res.Role) {
                    role.innerHTML = roles.childNodes[i].innerHTML;
                    role.setAttribute("value", res.Role)
                    role.style.color = "#ffffff";
                    role.style.backgroundColor = "#0177fd";
                }
            }
            let btnBlock = document.querySelector("#btn-block")
            if (res.Blocked === 1) {
                btnBlock.className = "btn unblock btn-block-unblock";
                btnBlock.innerHTML = "Разблокиовать";
            }
        }
    })
}

getUserData()