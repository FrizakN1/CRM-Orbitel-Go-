const setCookie = (name, value, days = 7, path = '/') => {
    const expires = new Date(Date.now() + days * 864e5).toUTCString()
    document.cookie = name + '=' + encodeURIComponent(value) + '; expires=' + expires + '; path=' + path;
}

const getCookie = (name) => {
    return document.cookie.split('; ').reduce((r, v) => {
        const parts = v.split('=')
        return parts[0] === name ? decodeURIComponent(parts[1]) : r
    }, '')
}

const deleteCookie = (name, path) => {
    setCookie(name, '', -1, path)
}

let switchTheme = document.querySelector(".switch");

if (switchTheme) {
    switchTheme.onclick = () => {
        if (getCookie("crm-theme") === "dark") {
            switchTheme.classList.remove("on");
            setCookie("crm-theme", "light");
            changeLinkPath("dark", "light")
        } else {
            switchTheme.classList.add("on");
            setCookie("crm-theme", "dark");
            changeLinkPath("light", "dark")
        }
    }

    function changeLinkPath(oldTheme, newTheme) {
        let links = document.querySelectorAll("link");
        for (let link of links) {
            if (link.href.includes(oldTheme) || link.href.includes(newTheme)) {
                link.href = link.href.replace(oldTheme, newTheme)
            }
        }
    }
}

let theme = getCookie("crm-theme");
if (theme === "dark"){
    switchTheme.classList.add("on");
}

let menuBtn = document.querySelector('.menuBtn');
let navigation = document.querySelector('nav');
if (menuBtn) {
    menuBtn.onclick = function () {
        navigation.classList.toggle('active');
    }
}

window.onload = function () {
    document.body.classList.add('loaded_hiding');
    window.setTimeout(function () {
        document.body.classList.add('loaded');
        document.body.classList.remove('loaded_hiding');
    }, 500);
}

let exit = document.querySelector("#exit");
if (exit){
    exit.onclick = () => {
        Send("DELETE", "/users/exit", null, function (res){
            if (res) {
                window.location.reload()
            }
        })
    }
}

let selects = document.querySelectorAll(".select")
if (selects.length > 0) {
    for (let select of selects) {
        let selectButton = select.querySelector("div")
        select.onclick = () => {
            select.classList.toggle("active");
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
                let applicationsTable = document.querySelector("#applications-table");
                if (applicationsTable) {
                    console.log(1)
                    switchTabs(null)
                }
            }
        }
    }
    window.onclick = (event) => {
        if (selects.length > 1) {
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
        } else {
            if (!selects[0].className.includes("active")) {
                selects[0].classList.remove("active");
            }
        }
    }
}