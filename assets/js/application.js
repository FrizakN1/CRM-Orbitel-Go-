let ths = document.querySelectorAll('th');
if (ths) {
    [].forEach.call(ths, function(th, index) {
        th.addEventListener('click', function() {
            [].forEach.call(ths, function (th) {
                th.style.color = "#b0b0b0";
            });
            th.style.color = "#0177fd";
            sortColumn(index);
        });
    });
}

let columnFlags = [0, 0, 0, 0, 0, 0];

function sortColumn(index) {
    let tbody = document.querySelector('tbody');
    let rows = tbody.querySelectorAll('tr');

    let newRows = Array.from(rows);

    newRows.sort(function(rowA, rowB) {
        let cellA = rowA.querySelectorAll('td')[index].innerHTML;
        let cellB = rowB.querySelectorAll('td')[index].innerHTML;

        if (columnFlags[index] === 1) {
            switch (true) {
                case cellB > cellA: return 1;
                case cellB < cellA: return -1;
                case cellB === cellA: return 0;
            }
        } else {
            switch (true) {
                case cellA > cellB: return 1;
                case cellA < cellB: return -1;
                case cellA === cellB: return 0;
            }
        }
    });

    if (columnFlags[index] === 1) {
        columnFlags = [0, 0, 0, 0, 0, 0]
    } else {
        columnFlags = [0, 0, 0, 0, 0, 0]
        columnFlags[index] = 1;
    }

    tbody.innerHTML = '';

    newRows.forEach(function(newRow) {
        tbody.appendChild(newRow);
    });
}