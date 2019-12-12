var sock = null;
var wsurl = "ws://127.0.0.1:1234/websocket";
window.onload = function () {
    console.log("onload");
    sock = new WebSocket(wsurl);
    sock.onopen = function () {
        console.log("connected to " + wsurl);
    }
    sock.onclose = function (e) {
        console.log("connection closed (" + e.code + ")");
    }

    sock.onmessage = function (e) {
        console.log("message received: " + e.data);
        result.innerHTML = e.data
    }
    var str = "";
    var btn = document.getElementsByTagName("button");
    var result = document.getElementById("result");
    var process = document.getElementById("process");
    for (var i = 0; i < btn.length; i++) {
        btn[i].onclick = function () {
            if (result.innerHTML == "" && this.value == ".") {
                result.innerHTML = "0.";
                str = "0.";
            }
            else if (!isNaN(this.value) || this.value == ".") {
                result.innerHTML += this.value;
                str += this.value;
            }
            else {
                if (this.value == "=") {
                    sock.send(str)
                    process.innerHTML = result.innerHTML
                }
                else if (this.value == "ce") {
                    result.innerHTML = result.innerHTML.substr(0, result.innerHTML.length - 1);
                    str = str.substr(0, str.length - 1);
                }
                else if (this.value == "c") {
                    result.innerHTML = "";
                    process.innerHTML = "";
                    str = "";
                }
                else {
                    result.innerHTML += this.value;
                    if (this.value == "²") {
                        str += "p";
                    }
                    else if (this.value == "√") {
                        str += "g";
                    }
                    else if (this.value == "×") {
                        str += "m";
                    }
                    else if (this.value == "÷") {
                        str += "d";
                    }
                    else {
                        str += this.value;
                    }
                }
            }
        }
    }
};