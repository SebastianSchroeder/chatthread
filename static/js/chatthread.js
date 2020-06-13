function replyLinkHandler(e) {
    e.preventDefault()

    const post = this.closest(".ct-post") // TODO Polyfill needed
    let postFormWrapper
    for (let i = 0; i < post.children.length; i++) {
        const child = post.children[i]
        if (child.classList.contains("ct-reply-form-wrapper")) {
            postFormWrapper = child
            break
        }
    }
    if (postFormWrapper.getElementsByTagName("form").length === 0) {
        const postFormTemplate = document.getElementById("ct-post-form-template").innerHTML
        const actionUrl = this.getAttribute("href")
        postFormWrapper.innerHTML = postFormTemplate
        postFormWrapper.getElementsByTagName("form")[0].setAttribute("action", actionUrl)
    } else {
        toggleVisibility(postFormWrapper)
    }
}

function toggleVisibility(element) {
    if (element.style.display === "none") {
        element.style.display = "block"
    } else {
        element.style.display = "none"
    }
}

const replyLinks = document.getElementsByClassName("ct-reply-action")

for (let i = 0; i < replyLinks.length; i++) {
    const replyLink = replyLinks[i]
    replyLink.addEventListener('click', replyLinkHandler, false)
}