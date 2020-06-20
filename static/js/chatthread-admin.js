const chatThreadAdminMain = document.getElementById("ct-admin-main");

chatThreadAdminMain.addEventListener(
    'click',
    function (event) {
        if (event.target.classList.contains("ct-admin-page-delete-action")) {
            deletePageHandler(event)
        }
    },
    false
);

let pageSubmitForm = document.getElementById("ct-admin-page-form");

pageSubmitForm.addEventListener("submit", addPageHandler, false)

function addPageHandler(e) {
    e.preventDefault()
    const form = e.target
    const xhr = new XMLHttpRequest()
    xhr.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            location.reload();
        }
    }
    xhr.open("POST", form.getAttribute("action"), true)
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(JSON.stringify(parseFormElements(form)))
}

function deletePageHandler(e) {
    e.preventDefault()
    const deleteLink = e.target
    const httpRequest = new XMLHttpRequest()
    httpRequest.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            const page = deleteLink.closest(".ct-admin-page")
            page.remove();
        }
    }
    httpRequest.open("DELETE", deleteLink.getAttribute("href"), true)
    httpRequest.send()
}

function parseFormElements(form) {
    let data = {}
    for (let i = 0; i < form.elements.length; i++) {
        const element = form.elements[i]
        if (element.type === "text" || element.type === "url" || element.type === "textarea") {
            data[element.name] = element.value
        }
    }
    return data
}