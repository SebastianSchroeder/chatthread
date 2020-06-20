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

function deletePageHandler(e) {
    e.preventDefault()
    const deleteLink = e.target
    const httpRequest = new XMLHttpRequest()
    httpRequest.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            const page = deleteLink.closest(".ct-admin-page")
            page.remove();
        }
    }
    httpRequest.open("DELETE", deleteLink.getAttribute("href"), true)
    httpRequest.send()
}