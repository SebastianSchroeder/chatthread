package admin

import "chatthread.net/app/main/util/hash"

var ChatThreadJsHash = hash.ComputeFileHash("static/js/chatthread-admin.js")
var ChatThreadCssHash = hash.ComputeFileHash("static/css/chatthread-admin.css")
