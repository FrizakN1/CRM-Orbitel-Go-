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
