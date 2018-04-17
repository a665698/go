let chatRoot = function () {
    this.documentInit();
    this.init();
};

chatRoot.prototype.documentInit = function () {
    // 禁止右键和刷新
    document.oncontextmenu = function () {
        return false;
    };
    document.onkeydown = function (event) {
        let e = event || window.event || arguments.callee.caller.arguments[0];
        if (e && e.keyCode == 116) {
            return false;
        }
    };
};
chatRoot.prototype.init = function () {
    let self = this;
    self.roomContent = $('#roomContent');
    self.sendButton = $('#send');
    self.sendButton.click(function () {
        self.sendMessage();
    });
    let ws = new WebSocket("ws://" + document.location.host + "/ws");
    ws.onopen = function () {};
    ws.onmessage = function (evt) {
        // console.log(evt, 'message');
        // var node = document.getElementById('content');
        // var p = document.createElement('p');
        // p.innerHTML = evt.data;
        // node.appendChild(p);
        // node.scrollTop = node.scrollHeight - node.offsetHeight;
        self.takeMessage(evt.data)
    };
    ws.onclose = function (evt) {};
    ws.onerror = function (evt) {};
    self.wx = ws;
};

chatRoot.prototype.takeMessage = function (data) {
    this.roomContent.children('.current').append(`<div class="chatroom-log myself">
        <span class="avatar"><img src="https://avatars0.githubusercontent.com/u/30884897?s=40&v=4" alt="我"></span>
        <span class="time"><b data-id="Q-2xC-3e2q46">我</b> 2018/4/16 下午5:51:30</span>
        <span class="detail">${data}</span>
     </div>`)
};

chatRoot.prototype.sendMessage = function () {
    let val = self.sendButton.prev('textarea').val();
    self.ws.send(val)
};



function submit() {
    var inp = document.getElementById("test").value;
    ws.send(inp);
    document.getElementById("test").value = '';
}