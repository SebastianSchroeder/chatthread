package presentation

import "chatthread.net/app/main/util/hash"

var chatThreadJsHash = hash.ComputeFileHash("static/js/chatthread.js")
var chatThreadCssHash = hash.ComputeFileHash("static/css/chatthread.css")
