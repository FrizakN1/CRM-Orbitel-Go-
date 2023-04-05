let optional = document.querySelector(".optional");
if (optional) {
    for (let i = 1; i < 8; i=i+2) {
        optional.childNodes[i].onclick = () => {
            for (let j = 1; j < 8; j=j+2) {
                optional.childNodes[j].classList.remove("active")
            }
            optional.childNodes[i].classList.add("active");
        }
    }
}