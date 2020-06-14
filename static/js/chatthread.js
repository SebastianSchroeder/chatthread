const chatThreadMain = document.getElementById("ct-main");

chatThreadMain.addEventListener(
    'click',
    function (event) {
        if (event.target.classList.contains("ct-reply-action")) {
            replyLinkHandler(event)
        }
    },
    false
);

function replyLinkHandler(e) {
    e.preventDefault()
    const replyLink = e.target

    const post = replyLink.closest(".ct-post")
    let postFormWrapper = retrieveChildWithClass(post, "ct-reply-form-wrapper");
    if (postFormWrapper.getElementsByTagName("form").length === 0) {
        const postFormTemplate = document.getElementById("ct-post-form-template").innerHTML
        const actionUrl = replyLink.getAttribute("href")
        postFormWrapper.innerHTML = postFormTemplate
        postFormWrapper.getElementsByTagName("form")[0].setAttribute("action", actionUrl)
    } else {
        toggleVisibility(postFormWrapper)
    }
}

function retrieveChildWithClass(element, elementClass) {
    for (let i = 0; i < element.children.length; i++) {
        const child = element.children[i]
        if (child.classList.contains(elementClass)) {
            return child
        }
    }
    return null
}

function toggleVisibility(element) {
    if (element.style.display === "none") {
        element.style.display = "block"
    } else {
        element.style.display = "none"
    }
}

/**
 * element.closest polyfill
 */
if (!Element.prototype.matches) {
    Element.prototype.matches = Element.prototype.msMatchesSelector ||
        Element.prototype.webkitMatchesSelector
}

if (!Element.prototype.closest) {
    Element.prototype.closest = function (s) {
        let el = this

        do {
            if (Element.prototype.matches.call(el, s)) return el
            el = el.parentElement || el.parentNode
        } while (el !== null && el.nodeType === 1)
        return null
    }
}