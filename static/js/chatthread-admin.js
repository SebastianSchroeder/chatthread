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
    const httpRequest = new XMLHttpRequest()
    httpRequest.onreadystatechange = function () {
        if (this.readyState === 4 && this.status === 200) {
            location.reload();
        }
    }
    httpRequest.open("POST", form.getAttribute("action"), true)
    httpRequest.send(new FormData(form))
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