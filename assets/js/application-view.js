let startProcessing = document.querySelector("#start-processing");
if (startProcessing) {
    startProcessing.onclick = () => {
        Send("POST", "/application/processing-start/"+startProcessing.getAttribute("data-application_id"), null, (res) => {
            if (res) {
                window.location.reload()
            }
        })
    }
}

let stopProcessing = document.querySelector("#stop-processing");
if (stopProcessing) {
    stopProcessing.onclick = () => {
        let data = {
            Comment: ""
        }
        Send("POST", "/application/processing-stop/"+stopProcessing.getAttribute("data-application_id"), data, (res) => {
            if (res) {
                window.location.reload()
            }
        })
    }
}

let reopenProcessing = document.querySelector("#reopen-processing");
if (reopenProcessing) {
    reopenProcessing.onclick = () => {
        Send("POST", "/application/processing-reopen/"+reopenProcessing.getAttribute("data-application_id"), null, (res) => {
            if (res) {
                window.location.reload()
            }
        })
    }
}

let departmentSubmit = document.querySelector("#department-submit");
if (departmentSubmit) {
    departmentSubmit.onclick = () => {
        let department = document.querySelector("#department").getAttribute("value");
        if (department) {
            let nameDepartment;
            switch (department) {
                case "1": nameDepartment = "Админы"; break;
                case "2": nameDepartment = "Менеджеры"; break;
                case "3": nameDepartment = "Техники"; break;
                case "4": nameDepartment = "ТВ"; break;
            }
            let data = {
                Application: {
                    Department: {
                        ID: Number(department),
                        Name: nameDepartment,
                    }
                }
            }
            Send("POST","/application/switch-department/"+departmentSubmit.getAttribute("data-application_id"), data, (res) => {
                if (res) {
                    window.location.reload()
                }
            })
        }
    }
}

let prioritySubmit = document.querySelector("#priority-submit");
if (prioritySubmit) {
    prioritySubmit.onclick = () => {
        let priority = document.querySelector("#priority").getAttribute("value");
        if (priority) {
            let namePriority;
            switch (priority) {
                case "1": namePriority = "Критический"; break;
                case "2": namePriority = "Высокий"; break;
                case "3": namePriority = "Нормальный"; break;
                case "4": namePriority = "Низкий"; break;
            }
            let data = {
                Application: {
                    Priority: {
                        ID: Number(priority),
                        Name: namePriority,
                    }
                }
            }
            Send("POST","/application/switch-priority/"+prioritySubmit.getAttribute("data-application_id"), data, (res) => {
                if (res) {
                    window.location.reload()
                }
            })
        }
    }
}

let commentSubmit = document.querySelector("#comment-submit");
if (commentSubmit) {
    commentSubmit.onclick = () => {
        let comment = document.querySelector("#comment");
        if (comment) {
            let data = {
                Comment: comment.value
            }
            Send("POST","/application/switch-comment/"+commentSubmit.getAttribute("data-application_id"), data, (res) => {
                if (res) {
                    window.location.reload()
                }
            })
        }
    }
}