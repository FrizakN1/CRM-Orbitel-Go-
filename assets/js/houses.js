let search = document.querySelector("#search");
if (search) {
    let section = document.querySelector("section")
    let addresses = section.querySelectorAll("span");
    console.log(addresses)
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
}