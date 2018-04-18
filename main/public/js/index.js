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
    self.name = window.localStorage.getItem('name');
    self.roomContent = $('#roomContent');
    self.sendButton = $('#send');
    self.sendButton.click(function () {
        self.sendMessage();
    });
    let ws = new WebSocket("ws://" + document.location.host + "/ws");
    ws.onopen = function () {};
    ws.onmessage = function (evt) {
        console.log(evt);
        // var node = document.getElementById('content');
        // var p = document.createElement('p');
        // p.innerHTML = evt.data;
        // node.appendChild(p);
        // node.scrollTop = node.scrollHeight - node.offsetHeight;
        self.takeMessage(JSON.parse(evt.data))
    };
    ws.onclose = function (evt) {};
    ws.onerror = function (evt) {};
    self.ws = ws;
};

chatRoot.prototype.takeMessage = function (data) {
    let c = this.name === data.name ? 'myself' : '';
    let current = this.roomContent.children('.current');
    current.append(`<div class="chatroom-log ${c}">
        <span class="avatar"><img src="https://avatars0.githubusercontent.com/u/30884897?s=40&v=4" alt="${data.name}"></span>
        <span class="time"><b data-id="Q-2xC-3e2q46">${data.name}</b> 2018/4/16 下午5:51:30</span>
        <span class="detail">${data.info}</span>
     </div>`);
    let scrollTop = current[0].scrollHeight;
    this.roomContent.scrollTop(scrollTop);
};

chatRoot.prototype.sendMessage = function () {
    let t = this.sendButton.prev('textarea');
    let val = t.val();
    t.val('');
    this.ws.send(val)
};


new chatRoot();